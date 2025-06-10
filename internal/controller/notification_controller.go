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
// @Summary      Create a new notification
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
// @Success      200  {object}  model.JsonDTORsp[model.Notification]
// @Failure      404  {object}  model.JsonDTORsp[model.Notification]
// @Failure      500  {object}  model.JsonDTORsp[model.Notification]
// @Router       /notifications/{id} [get]
// @Security     BearerAuth
func GetNotificationByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Notification]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.Notification, model.Notification](c.Param("id"))
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
// @Success      200  {object}  model.JsonDTORsp[[]model.Notification]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Notification]
// @Router       /notifications [get]
// @Security     BearerAuth
func GetNotifications(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Notification]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Notification, model.Notification]("")
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

// GetUserNotifications godoc
// @Summary      Get user's notifications
// @Description  Returns all notifications for a specific user.
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "User ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Notification]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Notification]
// @Router       /notifications/user/{user_id} [get]
// @Security     BearerAuth
func GetUserNotifications(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Notification]()

	userID := c.Param("user_id")
	filter := fmt.Sprintf("user_id = '%s' ORDER BY created_at DESC", userID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Notification, model.Notification](filter)
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

// GetUnreadNotifications godoc
// @Summary      Get user's unread notifications
// @Description  Returns all unread notifications for a specific user.
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "User ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Notification]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Notification]
// @Router       /notifications/unread/{user_id} [get]
// @Security     BearerAuth
func GetUnreadNotifications(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Notification]()

	userID := c.Param("user_id")
	filter := fmt.Sprintf("user_id = '%s' AND is_read = false ORDER BY created_at DESC", userID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Notification, model.Notification](filter)
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

// MarkNotificationAsRead godoc
// @Summary      Mark notification as read
// @Description  Updates the is_read status of a notification to true.
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Notification ID"
// @Success      200  {object}  model.JsonDTORsp[model.Notification]
// @Failure      404  {object}  model.JsonDTORsp[model.Notification]
// @Failure      500  {object}  model.JsonDTORsp[model.Notification]
// @Router       /notifications/{id}/read [put]
// @Security     BearerAuth
func MarkNotificationAsRead(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Notification]()

	notificationID := c.Param("id")
	dto, err := reposity.UpdateItemByIDFromDTO[model.Notification, model.Notification](notificationID, model.Notification{IsRead: true})
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}
