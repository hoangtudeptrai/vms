package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateCourseMaterial godoc
// @Summary      Create a new Course Material
// @Description  Takes a course material JSON and stores in DB. Returns saved JSON.
// @Tags         CourseMaterial
// @Accept       json
// @Produce      json
// @Param        material  body  model.CreateCourseMaterial  true  "Course Material JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateCourseMaterial]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateCourseMaterial]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateCourseMaterial]
// @Router       /course-materials [post]
// @Security     BearerAuth
func CreateCourseMaterial(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateCourseMaterial]()

	var dto model.CreateCourseMaterial
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateCourseMaterial, model.CourseMaterial](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetCourseMaterialByID godoc
// @Summary      Get single course material by id
// @Description  Returns the course material whose ID value matches the id.
// @Tags         CourseMaterial
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read course material by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateCourseMaterial]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateCourseMaterial]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateCourseMaterial]
// @Router       /course-materials/{id} [get]
// @Security     BearerAuth
func GetCourseMaterialByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateCourseMaterial]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateCourseMaterial, model.CourseMaterial](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetCourseMaterials godoc
// @Summary      Get all course materials
// @Description  Returns all course materials from the database.
// @Tags         CourseMaterial
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.UpdateCourseMaterial]
// @Failure      500  {object}  model.JsonDTORsp[[]model.UpdateCourseMaterial]
// @Router       /course-materials [get]
// @Security     BearerAuth
func GetCourseMaterials(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.UpdateCourseMaterial]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.UpdateCourseMaterial, model.CourseMaterial]("")
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

// UpdateCourseMaterial godoc
// @Summary      Update single course material by id
// @Description  Updates and returns a single course material whose ID value matches the id.
// @Tags         CourseMaterial
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update course material by id"
// @Param        material body  model.UpdateCourseMaterial  true  "Course Material JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateCourseMaterial]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateCourseMaterial]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateCourseMaterial]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateCourseMaterial]
// @Router       /course-materials/{id} [put]
// @Security     BearerAuth
func UpdateCourseMaterial(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateCourseMaterial]()

	var dto model.UpdateCourseMaterial
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateCourseMaterial, model.CourseMaterial](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteCourseMaterial godoc
// @Summary      Remove single course material by id
// @Description  Deletes a single course material from the repository based on id.
// @Tags         CourseMaterial
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete course material by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.CourseMaterial]
// @Failure      500  {object}  model.JsonDTORsp[model.CourseMaterial]
// @Router       /course-materials/{id} [delete]
// @Security     BearerAuth
func DeleteCourseMaterial(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CourseMaterial]()

	err := reposity.DeleteItemByID[model.CourseMaterial](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}
