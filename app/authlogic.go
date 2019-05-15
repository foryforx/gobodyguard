package app

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// AuthLogic represents the entity which has connection to dependent modules
type AuthLogic struct {
	AuthRepo AuthDataStore
}

// NewAuthLogic returns a service initialized with repository.
func NewAuthLogic(db *gorm.DB) AuthOpns {
	var authRepo = NewAuthRepository(db)
	return &AuthLogic{AuthRepo: authRepo}
}

// GetPrincipal will get the principal matching UUID
func (a AuthLogic) GetPrincipal(UUID string) (Principal, error) {
	var principal Principal
	principal, err := a.AuthRepo.GetPrincipal(UUID)
	if err != nil {
		return principal, errors.Wrapf(err, "SVC:GetPrincipal failed")
	}
	return principal, nil
}

// AddPrincipal will add a new principal from system
func (a AuthLogic) AddPrincipal(principal Principal) (Principal, error) {
	principal.UUID = NewUUID()
	principal, err := a.AuthRepo.AddPrincipal(principal)
	if err != nil {
		return principal, errors.Wrapf(err, "SVC:AddPrincipal failed")
	}
	return principal, nil
}

// UpdatePrincipal will update the principal in the system
func (a AuthLogic) UpdatePrincipal(principal Principal) (Principal, error) {
	principal, err := a.AuthRepo.UpdatePrincipal(principal)
	if err != nil {
		return principal, errors.Wrapf(err, "SVC:UpdatePrincipal failed")
	}
	return principal, nil
}

// DeletePrincipal will delete the principal from system
func (a AuthLogic) DeletePrincipal(UUID string) error {
	err := a.AuthRepo.DeletePrincipal(UUID)
	if err != nil {
		return errors.Wrapf(err, "SVC: DeletePrincipal failed")
	}
	return nil
}

// GetResource will get a the resource matching the UUID
func (a AuthLogic) GetResource(UUID string) (Resource, error) {
	var resource Resource
	resource, err := a.AuthRepo.GetResource(UUID)
	if err != nil {
		return resource, errors.Wrapf(err, "SVC:GetResource failed")
	}
	return resource, nil
}

// AddResource will add a new resource to the system
func (a AuthLogic) AddResource(resource Resource) (Resource, error) {
	resource.UUID = NewUUID()
	resource, err := a.AuthRepo.AddResource(resource)
	if err != nil {
		return resource, errors.Wrapf(err, "SVC:AddResource failed")
	}
	return resource, nil
}

// UpdateResource will update the resource from system
func (a AuthLogic) UpdateResource(resource Resource) (Resource, error) {
	resource, err := a.AuthRepo.UpdateResource(resource)
	if err != nil {
		return resource, errors.Wrapf(err, "SVC:UpdateResource failed")
	}
	return resource, nil
}

// DeleteResource will delete the resource from system
func (a AuthLogic) DeleteResource(UUID string) error {
	err := a.AuthRepo.DeleteResource(UUID)
	if err != nil {
		return errors.Wrapf(err, "SVC: DeleteResource failed")
	}
	return nil
}

// GetOperation will get the operation matching with UUID
func (a AuthLogic) GetOperation(UUID string) (Operation, error) {
	var operation Operation
	operation, err := a.AuthRepo.GetOperation(UUID)
	if err != nil {
		return operation, errors.Wrapf(err, "SVC:GetOperation failed")
	}
	return operation, nil
}

// AddOperation will add a new operation to the system
func (a AuthLogic) AddOperation(operation Operation) (Operation, error) {
	operation.UUID = NewUUID()
	log.Debugln(operation)
	operation, err := a.AuthRepo.AddOperation(operation)
	if err != nil {
		return operation, errors.Wrapf(err, "SVC:AddOperation failed")
	}
	return operation, nil
}

// UpdateOperation will update the operation from system
func (a AuthLogic) UpdateOperation(operation Operation) (Operation, error) {
	operation, err := a.AuthRepo.UpdateOperation(operation)
	if err != nil {
		return operation, errors.Wrapf(err, "SVC:UpdateOperation failed")
	}
	return operation, nil
}

// DeleteOperation will delete the operation from system
func (a AuthLogic) DeleteOperation(UUID string) error {
	err := a.AuthRepo.DeleteOperation(UUID)
	if err != nil {
		return errors.Wrapf(err, "SVC: DeleteOperation failed")
	}
	return nil
}

// AddPermission will add a permission for the principal, operation and resource
func (a AuthLogic) AddPermission(principalUUID string, resourceUUID string,
	operationUUID string, userName string,
	permission string) (Policy, error) {
	var policy Policy
	policy.UUID = NewUUID()
	policy.PrincipalUUID = principalUUID
	policy.ResourceUUID = resourceUUID
	policy.OperationUUID = operationUUID
	policy.Permission = permission
	policy, err := a.AuthRepo.AddPolicy(policy)
	if err != nil {
		return policy, errors.Wrapf(err, "SVC:AddPermission failed")
	}
	return policy, nil
}

// DeletePermission will remove the permission of the respective UUID
func (a AuthLogic) DeletePermission(UUID string) error {
	err := a.AuthRepo.DeletePolicy(UUID)
	if err != nil {
		return errors.Wrapf(err, "SVC: DeletePermission failed")
	}
	return nil
}

// GetPermission will get the permission for principal, resource and operation UUID
func (a AuthLogic) GetPermission(principalUUID string, resourceUUID string,
	operationUUID string) (
	PermissionStatusCode, error) {
	// var policys = make([]Policy, 0)
	// policys, err := a.AuthRepo.GetPolicyForAllMatch(principalUUID, resourceUUID, operationUUID)
	opnsAccess := GetMemData().OperationAccess
	access, ok := opnsAccess[operationUUID]
	if !ok {
		return Denied, nil
	}
	resources, ok := access[principalUUID]
	if !ok {
		return Denied, nil
	}
	for _, v := range resources {
		if v == resourceUUID {
			return Granted, nil
		}
	}
	return Denied, nil
}
