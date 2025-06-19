package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateCourseEnrollment godoc
// @Summary      Create a new course enrollment
// @Description  Takes a course enrollment JSON and stores in DB. Returns saved JSON.
// @Tags         CourseEnrollment
// @Accept       json
// @Produce      json
// @Param        enrollment  body  model.CreateCourseEnrollment  true  "Course Enrollment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateCourseEnrollment]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateCourseEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateCourseEnrollment]
// @Router       /enrollments [post]
// @Security     BearerAuth
func CreateCourseEnrollment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateCourseEnrollment]()

	var dto model.CreateCourseEnrollment
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateCourseEnrollment, model.CourseEnrollment](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetCourseEnrollmentByID godoc
// @Summary      Get single course enrollment by id
// @Description  Returns the course enrollment whose ID value matches the id.
// @Tags         CourseEnrollment
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read course enrollment by id"
// @Success      200  {object}  model.JsonDTORsp[model.CourseEnrollment]
// @Failure      404  {object}  model.JsonDTORsp[model.CourseEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[model.CourseEnrollment]
// @Router       /enrollments/{id} [get]
// @Security     BearerAuth
func GetCourseEnrollmentByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CourseEnrollment]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.CourseEnrollment, model.CourseEnrollment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetCourseEnrollments godoc
// @Summary      Get all course enrollments
// @Description  Returns all course enrollments from the database.
// @Tags         CourseEnrollment
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.CourseEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.CourseEnrollment]
// @Router       /enrollments [get]
// @Security     BearerAuth
func GetCourseEnrollments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.CourseEnrollment]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.CourseEnrollment, model.CourseEnrollment]("")
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

// UpdateCourseEnrollment godoc
// @Summary      Update single course enrollment by id
// @Description  Updates and returns a single course enrollment whose ID value matches the id.
// @Tags         CourseEnrollment
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update course enrollment by id"
// @Param        enrollment  body  model.UpdateCourseEnrollment  true  "Course Enrollment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateCourseEnrollment]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateCourseEnrollment]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateCourseEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateCourseEnrollment]
// @Router       /enrollments/{id} [put]
// @Security     BearerAuth
func UpdateCourseEnrollment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateCourseEnrollment]()

	var dto model.UpdateCourseEnrollment
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateCourseEnrollment, model.CourseEnrollment](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteCourseEnrollment godoc
// @Summary      Remove single course enrollment by id
// @Description  Deletes a single course enrollment from the repository based on id.
// @Tags         CourseEnrollment
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete course enrollment by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.CourseEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[model.CourseEnrollment]
// @Router       /enrollments/{id} [delete]
// @Security     BearerAuth
func DeleteCourseEnrollment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CourseEnrollment]()

	err := reposity.DeleteItemByID[model.CourseEnrollment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetStudentEnrollments godoc
// @Summary      Get student's course enrollments
// @Description  Returns all course enrollments for a specific student.
// @Tags         CourseEnrollment
// @Accept       json
// @Produce      json
// @Param        student_id  path  string  true  "Student ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.CourseEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.CourseEnrollment]
// @Router       /enrollments/student/{student_id} [get]
// @Security     BearerAuth
func GetStudentEnrollments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.CourseEnrollment]()

	studentID := c.Param("student_id")
	query := reposity.NewQuery[model.CourseEnrollment, model.CourseEnrollment]()
	query.AddConditionOfTextField("AND", "student_id", "=", studentID)

	dtos, total, err := query.ExecNoPaging("-created_at")
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

// GetCourseEnrollmentsByCourse godoc
// @Summary      Get course enrollments for a course
// @Description  Returns all enrollments for a specific course.
// @Tags         CourseEnrollment
// @Accept       json
// @Produce      json
// @Param        course_id  path  string  true  "Course ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.CourseEnrollment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.CourseEnrollment]
// @Router       /enrollments/course/{course_id} [get]
// @Security     BearerAuth
func GetCourseEnrollmentsByCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.CourseEnrollment]()

	query := reposity.NewQuery[model.CourseEnrollment, model.CourseEnrollment]()
	query.AddConditionOfTextField("AND", "course_id", "=", c.Param("course_id"))

	dtos, _, err := query.ExecNoPaging("-created_at")
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dtos
	c.JSON(http.StatusOK, &jsonRsp)
}
