package app

import (
	"time"
)

// Principal represents the identity of a specific user or group of users.
type Principal struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedBy string     `json:"updatedBy"`
}

// Resource is understood to be a specific entity or container
// upon which permissions may be applied.
type Resource struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedBy string     `json:"updatedBy"`
}

// Operation is understood to be a specific function
// that may be performed by a Principal on a Resource
type Operation struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedBy string     `json:"updatedBy"`
}

// Policy that can be read as declarative statements using Principal, Resource and Operation
type Policy struct {
	UUID          string    `json:"uuid"`
	Principal     Principal `gorm:"foreignkey:PrincipalUUID;association_foreignkey:UUID"`
	PrincipalUUID string    `json:"principalUUID"`
	Resouce       Resource  `gorm:"foreignkey:ResourceUUID;association_foreignkey:UUID"`
	ResourceUUID  string    `json:"resourceUUID"`
	Operation     Operation `gorm:"foreignkey:OperationUUID;association_foreignkey:UUID"`
	OperationUUID string    `json:"operationUUID"`
	Permission    string    `json:"authStatus"`
}

// // PermissionTags that can be attached to permission for further grouping and enhacement
// type PermissionTags struct {
// 	UUID           string `json:"uuid"`
// 	Tag            string `json:"tag"`
// 	PermissionUUID string `json:"permissionUUID"`
// }

///////////// PermissionStatusCode Action enum : START

// PermissionStatusCode is the code used for auth status
type PermissionStatusCode int

const (
	// Granted is the state When principal is specifically granted permission
	// on resource for a operation
	Granted PermissionStatusCode = iota
	// Denied is the state When principal is specifically denied permission
	// on resource for a operation
	Denied
)

// PermissionStatusCodeNames maps PermissionStatusCode to its string representation
var PermissionStatusCodeNames = map[PermissionStatusCode]string{
	Granted: "Granted",
	Denied:  "Denied",
}

// PermissionStatusCodeIDs maps string to PermissionStatusCode type
var PermissionStatusCodeIDs = map[string]PermissionStatusCode{
	"Granted": Granted,
	"Denied":  Denied,
}

// String will convert PermissionStatusCode to string value
func (a PermissionStatusCode) String() string {
	// Return unknown if type is outside of enum range
	if _, ok := PermissionStatusCodeNames[a]; ok {
		return PermissionStatusCodeNames[a]
	}
	return PermissionStatusCodeNames[Denied]
}

///////////// PermissionStatusCode Action enum : END
