package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateDocument godoc
// @Summary      Create a new Document
// @Description  Takes a document JSON and stores in DB. Returns saved JSON.
// @Tags         Document
// @Accept       json
// @Produce      json
// @Param        document  body  model.CreateDocument  true  "Document JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateDocument]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateDocument]
// @Router       /documents [post]
// @Security     BearerAuth
func CreateDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateDocument]()

	var dto model.CreateDocument
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateDocument, model.Document](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetDocumentByID godoc
// @Summary      Get single document by id
// @Description  Returns the document whose ID value matches the id.
// @Tags         Document
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read document by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateDocument]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateDocument]
// @Router       /documents/{id} [get]
// @Security     BearerAuth
func GetDocumentByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateDocument]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateDocument, model.Document](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetDocuments godoc
// @Summary      Get all documents
// @Description  Returns all documents from the database.
// @Tags         Document
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.UpdateDocument]
// @Failure      500  {object}  model.JsonDTORsp[[]model.UpdateDocument]
// @Router       /documents [get]
// @Security     BearerAuth
func GetDocuments(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.UpdateDocument]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.UpdateDocument, model.Document]("")
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

// UpdateDocument godoc
// @Summary      Update single document by id
// @Description  Updates and returns a single document whose ID value matches the id.
// @Tags         Document
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update document by id"
// @Param        document body  model.UpdateDocument  true  "Document JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateDocument]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateDocument]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateDocument]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateDocument]
// @Router       /documents/{id} [put]
// @Security     BearerAuth
func UpdateDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateDocument]()

	var dto model.UpdateDocument
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateDocument, model.Document](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteDocument godoc
// @Summary      Remove single document by id
// @Description  Deletes a single document from the repository based on id.
// @Tags         Document
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete document by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Document]
// @Failure      500  {object}  model.JsonDTORsp[model.Document]
// @Router       /documents/{id} [delete]
// @Security     BearerAuth
func DeleteDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Document]()

	err := reposity.DeleteItemByID[model.Document](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}
