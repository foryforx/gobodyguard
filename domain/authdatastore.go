package domain

import (
	"github.com/karuppaiah/gobodyguard/models"
)

// AuthDataStore defines operations expected from data storage entity
type AuthDataStore interface {
	// Principal
	GetPrincipal(UUID string) (models.Principal, error)
	AddPrincipal(principal models.Principal) (models.Principal, error)
	UpdatePrincipal(principal models.Principal) (models.Principal, error)
	DeletePrincipal(UUID string) error

	// Resource
	GetResource(UUID string) (models.Resource, error)
	AddResource(resource models.Resource) (models.Resource, error)
	UpdateResource(resource models.Resource) (models.Resource, error)
	DeleteResource(UUID string) error

	// Operation
	GetOperation(UUID string) (models.Operation, error)
	AddOperation(operation models.Operation) (models.Operation, error)
	UpdateOperation(operation models.Operation) (models.Operation, error)
	DeleteOperation(UUID string) error

	// Policy
	GetPolicy(UUID string) (models.Policy, error)
	GetPolicyForAllMatch(principalUUID string, resourceUUID string,
		operationUUID string) ([]models.Policy, error)
	AddPolicy(policy models.Policy) (models.Policy, error)
	UpdatePolicy(policy models.Policy) (models.Policy, error)
	DeletePolicy(UUID string) error
}
