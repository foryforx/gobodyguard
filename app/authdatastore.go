package app

// AuthDataStore defines operations expected from data storage entity
type AuthDataStore interface {
	// Principal
	GetPrincipal(UUID string) (Principal, error)
	AddPrincipal(principal Principal) (Principal, error)
	UpdatePrincipal(principal Principal) (Principal, error)
	DeletePrincipal(UUID string) error

	// Resource
	GetResource(UUID string) (Resource, error)
	AddResource(resource Resource) (Resource, error)
	UpdateResource(resource Resource) (Resource, error)
	DeleteResource(UUID string) error

	// Operation
	GetOperation(UUID string) (Operation, error)
	AddOperation(operation Operation) (Operation, error)
	UpdateOperation(operation Operation) (Operation, error)
	DeleteOperation(UUID string) error

	// Policy
	GetPolicy(UUID string) (Policy, error)
	GetPolicyForAllMatch(principalUUID string, resourceUUID string,
		operationUUID string) ([]Policy, error)
	AddPolicy(policy Policy) (Policy, error)
	UpdatePolicy(policy Policy) (Policy, error)
	DeletePolicy(UUID string) error

	LoadAccess() (map[string](map[string][]string), error)
}
