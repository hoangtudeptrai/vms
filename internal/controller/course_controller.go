package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateCourse godoc
// @Summary      Create a new course
// @Description  Takes a course JSON and stores in DB. Returns saved JSON.
// @Tags         Course
// @Accept       json
// @Produce      json
// @Param        course  body  model.CreateCourse  true  "Course JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateCourse]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateCourse]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateCourse]
// @Router       /courses [post]
// @Security     BearerAuth
func CreateCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateCourse]()

	var dto model.CreateCourse
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateCourse, model.Course](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetCourseByID godoc
// @Summary      Get single course by id
// @Description  Returns the course whose ID value matches the id.
// @Tags         Course
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read course by id"
// @Success      200  {object}  model.JsonDTORsp[model.Course]
// @Failure      404  {object}  model.JsonDTORsp[model.Course]
// @Failure      500  {object}  model.JsonDTORsp[model.Course]
// @Router       /courses/{id} [get]
// @Security     BearerAuth
func GetCourseByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Course]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.Course, model.Course](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetCourses godoc
// @Summary      Get all courses
// @Description  Returns all courses from the database.
// @Tags         Course
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.Course]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Course]
// @Router       /courses [get]
// @Security     BearerAuth
func GetCourses(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Course]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Course, model.Course]("")
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

// UpdateCourse godoc
// @Summary      Update single course by id
// @Description  Updates and returns a single course whose ID value matches the id.
// @Tags         Course
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update course by id"
// @Param        course  body  model.UpdateCourse  true  "Course JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateCourse]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateCourse]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateCourse]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateCourse]
// @Router       /courses/{id} [put]
// @Security     BearerAuth
func UpdateCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateCourse]()

	var dto model.UpdateCourse
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateCourse, model.Course](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteCourse godoc
// @Summary      Remove single course by id
// @Description  Deletes a single course from the repository based on id.
// @Tags         Course
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete course by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Course]
// @Failure      500  {object}  model.JsonDTORsp[model.Course]
// @Router       /courses/{id} [delete]
// @Security     BearerAuth
func DeleteCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Course]()

	err := reposity.DeleteItemByID[model.Course](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetCoursesByInstructor godoc
// @Summary      Get courses by instructor id
// @Description  Returns all courses taught by a specific instructor.
// @Tags         Course
// @Accept       json
// @Produce      json
// @Param        instructor_id  path  string  true  "Instructor ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Course]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Course]
// @Router       /courses/instructor/{instructor_id} [get]
// @Security     BearerAuth
func GetCoursesByInstructor(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Course]()

	instructorID := c.Param("instructor_id")
	filter := fmt.Sprintf("instructor_id = '%s'", instructorID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Course, model.Course](filter)
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

// GetEnrolledCourses godoc
// @Summary      Get enrolled courses by student id
// @Description  Returns all courses a student is enrolled in.
// @Tags         Course
// @Accept       json
// @Produce      json
// @Param        student_id  path  string  true  "Student ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Course]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Course]
// @Router       /courses/enrolled/{student_id} [get]
// @Security     BearerAuth
func GetEnrolledCourses(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Course]()

	studentID := c.Param("student_id")
	filter := fmt.Sprintf("id IN (SELECT course_id FROM course_enrollments WHERE student_id = '%s')", studentID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Course, model.Course](filter)
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
