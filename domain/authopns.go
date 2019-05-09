package domain

import (
	"github.com/karuppaiah/gobodyguard/models"
)

// AuthOpns defines operations exposed from domain layer
type AuthOpns interface {
	// Principal
	GetPrincipal(uuid string) (models.Principal, error)
	AddPrincipal(principal models.Principal) (models.Principal, error)
	UpdatePrincipal(principal models.Principal) (models.Principal, error)
	DeletePrincipal(uuid string) error

	// Resource
	GetResource(uuid string) (models.Resource, error)
	AddResource(resource models.Resource) (models.Resource, error)
	UpdateResource(resource models.Resource) (models.Resource, error)
	DeleteResource(uuid string) error

	// Operation
	GetOperation(uuid string) (models.Operation, error)
	AddOperation(operation models.Operation) (models.Operation, error)
	UpdateOperation(operation models.Operation) (models.Operation, error)
	DeleteOperation(uuid string) error

	// Permission
	AddPermission(principalUUID string, resourceUUID string,
		operationUUID string, userName string,
		permissionCode string) (models.Policy, error)
	DeletePermission(uuid string) error

	// Verification Policy
	GetPermission(principalUUID string, resourceUUID string,
		operationUUID string) (
		models.PermissionStatusCode, error)
}
