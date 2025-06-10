package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateAssignmentSubmission godoc
// @Summary      Create a new assignment submission
// @Description  Takes an assignment submission JSON and stores in DB. Returns saved JSON.
// @Tags         AssignmentSubmission
// @Accept       json
// @Produce      json
// @Param        submission  body  model.CreateAssignmentSubmission  true  "Assignment Submission JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateAssignmentSubmission]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateAssignmentSubmission]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateAssignmentSubmission]
// @Router       /assignment-submissions [post]
// @Security     BearerAuth
func CreateAssignmentSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateAssignmentSubmission]()

	var dto model.CreateAssignmentSubmission
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateAssignmentSubmission, model.AssignmentSubmission](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetAssignmentSubmissionByID godoc
// @Summary      Get single assignment submission by id
// @Description  Returns the assignment submission whose ID value matches the id.
// @Tags         AssignmentSubmission
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read assignment submission by id"
// @Success      200  {object}  model.JsonDTORsp[model.AssignmentSubmission]
// @Failure      404  {object}  model.JsonDTORsp[model.AssignmentSubmission]
// @Failure      500  {object}  model.JsonDTORsp[model.AssignmentSubmission]
// @Router       /assignment-submissions/{id} [get]
// @Security     BearerAuth
func GetAssignmentSubmissionByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.AssignmentSubmission]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.AssignmentSubmission, model.AssignmentSubmission](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetAssignmentSubmissions godoc
// @Summary      Get all assignment submissions
// @Description  Returns all assignment submissions from the database.
// @Tags         AssignmentSubmission
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.AssignmentSubmission]
// @Failure      500  {object}  model.JsonDTORsp[[]model.AssignmentSubmission]
// @Router       /assignment-submissions [get]
// @Security     BearerAuth
func GetAssignmentSubmissions(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.AssignmentSubmission]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.AssignmentSubmission, model.AssignmentSubmission]("")
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

// UpdateAssignmentSubmission godoc
// @Summary      Update single assignment submission by id
// @Description  Updates and returns a single assignment submission whose ID value matches the id.
// @Tags         AssignmentSubmission
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update assignment submission by id"
// @Param        submission  body  model.UpdateAssignmentSubmission  true  "Assignment Submission JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateAssignmentSubmission]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateAssignmentSubmission]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateAssignmentSubmission]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateAssignmentSubmission]
// @Router       /assignment-submissions/{id} [put]
// @Security     BearerAuth
func UpdateAssignmentSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateAssignmentSubmission]()

	var dto model.UpdateAssignmentSubmission
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateAssignmentSubmission, model.AssignmentSubmission](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteAssignmentSubmission godoc
// @Summary      Remove single assignment submission by id
// @Description  Deletes a single assignment submission from the repository based on id.
// @Tags         AssignmentSubmission
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete assignment submission by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.AssignmentSubmission]
// @Failure      500  {object}  model.JsonDTORsp[model.AssignmentSubmission]
// @Router       /assignment-submissions/{id} [delete]
// @Security     BearerAuth
func DeleteAssignmentSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.AssignmentSubmission]()

	err := reposity.DeleteItemByID[model.AssignmentSubmission](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetAssignmentSubmissionsByAssignment godoc
// @Summary      Get submissions for an assignment
// @Description  Returns all submissions for a specific assignment.
// @Tags         AssignmentSubmission
// @Accept       json
// @Produce      json
// @Param        assignment_id  path  string  true  "Assignment ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.AssignmentSubmission]
// @Failure      500  {object}  model.JsonDTORsp[[]model.AssignmentSubmission]
// @Router       /assignment-submissions/assignment/{assignment_id} [get]
// @Security     BearerAuth
func GetAssignmentSubmissionsByAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.AssignmentSubmission]()

	assignmentID := c.Param("assignment_id")
	filter := fmt.Sprintf("assignment_id = '%s' ORDER BY submitted_at DESC", assignmentID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.AssignmentSubmission, model.AssignmentSubmission](filter)
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

// GetAssignmentSubmissionsByStudent godoc
// @Summary      Get submissions for a student
// @Description  Returns all submissions for a specific student.
// @Tags         AssignmentSubmission
// @Accept       json
// @Produce      json
// @Param        student_id  path  string  true  "Student ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.AssignmentSubmission]
// @Failure      500  {object}  model.JsonDTORsp[[]model.AssignmentSubmission]
// @Router       /assignment-submissions/student/{student_id} [get]
// @Security     BearerAuth
func GetAssignmentSubmissionsByStudent(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.AssignmentSubmission]()

	studentID := c.Param("student_id")
	filter := fmt.Sprintf("student_id = '%s' ORDER BY submitted_at DESC", studentID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.AssignmentSubmission, model.AssignmentSubmission](filter)
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
