package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateEnrollment godoc
// @Summary      Create a new Enrollment
// @Description  Takes an enrollment JSON and stores in DB. Returns saved JSON.
// @Tags         Enrollment
// @Accept       json
// @Produce      json
// @Param        enrollment  body  model.CreateEnrollment  true  "Enrollment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateEnrollment]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateEnrollment]
// @Router       /enrollments [post]
// @Security     BearerAuth
func CreateEnrollment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateEnrollment]()

	var dto model.CreateEnrollment
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateEnrollment, model.Enrollment](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetEnrollmentByID godoc
// @Summary      Get single enrollment by id
// @Description  Returns the enrollment whose ID value matches the id.
// @Tags         Enrollment
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read enrollment by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateEnrollment]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateEnrollment]
// @Router       /enrollments/{id} [get]
// @Security     BearerAuth
func GetEnrollmentByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateEnrollment]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateEnrollment, model.Enrollment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetEnrollments godoc
// @Summary      Get all enrollments
// @Description  Returns all enrollments from the database.
// @Tags         Enrollment
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.UpdateEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.UpdateEnrollment]
// @Router       /enrollments [get]
// @Security     BearerAuth
func GetEnrollments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.UpdateEnrollment]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.UpdateEnrollment, model.Enrollment]("")
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.Header("X-Total-Count", fmt.Sprintf("%d", total))
	jsonRsp.Data = dtos
	c.JSON(http.StatusOK, &jsonRsp)
}

// UpdateEnrollment godoc
// @Summary      Update single enrollment by id
// @Description  Updates and returns a single enrollment whose ID value matches the id.
// @Tags         Enrollment
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update enrollment by id"
// @Param        enrollment body  model.UpdateEnrollment  true  "Enrollment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateEnrollment]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateEnrollment]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateEnrollment]
// @Router       /enrollments/{id} [put]
// @Security     BearerAuth
func UpdateEnrollment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateEnrollment]()

	var dto model.UpdateEnrollment
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateEnrollment, model.Enrollment](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteEnrollment godoc
// @Summary      Remove single enrollment by id
// @Description  Deletes a single enrollment from the repository based on id.
// @Tags         Enrollment
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete enrollment by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Enrollment]
// @Failure      500  {object}  model.JsonDTORsp[model.Enrollment]
// @Router       /enrollments/{id} [delete]
// @Security     BearerAuth
func DeleteEnrollment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Enrollment]()

	err := reposity.DeleteItemByID[model.Enrollment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}
