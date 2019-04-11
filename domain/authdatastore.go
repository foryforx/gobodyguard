package domain

import (
	"github.com/karuppaiah/gobodyguard/models"
)

// AuthDataStore defines operations expected from data storage entity
type AuthDataStore interface {
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

	// Policy
	GetPolicy(uuid string) (models.Policy, error)
	AddPolicy(policy models.Policy) (models.Policy, error)
	UpdatePolicy(policy models.Policy) (models.Policy, error)
	DeletePolicy(uuid string) error

	// Tags
	AddTagsToPolicy(uuid string, tags []string) error
	RemoveTagsFromPolicy(uuid string, tag string) error
	SearchForTagInPolicy(tag []string) error
}
