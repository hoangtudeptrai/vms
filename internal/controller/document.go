package controllers

import (
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
// @Produce      json
// @Param        document  body  model.CreateDocument  true  "Document JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateDocument]
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

// ReadDocument godoc
// @Summary      Get single document by id
// @Description  Returns the document whose ID value matches the id.
// @Tags         Document
// @Produce      json
// @Param        id  path  string  true  "Read document by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateDocument]
// @Router       /documents/{id} [get]
// @Security     BearerAuth
func ReadDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateDocument]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateDocument, model.Document](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// UpdateDocument godoc
// @Summary      Update single document by id
// @Description  Updates and returns a single document whose ID value matches the id. New data must be passed in the body.
// @Tags         Document
// @Produce      json
// @Param        id  path  string  true  "Update document by id"
// @Param        document  body  model.CreateDocument  true  "Document JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateDocument]
// @Router       /documents/{id} [put]
// @Security     BearerAuth
func UpdateDocument(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateDocument]()

	var dto model.CreateDocument
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = http.StatusBadRequest
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.CreateDocument, model.Document](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = http.StatusInternalServerError
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
// @Produce      json
// @Param        id  path  string  true  "Delete document by id"
// @Success      204
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
