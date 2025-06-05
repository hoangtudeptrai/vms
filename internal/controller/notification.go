package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateNotification godoc
// @Summary      Create a new Notification
// @Description  Takes a notification JSON and stores in DB. Returns saved JSON.
// @Tags         Notification
// @Produce      json
// @Param        notification  body  model.CreateNotification  true  "Notification JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateNotification]
// @Router       /notifications [post]
// @Security     BearerAuth
func CreateNotification(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateNotification]()

	var dto model.CreateNotification
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateNotification, model.Notification](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// ReadNotification godoc
// @Summary      Get single notification by id
// @Description  Returns the notification whose ID value matches the id.
// @Tags         Notification
// @Produce      json
// @Param        id  path  string  true  "Read notification by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateNotification]
// @Router       /notifications/{id} [get]
// @Security     BearerAuth
func ReadNotification(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateNotification]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateNotification, model.Notification](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// UpdateNotification godoc
// @Summary      Update single notification by id
// @Description  Updates and returns a single notification whose ID value matches the id. New data must be passed in the body.
// @Tags         Notification
// @Produce      json
// @Param        id  path  string  true  "Update notification by id"
// @Param        notification  body  model.CreateNotification  true  "Notification JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateNotification]
// @Router       /notifications/{id} [put]
// @Security     BearerAuth
func UpdateNotification(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateNotification]()

	var dto model.CreateNotification
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = http.StatusBadRequest
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.CreateNotification, model.Notification](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = http.StatusInternalServerError
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteNotification godoc
// @Summary      Remove single notification by id
// @Description  Deletes a single notification from the repository based on id.
// @Tags         Notification
// @Produce      json
// @Param        id  path  string  true  "Delete notification by id"
// @Success      204
// @Router       /notifications/{id} [delete]
// @Security     BearerAuth
func DeleteNotification(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Notification]()

	err := reposity.DeleteItemByID[model.Notification](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}
