package lib

import "gorm.io/gorm"

// DatabaseInstance represent a singleton collection of database connection, a single object for all necessary database operations.
type DatabaseInstance struct {
	mysql *gorm.DB
}

// NewDatabaseInstance constructs DatabaseInstance.
func NewDatabaseInstance() *DatabaseInstance {
	return &DatabaseInstance{}
}
