package models

import (
	"time"
)

// Principal represents the identity of a specific user or group of users.
type Principal struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	CreatedBy string    `json:"createdBy"`
}

// Resource is understood to be a specific entity or container
// upon which permissions may be applied.
type Resource struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	CreatedBy string    `json:"createdBy"`
}

// Operation is understood to be a specific function
// that may be performed by a Principal on a Resource
type Operation struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	CreatedBy string    `json:"createdBy"`
}

// Permission that can be read as declarative statements using Principal, Resource and Operation
type Permission struct {
	UUID          string           `json:"uuid"`
	Principal     Principal        `gorm:"foreignkey:PrincipalUUID;association_foreignkey:UUID"`
	PrincipalUUID string           `json:"principalUUID"`
	Resouce       Resource         `gorm:"foreignkey:ResourceUUID;association_foreignkey:UUID"`
	ResourceUUID  string           `json:"resourceUUID"`
	Operation     Operation        `gorm:"foreignkey:OperationUUID;association_foreignkey:UUID"`
	OperationUUID string           `json:"operationUUID"`
	AuthStatus    AuthStatusCode   `json:"authStatus"`
	Tags          []PermissionTags `json:"tags" gorm:"foreignkey:PermissionUUID"`
}

// PermissionTags that can be attached to permission for further grouping and enhacement
type PermissionTags struct {
	UUID           string `json:"uuid"`
	Tag            string `json:"tag"`
	PermissionUUID string `json:"permissionUUID"`
}

///////////// AuthStatusCode Action enum : START

// AuthStatusCode is the code used for auth status
type AuthStatusCode int

const (
	// Unknown is the state When something wierd happens but is equally treated like
	// Denied
	Unknown AuthStatusCode = iota
	// Granted is the state When principal is specifically granted permission
	// on resource for a operation
	Granted
	// Denied is the state When principal is specifically denied permission
	// on resource for a operation
	Denied
)

// AuthStatusCodeNames maps AuthStatusCodeCode to its string representation
var AuthStatusCodeNames = map[AuthStatusCode]string{
	Unknown: "Unknown",
	Granted: "Granted",
	Denied:  "Denied",
}

// AuthStatusCodeIDs maps string to AuthStatusCode type
var AuthStatusCodeIDs = map[string]AuthStatusCode{
	"Unknown": Unknown,
	"Granted": Granted,
	"Denied":  Denied,
}

// String will convert AuthStatusCode to string value
func (a AuthStatusCode) String() string {
	// Return unknown if type is outside of enum range
	if _, ok := AuthStatusCodeNames[a]; ok {
		return AuthStatusCodeNames[a]
	}
	return AuthStatusCodeNames[Unknown]
}

///////////// AuthStatusCode Action enum : END
