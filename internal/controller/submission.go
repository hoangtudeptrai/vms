package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateSubmission godoc
// @Summary      Create a new Submission
// @Description  Takes a submission JSON and stores in DB. Returns saved JSON.
// @Tags         Submission
// @Produce      json
// @Param        submission  body  model.CreateSubmission  true  "Submission JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateSubmission]
// @Router       /submissions [post]
// @Security     BearerAuth
func CreateSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateSubmission]()

	var dto model.CreateSubmission
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateSubmission, model.Submission](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// ReadSubmission godoc
// @Summary      Get single submission by id
// @Description  Returns the submission whose ID value matches the id.
// @Tags         Submission
// @Produce      json
// @Param        id  path  string  true  "Read submission by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateSubmission]
// @Router       /submissions/{id} [get]
// @Security     BearerAuth
func ReadSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateSubmission]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateSubmission, model.Submission](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// UpdateSubmission godoc
// @Summary      Update single submission by id
// @Description  Updates and returns a single submission whose ID value matches the id. New data must be passed in the body.
// @Tags         Submission
// @Produce      json
// @Param        id  path  string  true  "Update submission by id"
// @Param        submission  body  model.CreateSubmission  true  "Submission JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateSubmission]
// @Router       /submissions/{id} [put]
// @Security     BearerAuth
func UpdateSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateSubmission]()

	var dto model.CreateSubmission
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = http.StatusBadRequest
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.CreateSubmission, model.Submission](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = http.StatusInternalServerError
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteSubmission godoc
// @Summary      Remove single submission by id
// @Description  Deletes a single submission from the repository based on id.
// @Tags         Submission
// @Produce      json
// @Param        id  path  string  true  "Delete submission by id"
// @Success      204
// @Router       /submissions/{id} [delete]
// @Security     BearerAuth
func DeleteSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Submission]()

	err := reposity.DeleteItemByID[model.Submission](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}
