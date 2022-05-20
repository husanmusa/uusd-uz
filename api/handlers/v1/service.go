package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	l "github.com/husanmusa/uusd-uz/pkg/logger"
	"github.com/husanmusa/uusd-uz/pkg/structs"
	"github.com/husanmusa/uusd-uz/storage/postgres"
)

// CreateService ...
// @Summary CreateService
// @Description This API for creating a new service
// @Tags service
// @Accept json
// @Produce json
// @Param service request body structs.CreateService true "ServiceCreateRequest"
// @Success 200 {object} structs.ServiceStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/services [post]
func (h *handlerV1) CreateService(c *gin.Context) {
	var body structs.CreateService
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := postgres.NewServiceRepo(h.db).CreateService(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create service", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetService ...
// @Summary GetService
// @Description This API for getting service detail
// @Tags service
// @Accept json
// @Produce json
// @Param id path string true "ServiceId"
// @Success 200 {object} structs.ServiceStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/services/{id} [get]
func (h *handlerV1) GetService(c *gin.Context) {
	guid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	response, err := postgres.NewServiceRepo(h.db).GetService(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get service", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetListServices ...
// @Summary ListServices
// @Description This API for getting list of services
// @Tags service
// @Accept json
// @Produce json
// @Param id query string true "CompanyId"
// @Success 200 {object} []structs.ServiceStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/services [get]
func (h *handlerV1) GetListServices(c *gin.Context) {
	id := c.Request.URL.Query()
	guid, err := strconv.Atoi(id["id"][0])
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}
	response, err := postgres.NewServiceRepo(h.db).GetListServices(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list service", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"services": response,
	})
}

// UpdateService ...
// @Summary UpdateService
// @Description This API for updating service
// @Tags service
// @Accept json
// @Produce json
// @Param id path string true "ServiceId"
// @Param User request body structs.ServiceStruct true "ServiceUpdateRequest"
// @Success 200 {object} structs.ServiceStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/services/{id} [put]
func (h *handlerV1) UpdateService(c *gin.Context) {
	var body structs.ServiceStruct

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	body.Id = id
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	response, err := postgres.NewServiceRepo(h.db).UpdateService(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update service", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteService ...
// @Summary DeleteService
// @Description This API for deleting the service
// @Tags service
// @Accept json
// @Produce json
// @Param id path string true "ServiceId"
// @Success 200
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/services/{id} [delete]
func (h *handlerV1) DeleteService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	err = postgres.NewServiceRepo(h.db).DeleteService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete service", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Deleted",
	})
}
