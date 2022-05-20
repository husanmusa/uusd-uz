package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	l "github.com/husanmusa/uusd-uz/pkg/logger"
	"github.com/husanmusa/uusd-uz/pkg/structs"
	"github.com/husanmusa/uusd-uz/storage/postgres"
)

// CreateCompany ...
// @Summary CreateCompany
// @Description This API for creating a new company
// @Tags company
// @Accept json
// @Produce json
// @Param company request body structs.CreateCompany true "CompanyCreateRequest"
// @Success 200 {object} structs.CompanyStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/companies [post]
func (h *handlerV1) CreateCompany(c *gin.Context) {
	var body structs.CreateCompany
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := postgres.NewCompanyRepo(h.db).CreateCompany(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create company", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetCompany ...
// @Summary GetCompany
// @Description This API for getting company detail
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "CompanyId"
// @Success 200 {object} structs.CompanyStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/companies/{id} [get]
func (h *handlerV1) GetCompany(c *gin.Context) {
	guid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	response, err := postgres.NewCompanyRepo(h.db).GetCompany(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get company", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetListCompanies ...
// @Summary ListCompanies
// @Description This API for getting list of companies
// @Tags company
// @Accept json
// @Produce json
// @Success 200 {object} []structs.CompanyStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/companies [get]
func (h *handlerV1) GetListCompanies(c *gin.Context) {
	response, err := postgres.NewCompanyRepo(h.db).GetListCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list company", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"companies": response,
	})
}

// UpdateCompany ...
// @Summary UpdateCompany
// @Description This API for updating company
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "CompanyId"
// @Param User request body structs.CompanyStruct true "CompanyUpdateRequest"
// @Success 200 {object} structs.CompanyStruct
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/companies/{id} [put]
func (h *handlerV1) UpdateCompany(c *gin.Context) {
	var body structs.CompanyStruct

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

	response, err := postgres.NewCompanyRepo(h.db).UpdateCompany(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update company", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteCompany ...
// @Summary DeleteCompany
// @Description This API for deleting the company
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "CompanyId"
// @Success 200
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/companies/{id} [delete]
func (h *handlerV1) DeleteCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Failed to parse string to int", l.Error(err))
	}

	err = postgres.NewCompanyRepo(h.db).DeleteCompany(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete company", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Deleted",
	})
}
