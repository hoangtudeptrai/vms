package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateGrade godoc
// @Summary      Create a new Grade
// @Description  Takes a grade JSON and stores in DB. Returns saved JSON.
// @Tags         Grade
// @Produce      json
// @Param        grade  body  model.CreateGrade  true  "Grade JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateGrade]
// @Router       /grades [post]
// @Security     BearerAuth
func CreateGrade(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateGrade]()

	var dto model.CreateGrade
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateGrade, model.Grade](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// ReadGrade godoc
// @Summary      Get single grade by id
// @Description  Returns the grade whose ID value matches the id.
// @Tags         Grade
// @Produce      json
// @Param        id  path  string  true  "Read grade by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateGrade]
// @Router       /grades/{id} [get]
// @Security     BearerAuth
func ReadGrade(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateGrade]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateGrade, model.Grade](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// UpdateGrade godoc
// @Summary      Update single grade by id
// @Description  Updates and returns a single grade whose ID value matches the id. New data must be passed in the body.
// @Tags         Grade
// @Produce      json
// @Param        id  path  string  true  "Update grade by id"
// @Param        grade  body  model.CreateGrade  true  "Grade JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateGrade]
// @Router       /grades/{id} [put]
// @Security     BearerAuth
func UpdateGrade(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateGrade]()

	var dto model.CreateGrade
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = http.StatusBadRequest
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.CreateGrade, model.Grade](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = http.StatusInternalServerError
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteGrade godoc
// @Summary      Remove single grade by id
// @Description  Deletes a single grade from the repository based on id.
// @Tags         Grade
// @Produce      json
// @Param        id  path  string  true  "Delete grade by id"
// @Success      204
// @Router       /grades/{id} [delete]
// @Security     BearerAuth
func DeleteGrade(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Grade]()

	err := reposity.DeleteItemByID[model.Grade](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}
