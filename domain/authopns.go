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

	// Policy
	GrantPolicy(principalUUID string, resourceUUID string,
		operationUUID string, tags []string, userName string) (models.Policy, error)
	DenyPolicy(principalUUID string, resourceUUID string,
		operationUUID string, tags []string, userName string) (models.Policy, error)
	IsGranted(principalUUID string, resourceUUID string,
		operationUUID string, tags string) (bool, error)
	IsDenied(principalUUID string, resourceUUID string,
		operationUUID string, tags string) (bool, error)
	DeletePolicy(uuid string) error

	// Tags
	AddTagToPolicy(uuid string, tags string) error
	RemoveTagFromPolicy(uuid string, tag string) error
	SearchForTagInPolicy(tag []string) error
}
