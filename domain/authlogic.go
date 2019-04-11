package domain

import (
	"github.com/karuppaiah/gobodyguard/models"
)

// AuthLogic represents the entity which has connection to dependent modules
type AuthLogic struct {
	AuthRepo AuthDataStore
}

// NewAuthLogic returns a service initialized with repository.
func NewAuthLogic(authRepo AuthDataStore) AuthOpns {
	return &AuthLogic{AuthRepo: authRepo}
}

// Principal
func (a AuthLogic) GetPrincipal(uuid string) (models.Principal, error) {

}
func (a AuthLogic) AddPrincipal(principal models.Principal) (models.Principal, error) {

}
func (a AuthLogic) UpdatePrincipal(principal models.Principal) (models.Principal, error) {

}
func (a AuthLogic) DeletePrincipal(uuid string) error {

}

// Resource
func (a AuthLogic) GetResource(uuid string) (models.Resource, error) {

}
func (a AuthLogic) AddResource(resource models.Resource) (models.Resource, error) {

}
func (a AuthLogic) UpdateResource(resource models.Resource) (models.Resource, error) {

}

func (a AuthLogic) DeleteResource(uuid string) error {

}

// Operation
func (a AuthLogic) GetOperation(uuid string) (models.Operation, error) {

}
func (a AuthLogic) AddOperation(operation models.Operation) (models.Operation, error) {

}
func (a AuthLogic) UpdateOperation(operation models.Operation) (models.Operation, error) {

}
func (a AuthLogic) DeleteOperation(uuid string) error {

}

// Policy

func (a AuthLogic) GrantPolicy(principalUUID string, resourceUUID string,
	operationUUID string, tags []string, userName string) (models.Policy, error) {

}
func (a AuthLogic) DenyPolicy(principalUUID string, resourceUUID string,
	operationUUID string, tags []string, userName string) (models.Policy, error) {

}
func (a AuthLogic) IsGranted(principalUUID string, resourceUUID string,
	operationUUID string, tags string) (bool, error) {

}
func (a AuthLogic) IsDenied(principalUUID string, resourceUUID string,
	operationUUID string, tags string) (bool, error) {

}
func (a AuthLogic) DeletePolicy(uuid string) error {

}

// Tags
func (a AuthLogic) AddTagToPolicy(uuid string, tags string) error {

}
func (a AuthLogic) RemoveTagFromPolicy(uuid string, tag string) error {

}
func (a AuthLogic) SearchForTagInPolicy(tag []string) error {

}
