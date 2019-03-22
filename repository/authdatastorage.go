package repository

import (
	"github.com/jinzhu/gorm"
)

type authStorage struct {
	Conn *gorm.DB
}

// NewWorkflowRepository To create new Repository with connection to DB
func NewWorkflowRepository(conn *gorm.DB) *AuthDataStore {
	return &authStorage{conn}
}

func (a *authStorage) GetConnection() *gorm.DB {
	return a.Conn
}
