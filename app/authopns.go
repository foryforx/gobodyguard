package app

// AuthOpns defines operations exposed from domain layer
type AuthOpns interface {
	// Principal
	GetPrincipal(uuid string) (Principal, error)
	AddPrincipal(principal Principal) (Principal, error)
	UpdatePrincipal(principal Principal) (Principal, error)
	DeletePrincipal(uuid string) error

	// Resource
	GetResource(uuid string) (Resource, error)
	AddResource(resource Resource) (Resource, error)
	UpdateResource(resource Resource) (Resource, error)
	DeleteResource(uuid string) error

	// Operation
	GetOperation(uuid string) (Operation, error)
	AddOperation(operation Operation) (Operation, error)
	UpdateOperation(operation Operation) (Operation, error)
	DeleteOperation(uuid string) error

	// Permission
	AddPermission(principalUUID string, resourceUUID string,
		operationUUID string, userName string,
		permissionCode string) (Policy, error)
	DeletePermission(uuid string) error

	// Verification Policy
	GetPermission(principalUUID string, resourceUUID string,
		operationUUID string) (
		PermissionStatusCode, error)
}
