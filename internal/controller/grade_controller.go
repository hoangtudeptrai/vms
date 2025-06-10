package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateGrade godoc
// @Summary      Create a new grade
// @Description  Takes a grade JSON and stores in DB. Returns saved JSON.
// @Tags         Grade
// @Accept       json
// @Produce      json
// @Param        grade  body  model.CreateGrade  true  "Grade JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateGrade]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateGrade]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateGrade]
// @Router       /grades [post]
// @Security     BearerAuth
func CreateGrade(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateGrade]()

	var dto model.CreateGrade
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateGrade, model.Grade](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetGradeByID godoc
// @Summary      Get single grade by id
// @Description  Returns the grade whose ID value matches the id.
// @Tags         Grade
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read grade by id"
// @Success      200  {object}  model.JsonDTORsp[model.Grade]
// @Failure      404  {object}  model.JsonDTORsp[model.Grade]
// @Failure      500  {object}  model.JsonDTORsp[model.Grade]
// @Router       /grades/{id} [get]
// @Security     BearerAuth
func GetGradeByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Grade]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.Grade, model.Grade](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetGrades godoc
// @Summary      Get all grades
// @Description  Returns all grades from the database.
// @Tags         Grade
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.Grade]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Grade]
// @Router       /grades [get]
// @Security     BearerAuth
func GetGrades(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Grade]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Grade, model.Grade]("")
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

// UpdateGrade godoc
// @Summary      Update single grade by id
// @Description  Updates and returns a single grade whose ID value matches the id.
// @Tags         Grade
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update grade by id"
// @Param        grade  body  model.UpdateGrade  true  "Grade JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateGrade]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateGrade]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateGrade]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateGrade]
// @Router       /grades/{id} [put]
// @Security     BearerAuth
func UpdateGrade(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateGrade]()

	var dto model.UpdateGrade
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateGrade, model.Grade](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteGrade godoc
// @Summary      Remove single grade by id
// @Description  Deletes a single grade from the repository based on id.
// @Tags         Grade
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete grade by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Grade]
// @Failure      500  {object}  model.JsonDTORsp[model.Grade]
// @Router       /grades/{id} [delete]
// @Security     BearerAuth
func DeleteGrade(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Grade]()

	err := reposity.DeleteItemByID[model.Grade](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetStudentGrades godoc
// @Summary      Get student's grades
// @Description  Returns all grades for a specific student.
// @Tags         Grade
// @Accept       json
// @Produce      json
// @Param        student_id  path  string  true  "Student ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Grade]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Grade]
// @Router       /grades/student/{student_id} [get]
// @Security     BearerAuth
func GetStudentGrades(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Grade]()

	studentID := c.Param("student_id")
	filter := fmt.Sprintf("student_id = '%s' ORDER BY created_at DESC", studentID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Grade, model.Grade](filter)
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

// GetAssignmentGrades godoc
// @Summary      Get grades for an assignment
// @Description  Returns all grades for a specific assignment.
// @Tags         Grade
// @Accept       json
// @Produce      json
// @Param        assignment_id  path  string  true  "Assignment ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Grade]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Grade]
// @Router       /grades/assignment/{assignment_id} [get]
// @Security     BearerAuth
func GetAssignmentGrades(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Grade]()

	assignmentID := c.Param("assignment_id")
	filter := fmt.Sprintf("assignment_id = '%s' ORDER BY created_at DESC", assignmentID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Grade, model.Grade](filter)
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

// GetCourseGrades godoc
// @Summary      Get grades for a course
// @Description  Returns all grades for a specific course.
// @Tags         Grade
// @Accept       json
// @Produce      json
// @Param        course_id  path  string  true  "Course ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Grade]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Grade]
// @Router       /grades/course/{course_id} [get]
// @Security     BearerAuth
func GetCourseGrades(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Grade]()

	courseID := c.Param("course_id")
	filter := fmt.Sprintf("course_id = '%s' ORDER BY created_at DESC", courseID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Grade, model.Grade](filter)
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
