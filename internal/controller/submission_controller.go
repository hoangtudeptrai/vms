package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateSubmission godoc
// @Summary      Create a new submission
// @Description  Takes a submission JSON and stores in DB. Returns saved JSON.
// @Tags         Submission
// @Accept       json
// @Produce      json
// @Param        submission  body  model.CreateSubmission  true  "Submission JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateSubmission]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateSubmission]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateSubmission]
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

// GetSubmissionByID godoc
// @Summary      Get single submission by id
// @Description  Returns the submission whose ID value matches the id.
// @Tags         Submission
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read submission by id"
// @Success      200  {object}  model.JsonDTORsp[model.Submission]
// @Failure      404  {object}  model.JsonDTORsp[model.Submission]
// @Failure      500  {object}  model.JsonDTORsp[model.Submission]
// @Router       /submissions/{id} [get]
// @Security     BearerAuth
func GetSubmissionByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Submission]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.Submission, model.Submission](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetSubmissions godoc
// @Summary      Get all submissions
// @Description  Returns all submissions from the database.
// @Tags         Submission
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.Submission]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Submission]
// @Router       /submissions [get]
// @Security     BearerAuth
func GetSubmissions(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Submission]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Submission, model.Submission]("")
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

// UpdateSubmission godoc
// @Summary      Update single submission by id
// @Description  Updates and returns a single submission whose ID value matches the id.
// @Tags         Submission
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update submission by id"
// @Param        submission  body  model.UpdateSubmission  true  "Submission JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateSubmission]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateSubmission]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateSubmission]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateSubmission]
// @Router       /submissions/{id} [put]
// @Security     BearerAuth
func UpdateSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateSubmission]()

	var dto model.UpdateSubmission
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateSubmission, model.Submission](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
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
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete submission by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Submission]
// @Failure      500  {object}  model.JsonDTORsp[model.Submission]
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

// GetStudentSubmissions godoc
// @Summary      Get student's submissions
// @Description  Returns all submissions for a specific student.
// @Tags         Submission
// @Accept       json
// @Produce      json
// @Param        student_id  path  string  true  "Student ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Submission]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Submission]
// @Router       /submissions/student/{student_id} [get]
// @Security     BearerAuth
func GetStudentSubmissions(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Submission]()

	studentID := c.Param("student_id")
	filter := fmt.Sprintf("student_id = '%s' ORDER BY created_at DESC", studentID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Submission, model.Submission](filter)
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

// GetLatestSubmission godoc
// @Summary      Get latest submission for a student and assignment
// @Description  Returns the most recent submission for a specific student and assignment.
// @Tags         Submission
// @Accept       json
// @Produce      json
// @Param        student_id     path  string  true  "Student ID"
// @Param        assignment_id  path  string  true  "Assignment ID"
// @Success      200  {object}  model.JsonDTORsp[model.Submission]
// @Failure      404  {object}  model.JsonDTORsp[model.Submission]
// @Failure      500  {object}  model.JsonDTORsp[model.Submission]
// @Router       /submissions/latest/{student_id}/{assignment_id} [get]
// @Security     BearerAuth
func GetLatestSubmission(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Submission]()

	studentID := c.Param("student_id")
	assignmentID := c.Param("assignment_id")
	filter := fmt.Sprintf("student_id = '%s' AND assignment_id = '%s' ORDER BY created_at DESC LIMIT 1", studentID, assignmentID)

	dtos, _, err := reposity.ReadAllItemsIntoDTO[model.Submission, model.Submission](filter)
	if err != nil || len(dtos) == 0 {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = "No submission found"
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}

	jsonRsp.Data = dtos[0]
	c.JSON(http.StatusOK, &jsonRsp)
}
