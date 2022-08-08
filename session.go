package foxkit

import (
	"context"
	"crypto/subtle"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

// creates a new session in the given redis DB
// adds the session number to the UserID to keep the session ID unique
// returns sessionID, sessionKey, err
func CreateSession(ctx *context.Context, userID *string, redisClient *redis.Client, keyLength uint32, maxSessions int, sessionDuration time.Duration) (string, string, error) {
	// dummy error
	if maxSessions >= 0 {
		return "", "", errors.New("invalid value for maximum sessions, needs to be > 0")
	}

	// generate session key
	token, err := RandomString(keyLength)
	if err != nil {
		return "", "", err
	}

	// when only one session should be possible, we'll overwrite it directly
	if maxSessions == 1 {
		ctxLocal, cancel := context.WithTimeout(*ctx, time.Second*60)
		err := redisClient.Set(ctxLocal, *userID, token, sessionDuration).Err()
		cancel()
		if err != nil {
			return "", "", err
		}
		// return the successfully created session
		return *userID, token, nil
	}

	// create vars outside the loop to avoid re-allocation
	var count int64
	var ctxLocal context.Context
	var cancel context.CancelFunc
	// find the first free session
	for sessionNumber := 0; sessionNumber < maxSessions; sessionNumber++ {
		// check if the session exists
		ctxLocal, cancel = context.WithTimeout(*ctx, time.Second*60)
		count, err = redisClient.Exists(ctxLocal, *userID+strconv.Itoa(sessionNumber)).Result()
		cancel()
		if err != nil {
			return "", "", err
		} else if count == 0 { // create the session if the slot is free
			// remove the next session
			ctxLocal, cancel = context.WithTimeout(*ctx, time.Second*60)
			err := redisClient.Del(ctxLocal, *userID+strconv.Itoa(sessionNumber+1)).Err()
			cancel()
			if err != nil {
				return "", "", err
			}
			// save the session into redis
			ctxLocal, cancel = context.WithTimeout(*ctx, time.Second*60)
			err = redisClient.Set(ctxLocal, *userID+strconv.Itoa(sessionNumber), token, sessionDuration).Err()
			cancel()
			if err != nil {
				return "", "", err
			}
			// return the successfully created session
			return *userID + strconv.Itoa(sessionNumber), token, nil
		}
	}
	// we'll create another session when the maximum number is reached and remove the first one
	// if the removal of a session failed for some reason, this will always work for new sessions

	// remove the first session
	ctxLocal, cancel = context.WithTimeout(*ctx, time.Second*60)
	err = redisClient.Del(ctxLocal, *userID+"0").Err()
	cancel()
	if err != nil {
		return "", "", err
	}
	// save the session into redis
	ctxLocal, cancel = context.WithTimeout(*ctx, time.Second*60)
	err = redisClient.Set(ctxLocal, *userID+strconv.Itoa(maxSessions), token, sessionDuration).Err()
	cancel()
	if err != nil {
		return "", "", err
	}
	// return the successfully created session
	return *userID + strconv.Itoa(maxSessions), token, nil
}

// checks if the given session is valid
func ValidateSession(ctx *context.Context, uid, token *string, redisClient *redis.Client, sessionDuration time.Duration) (bool, error) {
	var res string
	var err error

	// the UUID session extension is part of the session, so no work is needed
	ctxLocal, cancel := context.WithTimeout(*ctx, time.Second*60)
	res, err = redisClient.Get(ctxLocal, *uid).Result()
	cancel()

	if err == redis.Nil {
		// the uid has no session stored in redis, it's not valid therefore
		return false, nil
	} else if err != nil {
		return false, err
	} else if subtle.ConstantTimeCompare([]byte(res), []byte(*token)) == 1 {
		// session and token match, the session is valid
		// We'll increase the TTL to keep the session alive
		ctxLocal, cancel := context.WithTimeout(*ctx, time.Second*60)
		err = redisClient.Expire(ctxLocal, *uid, sessionDuration).Err()
		cancel()
		if err != nil {
			return false, err
		}
		return true, nil
	}
	// the session seems to not being valid
	return false, nil
}
