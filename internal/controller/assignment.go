package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateAssignment godoc
// @Summary      Create a new Assignment
// @Description  Takes an assignment JSON and stores in DB. Returns saved JSON.
// @Tags         Assignment
// @Produce      json
// @Param        assignment  body  model.CreateAssignment  true  "Assignment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateAssignment]
// @Router       /assignments [post]
// @Security     BearerAuth
func CreateAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateAssignment]()

	var dto model.CreateAssignment
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateAssignment, model.Assignment](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// ReadAssignment godoc
// @Summary      Get single assignment by id
// @Description  Returns the assignment whose ID value matches the id.
// @Tags         Assignment
// @Produce      json
// @Param        id  path  string  true  "Read assignment by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateAssignment]
// @Router       /assignments/{id} [get]
// @Security     BearerAuth
func ReadAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateAssignment]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateAssignment, model.Assignment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// UpdateAssignment godoc
// @Summary      Update single assignment by id
// @Description  Updates and returns a single assignment whose ID value matches the id. New data must be passed in the body.
// @Tags         Assignment
// @Produce      json
// @Param        id  path  string  true  "Update assignment by id"
// @Param        assignment  body  model.CreateAssignment  true  "Assignment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateAssignment]
// @Router       /assignments/{id} [put]
// @Security     BearerAuth
func UpdateAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateAssignment]()

	var dto model.CreateAssignment
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = http.StatusBadRequest
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.CreateAssignment, model.Assignment](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = http.StatusInternalServerError
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteAssignment godoc
// @Summary      Remove single assignment by id
// @Description  Deletes a single assignment from the repository based on id.
// @Tags         Assignment
// @Produce      json
// @Param        id  path  string  true  "Delete assignment by id"
// @Success      204
// @Router       /assignments/{id} [delete]
// @Security     BearerAuth
func DeleteAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Assignment]()

	err := reposity.DeleteItemByID[model.Assignment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}
