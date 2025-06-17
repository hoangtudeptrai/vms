package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateAssignment godoc
// @Summary      Create a new Assignment
// @Description  Takes an assignment JSON and stores in DB. Returns saved JSON.
// @Tags         Assignment
// @Accept       json
// @Produce      json
// @Param        assignment  body  model.CreateAssignment  true  "Assignment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateAssignment]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateAssignment]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateAssignment]
// @Router       /assignments [post]
// @Security     BearerAuth
func CreateAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateAssignment]()

	var dto model.CreateAssignment
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateAssignment, model.Assignment](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetAssignmentByID godoc
// @Summary      Get single assignment by id
// @Description  Returns the assignment whose ID value matches the id.
// @Tags         Assignment
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read assignment by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateAssignment]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateAssignment]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateAssignment]
// @Router       /assignments/{id} [get]
// @Security     BearerAuth
func GetAssignmentByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateAssignment]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateAssignment, model.Assignment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetAssignments godoc
// @Summary      Get all assignments
// @Description  Returns all assignments from the database.
// @Tags         Assignment
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.Assignment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Assignment]
// @Router       /assignments [get]
// @Security     BearerAuth
func GetAssignments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Assignment]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Assignment, model.Assignment]("")
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

// UpdateAssignment godoc
// @Summary      Update single assignment by id
// @Description  Updates and returns a single assignment whose ID value matches the id.
// @Tags         Assignment
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update assignment by id"
// @Param        assignment  body  model.UpdateAssignment  true  "Assignment JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateAssignment]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateAssignment]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateAssignment]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateAssignment]
// @Router       /assignments/{id} [put]
// @Security     BearerAuth
func UpdateAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateAssignment]()

	var dto model.UpdateAssignment
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateAssignment, model.Assignment](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteAssignment godoc
// @Summary      Remove single assignment by id
// @Description  Deletes a single assignment from the repository based on id.
// @Tags         Assignment
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete assignment by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Assignment]
// @Failure      500  {object}  model.JsonDTORsp[model.Assignment]
// @Router       /assignments/{id} [delete]
// @Security     BearerAuth
func DeleteAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Assignment]()

	err := reposity.DeleteItemByID[model.Assignment](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetAssignmentsByCourseID godoc
// @Summary      Get assignments by course id
// @Description  Returns all assignments for a specific course.
// @Tags         Assignment
// @Accept       json
// @Produce      json
// @Param        course_id    path    string  true  "Course ID"
// @Param        title       query   string  false "Search by title"
// @Param        description query   string  false "Search by description"
// @Success      200  {object}  model.JsonDTORsp[[]model.Assignment]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Assignment]
// @Router       /assignments/course/{course_id} [get]
// @Security     BearerAuth
func GetAssignmentsByCourseID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Assignment]()

	courseID := c.Param("course_id")
	title := c.Query("title")
	description := c.Query("description")

	// Create new query
	query := reposity.NewQuery[model.Assignment, model.Assignment]()

	// Add base condition for course assignments
	query.AddRawCondition("", "course_id = ?", courseID)

	// Add title search if provided
	if title != "" {
		query.AddConditionOfTextField("AND", "title", "LIKE", title)
	}

	// Add description search if provided
	if description != "" {
		query.AddConditionOfTextField("AND", "description", "LIKE", description)
	}

	// Execute query without paging
	dtos, total, err := query.ExecNoPaging("")
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
