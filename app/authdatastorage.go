package app

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type authStorage struct {
	Conn *gorm.DB
}

// NewAuthRepository To create new Repository with connection to DB
func NewAuthRepository(conn *gorm.DB) AuthDataStore {
	return &authStorage{conn}
}

func (a *authStorage) GetConnection() *gorm.DB {
	return a.Conn
}

// GetPrincipal will get a principal data from datastore for a particular uuid
func (a *authStorage) GetPrincipal(UUID string) (Principal, error) {
	var principal Principal
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
func (a *authStorage) AddPrincipal(principal Principal) (Principal, error) {
	principal.DeletedAt = nil
	err := a.Conn.Save(&principal).Error
	if err != nil {
		return principal, errors.Wrapf(err, "AddPrincipal error %v", principal)
	}
	return principal, nil
}

// UpdatePrincipal will update the respective principal data in datastore
func (a *authStorage) UpdatePrincipal(principal Principal) (Principal, error) {
	err := a.Conn.Update(&principal).Error
	if err != nil {
		return principal, errors.Wrapf(err, "UpdatePrincipal error %v", principal)
	}
	return principal, nil
}

// DeletePrincipal will remove the principal data from data store
func (a *authStorage) DeletePrincipal(UUID string) error {
	err := a.Conn.Unscoped().Where("uuid::text = ?", UUID).Delete(&Principal{}).Error
	return err
}

// GetResource will get the resource data from data store for the UUID
func (a *authStorage) GetResource(UUID string) (Resource, error) {
	var resource Resource
	if len(UUID) <= 10 {
		return resource, errors.Errorf("UUID not valid:%v", UUID)
	}
	err := a.Conn.Where("uuid = ?", UUID).Find(&resource).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return resource, errors.Wrapf(err, "GetResource: error %v", UUID)
	}
	return resource, nil
}

// AddResource will add a new resource to data store
func (a *authStorage) AddResource(resource Resource) (Resource, error) {
	resource.DeletedAt = nil
	err := a.Conn.Debug().Save(&resource).Error
	if err != nil {
		return resource, errors.Wrapf(err, "AddResource error %v", resource)
	}
	return resource, nil
}

// UpdateResource will update the existing resource in data store
func (a *authStorage) UpdateResource(resource Resource) (Resource, error) {
	err := a.Conn.Update(&resource).Error
	if err != nil {
		return resource, errors.Wrapf(err, "UpdateResource error %v", resource)
	}
	return resource, nil
}

// DeleteResource will delete the existing resource from data store
func (a *authStorage) DeleteResource(UUID string) error {
	err := a.Conn.Unscoped().Where("uuid::text = ?", UUID).Delete(&Resource{}).Error
	return err
}

// GetOperation will the operation data from data store for UUID
func (a *authStorage) GetOperation(UUID string) (Operation, error) {
	var operation Operation
	if len(UUID) <= 10 {
		return operation, errors.Errorf("UUID not valid:%v", UUID)
	}
	err := a.Conn.Debug().Where("uuid = ?", UUID).Find(&operation).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return operation, errors.Wrapf(err, "GetOperation: error %v", UUID)
	}
	return operation, nil
}

// AddOperation will add a new operation to data store
func (a *authStorage) AddOperation(operation Operation) (Operation, error) {
	operation.DeletedAt = nil
	err := a.Conn.Debug().Save(&operation).Error
	if err != nil {
		return operation, errors.Wrapf(err, "AddOperation error %v", operation)
	}
	return operation, nil
}

// UpdateOperation will update the existing operation from data store
func (a *authStorage) UpdateOperation(operation Operation) (Operation, error) {
	err := a.Conn.Update(&operation).Error
	if err != nil {
		return operation, errors.Wrapf(err, "UpdateOperation error %v", operation)
	}
	return operation, nil
}

// DeleteOperation will delete the operation from data store
func (a *authStorage) DeleteOperation(UUID string) error {
	err := a.Conn.Unscoped().Where("uuid::text = ?", UUID).Delete(&Operation{}).Error
	return err
}

// GetPolicy will get the policy data from data store for UUID
func (a *authStorage) GetPolicy(UUID string) (Policy, error) {
	var policy Policy
	if len(UUID) <= 10 {
		return policy, errors.Errorf("UUID not valid:%v", UUID)
	}
	err := a.Conn.Where("uuid = ?", UUID).Find(&policy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return policy, errors.Wrapf(err, "GetPolicy: error %v", UUID)
	}
	return policy, nil
}

// GetPolicyForAllMatch will get the policy from data store which matches principal, operation and resource UUID
func (a *authStorage) GetPolicyForAllMatch(principalUUID string, resourceUUID string,
	operationUUID string) ([]Policy, error) {
	var policys = make([]Policy, 0)
	if len(principalUUID) <= 10 || len(resourceUUID) <= 10 || len(operationUUID) <= 10 {
		return policys, errors.Errorf("UUID's not valid:%v %v %v", principalUUID, resourceUUID, operationUUID)
	}
	err := a.Conn.Where("principal_uuid = ? AND resource_uuid = ? AND operation_UUID = ?",
		principalUUID,
		resourceUUID,
		operationUUID).Find(&policys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return policys, errors.Wrapf(err,
			"GetPolicy: error finding policy with %v %v %v",
			principalUUID,
			resourceUUID,
			operationUUID,
		)
	}
	return policys, nil
}

// AddPolicy will add a new policy to data store
func (a *authStorage) AddPolicy(policy Policy) (Policy, error) {
	err := a.Conn.Debug().Save(&policy).Error
	if err != nil {
		return policy, errors.Wrapf(err, "AddPolicy error %v", policy)
	}
	return policy, nil
}

// UpdatePolicy will update the existing policy from data store
func (a *authStorage) UpdatePolicy(policy Policy) (Policy, error) {
	err := a.Conn.Update(&policy).Error
	if err != nil {
		return policy, errors.Wrapf(err, "UpdatePolicy error %v", policy)
	}
	return policy, nil
}

// DeletePolicy will delete the existing policy from data store
func (a *authStorage) DeletePolicy(UUID string) error {
	err := a.Conn.Unscoped().Where("uuid::text = ?", UUID).Delete(&Policy{}).Error
	return err
}

func (a *authStorage) LoadAccess() (map[string](map[string][]string), error) {
	var operationAccess = make(map[string](map[string][]string))
	type returnGroup struct {
		PrincipalUUID string
		ResourceUUID  string
		OperationUUID string
	}
	var access = make(map[string][]string)
	var rgs []returnGroup
	err := a.Conn.Debug().Table("policies").
		Select("principal_uuid, resource_uuid, operation_uuid").
		Where("permission = 'Granted'").Find(&rgs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return operationAccess, errors.Wrapf(err, "LoadAccess failed during fetch access")
	}
	if err == gorm.ErrRecordNotFound {
		log.Warnln("LoadAccess no record found")
	}
	log.Println("length of rgs", len(rgs))
	for _, rg := range rgs {
		var resources []string
		resources, ok := access[rg.PrincipalUUID]
		if !ok {
			resources = make([]string, 0)
			access[rg.PrincipalUUID] = resources
		}
		access[rg.PrincipalUUID] = append(access[rg.PrincipalUUID], rg.ResourceUUID)
		var accessPrincipal map[string][]string
		accessPrincipal, ok = operationAccess[rg.OperationUUID]
		if !ok {
			accessPrincipal = make(map[string][]string)
			operationAccess[rg.OperationUUID] = accessPrincipal
		}
		operationAccess[rg.OperationUUID][rg.PrincipalUUID] = resources
	}
	return operationAccess, nil
}
