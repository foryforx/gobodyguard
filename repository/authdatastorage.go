package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/karuppaiah/gobodyguard/domain"
	"github.com/karuppaiah/gobodyguard/models"
	"github.com/pkg/errors"
)

type authStorage struct {
	Conn *gorm.DB
}

// NewWorkflowRepository To create new Repository with connection to DB
func NewWorkflowRepository(conn *gorm.DB) *domain.AuthDataStore {
	return &authStorage{conn}
}

func (a *authStorage) GetConnection() *gorm.DB {
	return a.Conn
}

// GetPrincipal will get a principal data from datastore for a particular uuid
func (a *authStorage) GetPrincipal(UUID string) (models.Principal, error) {
	var principal models.Principal
	if len(UUID) <= 10 {
		return principal, errors.Errorf("UUID not valid:%v", UUID)
	}
	err := a.Conn.Where("uuid = ?", UUID).Find(&principal).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return principal, errors.Wrapf(err, "GetPrincipal: error %v", UUID)
	}
	return principal, nil
}

// AddPrincipal will add a principal data to datastore
func (a *authStorage) AddPrincipal(principal models.Principal) (models.Principal, error) {
	err := a.Conn.Save(&principal).Error
	if err != nil {
		return principal, errors.Wrapf(err, "AddPrincipal error %v", principal)
	}
	return principal, nil
}

// UpdatePrincipal will update the respective principal data in datastore
func (a *authStorage) UpdatePrincipal(principal models.Principal) (models.Principal, error) {
	err := a.Conn.Update(&principal).Error
	if err != nil {
		return principal, errors.Wrapf(err, "UpdatePrincipal error %v", principal)
	}
	return principal, nil
}

// DeletePrincipal will remove the principal data from data store
func (a *authStorage) DeletePrincipal(UUID string) error {
	err := a.Conn.Unscoped().Where("uuid::text = ?", UUID).Delete(&models.Principal{}).Error
	return err
}

// GetResource will get the resource data from data store for the UUID
func (a *authStorage) GetResource(UUID string) (models.Resource, error) {
	if len(UUID) <= 10 {
		return errors.Errorf("UUID not valid:%v", UUID)
	}
	var resource models.Principal
	err := a.Conn.Where("uuid = ?", UUID).Find(&resource).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return resource, errors.Wrapf(err, "GetResource: error %v", UUID)
	}
	return resource, nil
}

// AddResource will add a new resource to data store
func (a *authStorage) AddResource(resource models.Resource) (models.Resource, error) {
	err := a.Conn.Save(&resource).Error
	if err != nil {
		return resource, errors.Wrapf(err, "AddResource error %v", resource)
	}
	return resource, nil
}

// UpdateResource will update the existing resource in data store
func (a *authStorage) UpdateResource(resource models.Resource) (models.Resource, error) {
	err := a.Conn.Update(&resource).Error
	if err != nil {
		return resource, errors.Wrapf(err, "UpdateResource error %v", resource)
	}
	return resource, nil
}

// DeleteResource will delete the existing resource from data store
func (a *authStorage) DeleteResource(UUID string) error {
	err := a.Conn.Unscoped().Where("uuid::text = ?", UUID).Delete(&models.Resource{}).Error
	return err
}

// GetOperation will the operation data from data store for UUID
func (a *authStorage) GetOperation(UUID string) (models.Operation, error) {
	if len(UUID) <= 10 {
		return errors.Errorf("UUID not valid:%v", UUID)
	}
	var operation models.Operation
	err := a.Conn.Where("uuid = ?", UUID).Find(&operation).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return operation, errors.Wrapf(err, "GetOperation: error %v", UUID)
	}
	return operation, nil
}

// AddOperation will add a new operation to data store
func (a *authStorage) AddOperation(operation models.Operation) (models.Operation, error) {
	err := a.Conn.Save(&operation).Error
	if err != nil {
		return operation, errors.Wrapf(err, "AddOperation error %v", operation)
	}
	return operation, nil
}

// UpdateOperation will update the existing operation from data store
func (a *authStorage) UpdateOperation(operation models.Operation) (models.Operation, error) {
	err := a.Conn.Update(&operation).Error
	if err != nil {
		return operation, errors.Wrapf(err, "UpdateOperation error %v", operation)
	}
	return operation, nil
}

// DeleteOperation will delete the operation from data store
func (a *authStorage) DeleteOperation(UUID string) error {
	err := a.Conn.Unscoped().Where("uuid::text = ?", UUID).Delete(&models.Operation{}).Error
	return err
}

// GetPolicy will get the policy data from data store for UUID
func (a *authStorage) GetPolicy(UUID string) (models.Policy, error) {
	if len(UUID) <= 10 {
		return errors.Errorf("UUID not valid:%v", UUID)
	}
	var policy models.Policy
	err := a.Conn.Where("uuid = ?", UUID).Find(&policy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return policy, errors.Errorf(err, "GetPolicy: error %v", UUID)
	}
	return policy, nil
}

// GetPolicyForAllMatch will get the policy from data store which matches principal, operation and resource UUID
func (a *authStorage) GetPolicyForAllMatch(principalUUID string, resourceUUID string,
	operationUUID string) ([]models.Policy, error) {
	if len(principalUUID) <= 10 || len(resourceUUID) <= 10 || len(operationUUID) <= 10 {
		return errors.Errorf("UUID's not valid:%v %v %v", principalUUID, resourceUUID, operationUUID)
	}
	var policys = make([]models.Policy, 0)
	err := a.Conn.Where("principal_uuid = ? AND resource_uuid = ? AND operation_UUID = ?",
		principalUUID,
		resourceUUID,
		operationUUID).Find(&policys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return policys, errors.Errorf(err, "GetPolicy: error %v", UUID)
	}
	return policys, nil
}

// AddPolicy will add a new policy to data store
func (a *authStorage) AddPolicy(policy models.Policy) (models.Policy, error) {
	err := a.Conn.Save(&policy).Error
	if err != nil {
		return policy, errors.Errorf(err, "AddPolicy error %v", policy)
	}
	return policy, nil
}

// UpdatePolicy will update the existing policy from data store
func (a *authStorage) UpdatePolicy(policy models.Policy) (models.Policy, error) {
	err := a.Conn.Update(&policy).Error
	if err != nil {
		return policy, errors.Wrapf(err, "UpdatePolicy error %v", policy)
	}
	return policy, nil
}

// DeletePolicy will delete the existing policy from data store
func (a *authStorage) DeletePolicy(UUID string) error {
	err := a.Conn.Unscoped().Where("uuid::text = ?", UUID).Delete(&models.Policy{}).Error
	return err
}
