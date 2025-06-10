package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateAssignmentDocument godoc
// @Summary      Create a new assignment document
// @Description  Takes an assignment document JSON and stores in DB. Returns saved JSON.
// @Tags         AssignmentDocument
// @Accept       json
// @Produce      json
// @Param        document  body  model.CreateAssignmentDocument  true  "Assignment Document JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateAssignmentDocument]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateAssignmentDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateAssignmentDocument]
// @Router       /assignment-documents [post]
// @Security     BearerAuth
func CreateAssignmentDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateAssignmentDocument]()

	var dto model.CreateAssignmentDocument
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateAssignmentDocument, model.AssignmentDocument](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetAssignmentDocumentByID godoc
// @Summary      Get single assignment document by id
// @Description  Returns the assignment document whose ID value matches the id.
// @Tags         AssignmentDocument
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read assignment document by id"
// @Success      200  {object}  model.JsonDTORsp[model.AssignmentDocument]
// @Failure      404  {object}  model.JsonDTORsp[model.AssignmentDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.AssignmentDocument]
// @Router       /assignment-documents/{id} [get]
// @Security     BearerAuth
func GetAssignmentDocumentByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.AssignmentDocument]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.AssignmentDocument, model.AssignmentDocument](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetAssignmentDocuments godoc
// @Summary      Get all assignment documents
// @Description  Returns all assignment documents from the database.
// @Tags         AssignmentDocument
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.AssignmentDocument]
// @Failure      500  {object}  model.JsonDTORsp[[]model.AssignmentDocument]
// @Router       /assignment-documents [get]
// @Security     BearerAuth
func GetAssignmentDocuments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.AssignmentDocument]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.AssignmentDocument, model.AssignmentDocument]("")
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

// UpdateAssignmentDocument godoc
// @Summary      Update single assignment document by id
// @Description  Updates and returns a single assignment document whose ID value matches the id.
// @Tags         AssignmentDocument
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update assignment document by id"
// @Param        document  body  model.UpdateAssignmentDocument  true  "Assignment Document JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateAssignmentDocument]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateAssignmentDocument]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateAssignmentDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateAssignmentDocument]
// @Router       /assignment-documents/{id} [put]
// @Security     BearerAuth
func UpdateAssignmentDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateAssignmentDocument]()

	var dto model.UpdateAssignmentDocument
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateAssignmentDocument, model.AssignmentDocument](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteAssignmentDocument godoc
// @Summary      Remove single assignment document by id
// @Description  Deletes a single assignment document from the repository based on id.
// @Tags         AssignmentDocument
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete assignment document by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.AssignmentDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.AssignmentDocument]
// @Router       /assignment-documents/{id} [delete]
// @Security     BearerAuth
func DeleteAssignmentDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.AssignmentDocument]()

	err := reposity.DeleteItemByID[model.AssignmentDocument](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetAssignmentDocumentsByAssignment godoc
// @Summary      Get documents for an assignment
// @Description  Returns all documents for a specific assignment.
// @Tags         AssignmentDocument
// @Accept       json
// @Produce      json
// @Param        assignment_id  path  string  true  "Assignment ID"
// @Success      200  {object}  model.JsonDTORsp[[]model.AssignmentDocument]
// @Failure      500  {object}  model.JsonDTORsp[[]model.AssignmentDocument]
// @Router       /assignment-documents/assignment/{assignment_id} [get]
// @Security     BearerAuth
func GetAssignmentDocumentsByAssignment(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.AssignmentDocument]()

	assignmentID := c.Param("assignment_id")
	filter := fmt.Sprintf("assignment_id = '%s' ORDER BY created_at DESC", assignmentID)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.AssignmentDocument, model.AssignmentDocument](filter)
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
