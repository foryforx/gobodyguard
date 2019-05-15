package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// AuthHandler will handle all API request to Auth service
type AuthHandler struct {
	AuthUsecase AuthOpns
}

// NewAuthHandler will create a new handler with use case and repo initialization
func NewAuthHandler(r *gin.Engine, db *gorm.DB) {
	authUsecase := NewAuthLogic(db)
	handler := &AuthHandler{
		AuthUsecase: authUsecase,
	}

	r.GET("/operation/:uuid", handler.FetchOperation)
	r.POST("/operation", handler.AddOperation)
	r.PUT("/operation/:uuid", handler.UpdateOperation)
	r.DELETE("/operation/:uuid", handler.DeleteOperation)

	r.GET("/resource/:uuid", handler.FetchResource)
	r.POST("/resource", handler.AddResource)
	r.PUT("/resource/:uuid", handler.UpdateResource)
	r.DELETE("/resource/:uuid", handler.DeleteResource)

	r.GET("/principal/:uuid", handler.FetchPrincipal)
	r.POST("/principal", handler.AddPrincipal)
	r.PUT("/principal/:uuid", handler.UpdatePrincipal)
	r.DELETE("/principal/:uuid", handler.DeletePrincipal)

	r.GET("/permission/:puuid/:ruuid/:ouuid", handler.CheckPermission)
	r.POST("/policy", handler.AddPolicy)
	r.DELETE("/policy/:uuid", handler.DeletePolicy)
}

// FetchOperation will respond with the operation data for the uuid requested
func (h *AuthHandler) FetchOperation(c *gin.Context) {
	UUID := c.Param("uuid")
	log.Debugln("UUID", UUID)
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	operation, err := h.AuthUsecase.GetOperation(UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"operation": operation})
}

// AddOperation will add a new operation data to the system
func (h *AuthHandler) AddOperation(c *gin.Context) {
	var operation Operation
	err := c.BindJSON(&operation)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	operation, err = h.AuthUsecase.AddOperation(operation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"operation": operation})
}

// UpdateOperation will update the operation data for the uuid provided
func (h *AuthHandler) UpdateOperation(c *gin.Context) {
	UUID := c.Param("uuid")
	log.Debugln("UUID", UUID)
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	var operation Operation
	err := c.BindJSON(&operation)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	operation.UUID = UUID
	operation, err = h.AuthUsecase.UpdateOperation(operation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"operation": operation})
}

// DeleteOperation will delete the operation from system for the uuid
func (h *AuthHandler) DeleteOperation(c *gin.Context) {
	UUID := c.Param("uuid")
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	err := h.AuthUsecase.DeleteOperation(UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// FetchResource will respond with the resource data for the uuid requested
func (h *AuthHandler) FetchResource(c *gin.Context) {
	UUID := c.Param("uuid")
	log.Debugln("UUID", UUID)
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	resource, err := h.AuthUsecase.GetResource(UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"resource": resource})
}

// AddResource will add a new resource data to the system
func (h *AuthHandler) AddResource(c *gin.Context) {
	var resource Resource
	err := c.BindJSON(&resource)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	resource, err = h.AuthUsecase.AddResource(resource)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"resource": resource})
}

// UpdateResource will update the resource data for the uuid provided
func (h *AuthHandler) UpdateResource(c *gin.Context) {
	UUID := c.Param("uuid")
	log.Debugln("UUID", UUID)
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	var resource Resource
	err := c.BindJSON(&resource)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	resource.UUID = UUID
	resource, err = h.AuthUsecase.UpdateResource(resource)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"resource": resource})
}

// DeleteResource will delete the resource from system for the uuid
func (h *AuthHandler) DeleteResource(c *gin.Context) {
	UUID := c.Param("uuid")
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	err := h.AuthUsecase.DeleteResource(UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// FetchPrincipal will respond with the principal data for the uuid requested
func (h *AuthHandler) FetchPrincipal(c *gin.Context) {
	UUID := c.Param("uuid")
	log.Debugln("UUID", UUID)
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	principal, err := h.AuthUsecase.GetPrincipal(UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"principal": principal})
}

// AddPrincipal will add a new principal data to the system
func (h *AuthHandler) AddPrincipal(c *gin.Context) {
	var principal Principal
	err := c.BindJSON(&principal)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	principal, err = h.AuthUsecase.AddPrincipal(principal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"principal": principal})
}

// UpdatePrincipal will update the principal data for the uuid provided
func (h *AuthHandler) UpdatePrincipal(c *gin.Context) {
	UUID := c.Param("uuid")
	log.Debugln("UUID", UUID)
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	var principal Principal
	err := c.BindJSON(&principal)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	principal.UUID = UUID
	principal, err = h.AuthUsecase.UpdatePrincipal(principal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"principal": principal})
}

// DeletePrincipal will delete the principal from system for the uuid
func (h *AuthHandler) DeletePrincipal(c *gin.Context) {
	UUID := c.Param("uuid")
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	err := h.AuthUsecase.DeletePrincipal(UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// CheckPermission will respond with the grant/deny for the principal, resource and operation requested
func (h *AuthHandler) CheckPermission(c *gin.Context) {
	principalUUID := c.Param("puuid")
	resourceUUID := c.Param("ruuid")
	operationUUID := c.Param("ouuid")
	if len(principalUUID) <= 10 || len(resourceUUID) <= 10 || len(operationUUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	permissionStatusCode, err := h.AuthUsecase.GetPermission(principalUUID, resourceUUID, operationUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"permission": permissionStatusCode.String()})
}

// AddPolicy will add a new policy data to the system
func (h *AuthHandler) AddPolicy(c *gin.Context) {
	var policy Policy
	err := c.BindJSON(&policy)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	policy, err = h.AuthUsecase.AddPermission(policy.PrincipalUUID, policy.ResourceUUID, policy.OperationUUID, "", policy.Permission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"policy": policy})
}

// DeletePolicy will delete the policy from system for the uuid
func (h *AuthHandler) DeletePolicy(c *gin.Context) {
	UUID := c.Param("uuid")
	if len(UUID) <= 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Request"})
		return
	}
	err := h.AuthUsecase.DeletePermission(UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
