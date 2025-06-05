package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateCourse godoc
// @Summary      Create a new Course
// @Description  Takes a course JSON and stores in DB. Returns saved JSON.
// @Tags         Course
// @Produce      json
// @Param        course  body  model.CreateCourse  true  "Course JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateCourse]
// @Router       /courses [post]
// @Security     BearerAuth
func CreateCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateCourse]()

	var dto model.CreateCourse
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateCourse, model.Course](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// ReadCourse godoc
// @Summary      Get single course by id
// @Description  Returns the course whose ID value matches the id.
// @Tags         Course
// @Produce      json
// @Param        id  path  string  true  "Read course by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateCourse]
// @Router       /courses/{id} [get]
// @Security     BearerAuth
func ReadCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateCourse]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateCourse, model.Course](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// UpdateCourse godoc
// @Summary      Update single course by id
// @Description  Updates and returns a single course whose ID value matches the id. New data must be passed in the body.
// @Tags         Course
// @Produce      json
// @Param        id  path  string  true  "Update course by id"
// @Param        course  body  model.CreateCourse  true  "Course JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateCourse]
// @Router       /courses/{id} [put]
// @Security     BearerAuth
func UpdateCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateCourse]()

	var dto model.CreateCourse
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = http.StatusBadRequest
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.CreateCourse, model.Course](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = http.StatusInternalServerError
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteCourse godoc
// @Summary      Remove single course by id
// @Description  Deletes a single course from the repository based on id.
// @Tags         Course
// @Produce      json
// @Param        id  path  string  true  "Delete course by id"
// @Success      204
// @Router       /courses/{id} [delete]
// @Security     BearerAuth
func DeleteCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Course]()

	err := reposity.DeleteItemByID[model.Course](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}
