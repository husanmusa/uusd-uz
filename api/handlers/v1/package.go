package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	l "github.com/husanmusa/uusd-uz/pkg/logger"
	"github.com/husanmusa/uusd-uz/pkg/structs"
	"github.com/husanmusa/uusd-uz/storage/postgres"
)

// CreatePackage ...
// @Summary CreatePackage
// @Description This API for creating a new package
// @Tags package
// @Accept json
// @Produce json
// @Param package request body structs.PackageStruct true "PackageCreateRequest"
// @Success 200 {object} structs.PackageStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/packages [post]
func (h *handlerV1) CreatePackage(c *gin.Context) {
	var body structs.PackageStruct
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := postgres.NewPackageRepo(h.db).CreatePackage(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create Package", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetPackage ...
// @Summary GetPackage
// @Description This API for getting package detail
// @Tags package
// @Accept json
// @Produce json
// @Param id path string true "PackageId"
// @Success 200 {object} structs.PackageStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/packages/{id} [get]
func (h *handlerV1) GetPackage(c *gin.Context) {
	guid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	response, err := postgres.NewPackageRepo(h.db).GetPackage(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get package", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetListPackages ...
// @Summary GetListPackages
// @Description This API for getting list of packages
// @Tags package
// @Accept json
// @Produce json
// @Success 200 {object} []structs.PackageStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/packages [get]
func (h *handlerV1) GetListPackages(c *gin.Context) {
	response, err := postgres.NewPackageRepo(h.db).GetListPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list set", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sets": response,
	})
}

// UpdatePackage ...
// @Summary UpdateSetUpdatePackage
// @Description This API for updating package
// @Tags package
// @Accept json
// @Produce json
// @Param id path string true "PackageId"
// @Param User request body structs.PackageStruct true "PackageUpdateRequest"
// @Success 200 {object} structs.PackageStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/packages/{id} [put]
func (h *handlerV1) UpdatePackage(c *gin.Context) {
	var body structs.PackageStruct

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

	response, err := postgres.NewPackageRepo(h.db).UpdatePackage(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update package", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeletePackage ...
// @Summary DeletePackage
// @Description This API for deleting the package
// @Tags package
// @Accept json
// @Produce json
// @Param id path string true "PackageId"
// @Success 200
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/packages/{id} [delete]
func (h *handlerV1) DeletePackage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	err = postgres.NewPackageRepo(h.db).DeletePackage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete package", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Deleted",
	})
}
