package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateComment godoc
// @Summary      Create a new comment
// @Description  Takes a comment JSON and stores in DB. Returns saved JSON.
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        comment  body  model.CreateComment  true  "Comment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateComment]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateComment]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateComment]
// @Router       /comments [post]
// @Security     BearerAuth
func CreateComment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateComment]()

	var dto model.CreateComment
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateComment, model.Comment](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetCommentByID godoc
// @Summary      Get single comment by id
// @Description  Returns the comment whose ID value matches the id.
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read comment by id"
// @Success      200  {object}  model.JsonDTORsp[model.Comment]
// @Failure      404  {object}  model.JsonDTORsp[model.Comment]
// @Failure      500  {object}  model.JsonDTORsp[model.Comment]
// @Router       /comments/{id} [get]
// @Security     BearerAuth
func GetCommentByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Comment]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.Comment, model.Comment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetComments godoc
// @Summary      Get all comments
// @Description  Returns all comments from the database.
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.Comment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Comment]
// @Router       /comments [get]
// @Security     BearerAuth
func GetComments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Comment]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Comment, model.Comment]("")
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

// UpdateComment godoc
// @Summary      Update single comment by id
// @Description  Updates and returns a single comment whose ID value matches the id.
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update comment by id"
// @Param        comment  body  model.UpdateComment  true  "Comment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateComment]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateComment]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateComment]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateComment]
// @Router       /comments/{id} [put]
// @Security     BearerAuth
func UpdateComment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateComment]()

	var dto model.UpdateComment
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateComment, model.Comment](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteComment godoc
// @Summary      Remove single comment by id
// @Description  Deletes a single comment from the repository based on id.
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete comment by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Comment]
// @Failure      500  {object}  model.JsonDTORsp[model.Comment]
// @Router       /comments/{id} [delete]
// @Security     BearerAuth
func DeleteComment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Comment]()

	err := reposity.DeleteItemByID[model.Comment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetSubmissionComments godoc
// @Summary      Get comments for a submission
// @Description  Returns all comments for a specific submission.
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        submission_id  path  string  true  "Submission ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Comment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Comment]
// @Router       /comments/submission/{submission_id} [get]
// @Security     BearerAuth
func GetSubmissionComments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Comment]()

	submissionID := c.Param("submission_id")
	filter := fmt.Sprintf("submission_id = '%s' ORDER BY created_at ASC", submissionID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Comment, model.Comment](filter)
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

// GetUserComments godoc
// @Summary      Get comments by a user
// @Description  Returns all comments made by a specific user.
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "User ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Comment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Comment]
// @Router       /comments/user/{user_id} [get]
// @Security     BearerAuth
func GetUserComments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Comment]()

	userID := c.Param("user_id")
	filter := fmt.Sprintf("user_id = '%s' ORDER BY created_at DESC", userID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Comment, model.Comment](filter)
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
