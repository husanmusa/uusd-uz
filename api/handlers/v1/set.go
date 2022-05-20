package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	l "github.com/husanmusa/uusd-uz/pkg/logger"
	"github.com/husanmusa/uusd-uz/pkg/structs"
	"github.com/husanmusa/uusd-uz/storage/postgres"
)

// CreateSet ...
// @Summary CreateSet
// @Description This API for creating a new set
// @Tags set
// @Accept json
// @Produce json
// @Param set request body structs.CreateSet true "SetCreateRequest"
// @Success 200 {object} structs.SetStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/sets [post]
func (h *handlerV1) CreateSet(c *gin.Context) {
	var body structs.CreateSet
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := postgres.NewSetRepo(h.db).CreateSet(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create set", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetSet ...
// @Summary GetSet
// @Description This API for getting set detail
// @Tags set
// @Accept json
// @Produce json
// @Param id path string true "SetId"
// @Success 200 {object} structs.SetStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/sets/{id} [get]
func (h *handlerV1) GetSet(c *gin.Context) {
	guid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	response, err := postgres.NewSetRepo(h.db).GetSet(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get set", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetListSets ...
// @Summary GetListSets
// @Description This API for getting list of sets
// @Tags set
// @Accept json
// @Produce json
// @Param id query string true "ServiceId"
// @Success 200 {object} []structs.SetStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/sets [get]
func (h *handlerV1) GetListSets(c *gin.Context) {
	id := c.Request.URL.Query()
	guid, err := strconv.Atoi(id["id"][0])
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	response, err := postgres.NewSetRepo(h.db).GetListSets(guid)
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

// UpdateSet ...
// @Summary UpdateSet
// @Description This API for updating set
// @Tags set
// @Accept json
// @Produce json
// @Param id path string true "SetId"
// @Param User request body structs.SetStruct true "SetUpdateRequest"
// @Success 200 {object} structs.SetStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/sets/{id} [put]
func (h *handlerV1) UpdateSet(c *gin.Context) {
	var body structs.SetStruct

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

	response, err := postgres.NewSetRepo(h.db).UpdateSet(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update set", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteSet ...
// @Summary DeleteSet
// @Description This API for deleting the set
// @Tags set
// @Accept json
// @Produce json
// @Param id path string true "SetId"
// @Success 200
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/sets/{id} [delete]
func (h *handlerV1) DeleteSet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	err = postgres.NewSetRepo(h.db).DeleteSet(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete set", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Deleted",
	})
}
