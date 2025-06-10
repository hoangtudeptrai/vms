package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateLesson godoc
// @Summary      Create a new lesson
// @Description  Takes a lesson JSON and stores in DB. Returns saved JSON.
// @Tags         Lesson
// @Accept       json
// @Produce      json
// @Param        lesson  body  model.CreateLesson  true  "Lesson JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateLesson]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateLesson]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateLesson]
// @Router       /lessons [post]
// @Security     BearerAuth
func CreateLesson(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateLesson]()

	var dto model.CreateLesson
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateLesson, model.Lesson](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetLessonByID godoc
// @Summary      Get single lesson by id
// @Description  Returns the lesson whose ID value matches the id.
// @Tags         Lesson
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read lesson by id"
// @Success      200  {object}  model.JsonDTORsp[model.Lesson]
// @Failure      404  {object}  model.JsonDTORsp[model.Lesson]
// @Failure      500  {object}  model.JsonDTORsp[model.Lesson]
// @Router       /lessons/{id} [get]
// @Security     BearerAuth
func GetLessonByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Lesson]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.Lesson, model.Lesson](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetLessons godoc
// @Summary      Get all lessons
// @Description  Returns all lessons from the database.
// @Tags         Lesson
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.Lesson]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Lesson]
// @Router       /lessons [get]
// @Security     BearerAuth
func GetLessons(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Lesson]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Lesson, model.Lesson]("")
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

// UpdateLesson godoc
// @Summary      Update single lesson by id
// @Description  Updates and returns a single lesson whose ID value matches the id.
// @Tags         Lesson
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update lesson by id"
// @Param        lesson  body  model.UpdateLesson  true  "Lesson JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateLesson]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateLesson]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateLesson]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateLesson]
// @Router       /lessons/{id} [put]
// @Security     BearerAuth
func UpdateLesson(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateLesson]()

	var dto model.UpdateLesson
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateLesson, model.Lesson](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteLesson godoc
// @Summary      Remove single lesson by id
// @Description  Deletes a single lesson from the repository based on id.
// @Tags         Lesson
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete lesson by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Lesson]
// @Failure      500  {object}  model.JsonDTORsp[model.Lesson]
// @Router       /lessons/{id} [delete]
// @Security     BearerAuth
func DeleteLesson(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Lesson]()

	err := reposity.DeleteItemByID[model.Lesson](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetLessonsByCourse godoc
// @Summary      Get lessons by course id
// @Description  Returns all lessons for a specific course.
// @Tags         Lesson
// @Accept       json
// @Produce      json
// @Param        course_id  path  string  true  "Course ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.Lesson]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Lesson]
// @Router       /lessons/course/{course_id} [get]
// @Security     BearerAuth
func GetLessonsByCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Lesson]()

	courseID := c.Param("course_id")
	filter := fmt.Sprintf("course_id = '%s' ORDER BY order_index ASC", courseID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Lesson, model.Lesson](filter)
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
