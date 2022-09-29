package foxkit

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// creates a connection to Postgres using ENV
func ConnectSQL() *gorm.DB {
	// build connection URI
	uri := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"))

	// connect to the DB
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		ErrorFatal("FoxKit", err)
	}

	// check for errors
	if db.Error != nil {
		ErrorFatal("FoxKit", db.Error)
	}

	return db
}

// migrates all tables, panics on error
func AutoMigrateSQL(pg_conn *gorm.DB, inf ...interface{}) {
	if err := pg_conn.AutoMigrate(inf...); err != nil {
		ErrorFatal("FoxKit", err)
	}
}

// tried to push the data into the DB table, sets status to 500 if false
func StoreDB(c *gin.Context, pg_conn *gorm.DB, inf interface{}) bool {
	if err := pg_conn.Create(&inf).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		LogError("FoxKit", err)
		return false
	}
	return true
}

// Finds Data Entry with condition and updates it with new data
func UpdateDB(c *gin.Context, pg_conn *gorm.DB, inf interface{}, condition interface{}) bool {
	if err := pg_conn.Where(condition).Updates(inf).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		LogError("FoxKit", err)
		return false
	}
	return true
}

// Finds and Returns true with the Data Entry or false if not found
func FindDB(c *gin.Context, pg_conn *gorm.DB, inf interface{}, condition interface{}) bool {
	if err := pg_conn.Where(condition).First(inf).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		LogError("FoxKit", err)
		return false
	}
	return true
}

// Deletes Data Entry with condition
func DeleteDB(c *gin.Context, pg_conn *gorm.DB, inf interface{}, condition interface{}) bool {
	if err := pg_conn.Where(condition).Delete(inf).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		LogError("FoxKit", err)
		return false
	}
	return true
}
