package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
)

// CreateProfile godoc
// @Summary      Create a new profile
// @Description  Takes a profile JSON and stores in DB. Returns saved JSON.
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        profile  body  model.CreateProfile  true  "Profile JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateProfile]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateProfile]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateProfile]
// @Router       /profiles [post]
// @Security     BearerAuth
func CreateProfile(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateProfile]()

	var dto model.CreateProfile
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.CreateItemFromDTO[model.CreateProfile, model.Profile](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetProfileByID godoc
// @Summary      Get single profile by id
// @Description  Returns the profile whose ID value matches the id.
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read profile by id"
// @Success      200  {object}  model.JsonDTORsp[model.Profile]
// @Failure      404  {object}  model.JsonDTORsp[model.Profile]
// @Failure      500  {object}  model.JsonDTORsp[model.Profile]
// @Router       /profiles/{id} [get]
// @Security     BearerAuth
func GetProfileByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Profile]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.Profile, model.Profile](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// GetProfiles godoc
// @Summary      Get all profiles
// @Description  Returns all profiles from the database.
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.Profile]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Profile]
// @Router       /profiles [get]
// @Security     BearerAuth
func GetProfiles(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Profile]()

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Profile, model.Profile]("")
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

// UpdateProfile godoc
// @Summary      Update single profile by id
// @Description  Updates and returns a single profile whose ID value matches the id.
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update profile by id"
// @Param        profile  body  model.UpdateProfile  true  "Profile JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateProfile]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateProfile]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateProfile]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateProfile]
// @Router       /profiles/{id} [put]
// @Security     BearerAuth
func UpdateProfile(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateProfile]()

	var dto model.UpdateProfile
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateProfile, model.Profile](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteProfile godoc
// @Summary      Remove single profile by id
// @Description  Deletes a single profile from the repository based on id.
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete profile by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.Profile]
// @Failure      500  {object}  model.JsonDTORsp[model.Profile]
// @Router       /profiles/{id} [delete]
// @Security     BearerAuth
func DeleteProfile(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.Profile]()

	err := reposity.DeleteItemByID[model.Profile](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetProfilesByRole godoc
// @Summary      Get profiles by role
// @Description  Returns all profiles with a specific role.
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        role  path  string  true  "Role (student/instructor)"
// @Success      200  {object}  model.JsonDTORsp[[]model.Profile]
// @Failure      500  {object}  model.JsonDTORsp[[]model.Profile]
// @Router       /profiles/role/{role} [get]
// @Security     BearerAuth
func GetProfilesByRole(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.Profile]()

	role := c.Param("role")
	filter := fmt.Sprintf("role = '%s'", role)

	dtos, total, err := reposity.ReadAllItemsIntoDTO[model.Profile, model.Profile](filter)
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
