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
func NewWorkflowRepository(conn *gorm.DB) *AuthDataStore {
	return &authStorage{conn}
}

func (a *authStorage) GetConnection() *gorm.DB {
	return a.Conn
}

// Principal
func (a *authStorage) GetPrincipal(uuid string) (models.Principal, error) {
	if len(uuid) <= 10 {
		return errors.Errorf("UUID not valid:%v", uuid)
	}
	var principal models.Principal
	err := a.Conn.Where("uuid = ?", uuid).Find(&principal).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return principal, errors.Errorf(err, "GetPrincipal: error %v", uuid)
	}
	return principal, nil
}
func (a *authStorage) AddPrincipal(principal models.Principal) (models.Principal, error) {
	err := a.Conn.Save(&principal).Error
	if err != nil {
		return principal, errors.Errorf(err, "AddPrincipal error %v", principal)
	}
	return principal, nil
}
func (a *authStorage) UpdatePrincipal(principal models.Principal) (models.Principal, error) {

}
func (a *authStorage) DeletePrincipal(uuid string) error {

}

// Resource
func (a *authStorage) GetResource(uuid string) (models.Resource, error) {
	if len(uuid) <= 10 {
		return errors.Errorf("UUID not valid:%v", uuid)
	}
	var resource models.Principal
	err := a.Conn.Where("uuid = ?", uuid).Find(&resource).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return resource, errors.Errorf(err, "GetResource: error %v", uuid)
	}
	return resource, nil
}
func (a *authStorage) AddResource(resource models.Resource) (models.Resource, error) {
	err := a.Conn.Save(&resource).Error
	if err != nil {
		return resource, errors.Errorf(err, "AddResource error %v", resource)
	}
	return resource, nil
}
func (a *authStorage) UpdateResource(resource models.Resource) (models.Resource, error) {

}
func (a *authStorage) DeleteResource(uuid string) error {

}

// Operation
func (a *authStorage) GetOperation(uuid string) (models.Operation, error) {
	if len(uuid) <= 10 {
		return errors.Errorf("UUID not valid:%v", uuid)
	}
	var operation models.Operation
	err := a.Conn.Where("uuid = ?", uuid).Find(&operation).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return operation, errors.Errorf(err, "GetOperation: error %v", uuid)
	}
	return operation, nil
}
func (a *authStorage) AddOperation(operation models.Operation) (models.Operation, error) {
	err := a.Conn.Save(&operation).Error
	if err != nil {
		return operation, errors.Errorf(err, "AddOperation error %v", operation)
	}
	return operation, nil
}
func (a *authStorage) UpdateOperation(operation models.Operation) (models.Operation, error) {

}
func (a *authStorage) DeleteOperation(uuid string) error {

}

// Policy
func (a *authStorage) GetPolicy(uuid string) (models.Policy, error) {
	if len(uuid) <= 10 {
		return errors.Errorf("UUID not valid:%v", uuid)
	}
	var policy models.Policy
	err := a.Conn.Where("uuid = ?", uuid).Find(&policy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return policy, errors.Errorf(err, "GetPolicy: error %v", uuid)
	}
	return policy, nil
}
func (a *authStorage) AddPolicy(policy models.Policy) (models.Policy, error) {
	err := a.Conn.Save(&policy).Error
	if err != nil {
		return policy, errors.Errorf(err, "AddPolicy error %v", policy)
	}
	return policy, nil
}
func (a *authStorage) UpdatePolicy(policy models.Policy) (models.Policy, error) {

}
func (a *authStorage) DeletePolicy(uuid string) error {

}

// Tags
func (a *authStorage) AddTagsToPolicy(uuid string, tags []string) error {

}
func (a *authStorage) RemoveTagsFromPolicy(uuid string, tag string) error {

}
func (a *authStorage) SearchForTagInPolicy(tag []string) error {

}
