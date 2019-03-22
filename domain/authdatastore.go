package domain

import (
	"github.com/karuppaiah/gobodyguard/models"
)

// AuthDataStore defines operations expected from data storage entity
type AuthDataStore interface {
	// Principal
	FetchPrincipal(uuid string) (models.Principal, error)
	AddPrincipal(principal models.Principal) (models.Principal, error)
	UpdatePrincipal(principal models.Principal) (models.Principal, error)
	DeletePrincipal(uuid string) error

	// Resource
	FetchResource(uuid string) (models.Resource, error)
	AddResource(resource models.Resource) (models.Resource, error)
	UpdateResource(resource models.Resource) (models.Resource, error)
	DeleteResource(uuid string) error

	// Operation
	FetchOperation(uuid string) (models.Operation, error)
	AddOperation(operation models.Operation) (models.Operation, error)
	UpdateOperation(operation models.Operation) (models.Operation, error)
	DeleteOperation(uuid string) error

	// Permission
	FetchPermission(uuid string) (models.Permission, error)
	AddPermission(permission models.Permission) (models.Permission, error)
	UpdatePermission(permission models.Permission) (models.Permission, error)
	DeletePermission(uuid string) error

	// Tags
	AddTagsToPermission(uuid string, tags []string) error
	RemoveTagsFromPermission(uuid string, tag string) error
}
