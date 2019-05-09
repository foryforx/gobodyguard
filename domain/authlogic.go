package domain

import (
	"github.com/karuppaiah/gobodyguard/helpers"
	"github.com/karuppaiah/gobodyguard/models"
	"github.com/pkg/errors"
)

// AuthLogic represents the entity which has connection to dependent modules
type AuthLogic struct {
	AuthRepo AuthDataStore
}

// NewAuthLogic returns a service initialized with repository.
func NewAuthLogic(authRepo AuthDataStore) AuthOpns {
	return &AuthLogic{AuthRepo: authRepo}
}

// GetPrincipal will get the principal matching UUID
func (a AuthLogic) GetPrincipal(UUID string) (models.Principal, error) {
	var principal models.Principal
	principal, err := a.AuthRepo.GetPrincipal(UUID)
	if err != nil {
		return principal, errors.Wrapf(err, "SVC:GetPrincipal failed")
	}
	return principal, nil
}

// AddPrincipal will add a new principal from system
func (a AuthLogic) AddPrincipal(principal models.Principal) (models.Principal, error) {
	principal.UUID = helpers.NewUUID()
	principal, err := a.AuthRepo.AddPrincipal(principal)
	if err != nil {
		return principal, errors.Wrapf(err, "SVC:AddPrincipal failed")
	}
	return principal, nil
}

// UpdatePrincipal will update the principal in the system
func (a AuthLogic) UpdatePrincipal(principal models.Principal) (models.Principal, error) {
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
func (a AuthLogic) GetResource(UUID string) (models.Resource, error) {
	var resource models.Resource
	resource, err := a.AuthRepo.GetResource(UUID)
	if err != nil {
		return resource, errors.Wrapf(err, "SVC:GetResource failed")
	}
	return resource, nil
}

// AddResource will add a new resource to the system
func (a AuthLogic) AddResource(resource models.Resource) (models.Resource, error) {
	resource.UUID = helpers.NewUUID()
	resource, err := a.AuthRepo.AddResource(resource)
	if err != nil {
		return resource, errors.Wrapf(err, "SVC:AddResource failed")
	}
	return resource, nil
}

// UpdateResource will update the resource from system
func (a AuthLogic) UpdateResource(resource models.Resource) (models.Resource, error) {
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
func (a AuthLogic) GetOperation(UUID string) (models.Operation, error) {
	var operation models.Operation
	operation, err := a.AuthRepo.GetOperation(UUID)
	if err != nil {
		return operation, errors.Wrapf(err, "SVC:GetOperation failed")
	}
	return operation, nil
}

// AddOperation will add a new operation to the system
func (a AuthLogic) AddOperation(operation models.Operation) (models.Operation, error) {
	operation.UUID = helpers.NewUUID()
	operation, err := a.AuthRepo.AddOperation(operation)
	if err != nil {
		return operation, errors.Wrapf(err, "SVC:AddPrincipal failed")
	}
	return operation, nil
}

// UpdateOperation will update the operation from system
func (a AuthLogic) UpdateOperation(operation models.Operation) (models.Operation, error) {
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
	permissionCode string) (models.Policy, error) {
	var policy models.Policy
	policy.UUID = helpers.NewUUID()
	policy.PrincipalUUID = principalUUID
	policy.ResourceUUID = resourceUUID
	policy.OperationUUID = operationUUID
	policy, err := a.AuthRepo.AddPolicy(policy)
	if err != nil {
		return policy, errors.Wrapf(err, "SVC:Addpolicy failed")
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
	models.PermissionStatusCode, error) {
	var policys = make([]models.Policy, 0)
	policys, err := a.AuthRepo.GetPolicyForAllMatch(principalUUID, resourceUUID, operationUUID)
	if err != nil {
		return models.Denied, errors.Wrapf(err, "SVC:GetOperation failed")
	}
	if len(policys) == 0 {
		return models.Denied, nil
	}
	for _, policy := range policys {
		if policy.Permission == models.Denied {
			return models.Denied, nil
		}
	}
	return models.Granted, nil
}
