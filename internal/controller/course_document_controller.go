package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateCourseDocument godoc
// @Summary      Create a new course document
// @Description  Takes a course document JSON and stores in DB. Returns saved JSON.
// @Tags         CourseDocument
// @Accept       json
// @Produce      json
// @Param        document  body  model.CreateCourseDocument  true  "Course Document JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateCourseDocument]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateCourseDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateCourseDocument]
// @Router       /course-documents [post]
// @Security     BearerAuth
func CreateCourseDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateCourseDocument]()

	var dto model.CreateCourseDocument
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateCourseDocument, model.CourseDocument](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetCourseDocumentByID godoc
// @Summary      Get single course document by id
// @Description  Returns the course document whose ID value matches the id.
// @Tags         CourseDocument
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read course document by id"
// @Success      200  {object}  model.JsonDTORsp[model.CourseDocument]
// @Failure      404  {object}  model.JsonDTORsp[model.CourseDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.CourseDocument]
// @Router       /course-documents/{id} [get]
// @Security     BearerAuth
func GetCourseDocumentByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CourseDocument]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.CourseDocument, model.CourseDocument](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetCourseDocuments godoc
// @Summary      Get all course documents
// @Description  Returns all course documents from the database.
// @Tags         CourseDocument
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.CourseDocument]
// @Failure      500  {object}  model.JsonDTORsp[[]model.CourseDocument]
// @Router       /course-documents [get]
// @Security     BearerAuth
func GetCourseDocuments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.CourseDocument]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.CourseDocument, model.CourseDocument]("")
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

// UpdateCourseDocument godoc
// @Summary      Update single course document by id
// @Description  Updates and returns a single course document whose ID value matches the id.
// @Tags         CourseDocument
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update course document by id"
// @Param        document  body  model.UpdateCourseDocument  true  "Course Document JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateCourseDocument]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateCourseDocument]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateCourseDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateCourseDocument]
// @Router       /course-documents/{id} [put]
// @Security     BearerAuth
func UpdateCourseDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateCourseDocument]()

	var dto model.UpdateCourseDocument
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateCourseDocument, model.CourseDocument](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteCourseDocument godoc
// @Summary      Remove single course document by id
// @Description  Deletes a single course document from the repository based on id.
// @Tags         CourseDocument
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete course document by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.CourseDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.CourseDocument]
// @Router       /course-documents/{id} [delete]
// @Security     BearerAuth
func DeleteCourseDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CourseDocument]()

	err := reposity.DeleteItemByID[model.CourseDocument](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetCourseDocumentsByCourse godoc
// @Summary      Get documents for a course
// @Description  Returns all documents for a specific course.
// @Tags         CourseDocument
// @Accept       json
// @Produce      json
// @Param        course_id  path  string  true  "Course ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.CourseDocument]
// @Failure      500  {object}  model.JsonDTORsp[[]model.CourseDocument]
// @Router       /course-documents/course/{course_id} [get]
// @Security     BearerAuth
func GetCourseDocumentsByCourse(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.CourseDocument]()

	courseID := c.Param("course_id")
	filter := fmt.Sprintf("course_id = '%s' ORDER BY created_at DESC", courseID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.CourseDocument, model.CourseDocument](filter)
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
