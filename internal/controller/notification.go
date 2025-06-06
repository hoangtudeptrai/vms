package controllers

import (
	"fmt"
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
// @Accept       json
// @Produce      json
// @Param        notification  body  model.CreateNotification  true  "Notification JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateNotification]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateNotification]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateNotification]
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

// GetNotificationByID godoc
// @Summary      Get single notification by id
// @Description  Returns the notification whose ID value matches the id.
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read notification by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateNotification]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateNotification]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateNotification]
// @Router       /notifications/{id} [get]
// @Security     BearerAuth
func GetNotificationByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateNotification]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateNotification, model.Notification](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetNotifications godoc
// @Summary      Get all notifications
// @Description  Returns all notifications from the database.
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.UpdateNotification]
// @Failure      500  {object}  model.JsonDTORsp[[]model.UpdateNotification]
// @Router       /notifications [get]
// @Security     BearerAuth
func GetNotifications(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.UpdateNotification]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.UpdateNotification, model.Notification]("")
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

// UpdateNotification godoc
// @Summary      Update single notification by id
// @Description  Updates and returns a single notification whose ID value matches the id.
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update notification by id"
// @Param        notification body  model.UpdateNotification  true  "Notification JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateNotification]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateNotification]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateNotification]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateNotification]
// @Router       /notifications/{id} [put]
// @Security     BearerAuth
func UpdateNotification(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateNotification]()

	var dto model.UpdateNotification
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateNotification, model.Notification](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
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
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete notification by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Notification]
// @Failure      500  {object}  model.JsonDTORsp[model.Notification]
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
