package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateMessage godoc
// @Summary      Create a new message
// @Description  Takes a message JSON and stores in DB. Returns saved JSON.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        message  body  model.CreateMessage  true  "Message JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateMessage]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateMessage]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateMessage]
// @Router       /messages [post]
// @Security     BearerAuth
func CreateMessage(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateMessage]()

	var dto model.CreateMessage
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateMessage, model.Message](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetMessageByID godoc
// @Summary      Get single message by id
// @Description  Returns the message whose ID value matches the id.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read message by id"
// @Success      200  {object}  model.JsonDTORsp[model.Message]
// @Failure      404  {object}  model.JsonDTORsp[model.Message]
// @Failure      500  {object}  model.JsonDTORsp[model.Message]
// @Router       /messages/{id} [get]
// @Security     BearerAuth
func GetMessageByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Message]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.Message, model.Message](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetMessages godoc
// @Summary      Get all messages
// @Description  Returns all messages from the database.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.Message]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Message]
// @Router       /messages [get]
// @Security     BearerAuth
func GetMessages(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Message]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Message, model.Message]("")
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

// UpdateMessage godoc
// @Summary      Update single message by id
// @Description  Updates and returns a single message whose ID value matches the id.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update message by id"
// @Param        message  body  model.UpdateMessage  true  "Message JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateMessage]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateMessage]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateMessage]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateMessage]
// @Router       /messages/{id} [put]
// @Security     BearerAuth
func UpdateMessage(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateMessage]()

	var dto model.UpdateMessage
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateMessage, model.Message](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteMessage godoc
// @Summary      Remove single message by id
// @Description  Deletes a single message from the repository based on id.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete message by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Message]
// @Failure      500  {object}  model.JsonDTORsp[model.Message]
// @Router       /messages/{id} [delete]
// @Security     BearerAuth
func DeleteMessage(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Message]()

	err := reposity.DeleteItemByID[model.Message](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetUserInbox godoc
// @Summary      Get user's inbox messages
// @Description  Returns all messages where the user is the receiver.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "User ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Message]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Message]
// @Router       /messages/inbox/{user_id} [get]
// @Security     BearerAuth
func GetUserInbox(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Message]()

	userID := c.Param("user_id")
	filter := fmt.Sprintf("receiver_id = '%s' ORDER BY created_at DESC", userID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Message, model.Message](filter)
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

// GetUserSentMessages godoc
// @Summary      Get user's sent messages
// @Description  Returns all messages where the user is the sender.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "User ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Message]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Message]
// @Router       /messages/sent/{user_id} [get]
// @Security     BearerAuth
func GetUserSentMessages(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Message]()

	userID := c.Param("user_id")
	filter := fmt.Sprintf("sender_id = '%s' ORDER BY created_at DESC", userID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Message, model.Message](filter)
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

// GetMessageThread godoc
// @Summary      Get message thread
// @Description  Returns all messages in a thread based on the replied_to field.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        message_id  path  string  true  "Original Message ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Message]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Message]
// @Router       /messages/thread/{message_id} [get]
// @Security     BearerAuth
func GetMessageThread(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Message]()

	messageID := c.Param("message_id")
	filter := fmt.Sprintf("id = '%s' OR replied_to = '%s' ORDER BY created_at ASC", messageID, messageID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Message, model.Message](filter)
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
