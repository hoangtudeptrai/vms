package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateAssignmentSubmissionFile godoc
// @Summary      Create a new assignment submission file
// @Description  Takes an assignment submission file JSON and stores in DB. Returns saved JSON.
// @Tags         AssignmentSubmissionFile
// @Accept       json
// @Produce      json
// @Param        file  body  model.CreateAssignmentSubmissionFile  true  "Assignment Submission File JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateAssignmentSubmissionFile]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateAssignmentSubmissionFile]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateAssignmentSubmissionFile]
// @Router       /assignment-submission-files [post]
// @Security     BearerAuth
func CreateAssignmentSubmissionFile(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateAssignmentSubmissionFile]()

	var dto model.CreateAssignmentSubmissionFile
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateAssignmentSubmissionFile, model.AssignmentSubmissionFile](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetAssignmentSubmissionFileByID godoc
// @Summary      Get single assignment submission file by id
// @Description  Returns the assignment submission file whose ID value matches the id.
// @Tags         AssignmentSubmissionFile
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read assignment submission file by id"
// @Success      200  {object}  model.JsonDTORsp[model.AssignmentSubmissionFile]
// @Failure      404  {object}  model.JsonDTORsp[model.AssignmentSubmissionFile]
// @Failure      500  {object}  model.JsonDTORsp[model.AssignmentSubmissionFile]
// @Router       /assignment-submission-files/{id} [get]
// @Security     BearerAuth
func GetAssignmentSubmissionFileByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.AssignmentSubmissionFile]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.AssignmentSubmissionFile, model.AssignmentSubmissionFile](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetAssignmentSubmissionFiles godoc
// @Summary      Get all assignment submission files
// @Description  Returns all assignment submission files from the database.
// @Tags         AssignmentSubmissionFile
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.AssignmentSubmissionFile]
// @Failure      500  {object}  model.JsonDTORsp[[]model.AssignmentSubmissionFile]
// @Router       /assignment-submission-files [get]
// @Security     BearerAuth
func GetAssignmentSubmissionFiles(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.AssignmentSubmissionFile]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.AssignmentSubmissionFile, model.AssignmentSubmissionFile]("")
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

// UpdateAssignmentSubmissionFile godoc
// @Summary      Update single assignment submission file by id
// @Description  Updates and returns a single assignment submission file whose ID value matches the id.
// @Tags         AssignmentSubmissionFile
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update assignment submission file by id"
// @Param        file  body  model.UpdateAssignmentSubmissionFile  true  "Assignment Submission File JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateAssignmentSubmissionFile]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateAssignmentSubmissionFile]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateAssignmentSubmissionFile]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateAssignmentSubmissionFile]
// @Router       /assignment-submission-files/{id} [put]
// @Security     BearerAuth
func UpdateAssignmentSubmissionFile(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateAssignmentSubmissionFile]()

	var dto model.UpdateAssignmentSubmissionFile
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateAssignmentSubmissionFile, model.AssignmentSubmissionFile](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteAssignmentSubmissionFile godoc
// @Summary      Remove single assignment submission file by id
// @Description  Deletes a single assignment submission file from the repository based on id.
// @Tags         AssignmentSubmissionFile
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete assignment submission file by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.AssignmentSubmissionFile]
// @Failure      500  {object}  model.JsonDTORsp[model.AssignmentSubmissionFile]
// @Router       /assignment-submission-files/{id} [delete]
// @Security     BearerAuth
func DeleteAssignmentSubmissionFile(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.AssignmentSubmissionFile]()

	err := reposity.DeleteItemByID[model.AssignmentSubmissionFile](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetAssignmentSubmissionFilesBySubmission godoc
// @Summary      Get files for a submission
// @Description  Returns all files for a specific submission.
// @Tags         AssignmentSubmissionFile
// @Accept       json
// @Produce      json
// @Param        submission_id  path  string  true  "Submission ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.AssignmentSubmissionFile]
// @Failure      500  {object}  model.JsonDTORsp[[]model.AssignmentSubmissionFile]
// @Router       /assignment-submission-files/submission/{submission_id} [get]
// @Security     BearerAuth
func GetAssignmentSubmissionFilesBySubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.AssignmentSubmissionFile]()

	submissionID := c.Param("submission_id")
	filter := fmt.Sprintf("submission_id = '%s' ORDER BY uploaded_at DESC", submissionID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.AssignmentSubmissionFile, model.AssignmentSubmissionFile](filter)
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
