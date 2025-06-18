package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/config"
	"github.com/hoangtu1372k2/vms/internal/model"
	"github.com/hoangtu1372k2/vms/pkg/statuscode"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser godoc
// @Summary      Create a new User
// @Description  Takes a user JSON and stores in DB. Returns saved JSON.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body  model.CreateUser  true  "User JSON"
// @Success      200  {object}  model.JsonDTORsp[model.CreateUser]
// @Failure      400  {object}  model.JsonDTORsp[model.CreateUser]
// @Failure      500  {object}  model.JsonDTORsp[model.CreateUser]
// @Router       /users [post]
// @Security     BearerAuth
func CreateUser(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.CreateUser]()

	var dto model.CreateUser
	if err := c.BindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}
	dto.Password = string(hashedPassword)
	dto, err = reposity.CreateItemFromDTO[model.CreateUser, model.User](dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusCreateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusCreated, &jsonRsp)
}

// GetUserByID godoc
// @Summary      Get single user by id
// @Description  Returns the user whose ID value matches the id.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Read user by id"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateUser]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateUser]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateUser]
// @Router       /users/{id} [get]
// @Security     BearerAuth
func GetUserByID(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateUser]()
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateUser, model.User](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusNotFound, &jsonRsp)
		return
	}
	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// Get godoc
// @Summary      Get all users
// @Description  Returns all users from the database.
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.JsonDTORsp[[]model.UpdateUser]
// @Failure      500  {object}  model.JsonDTORsp[[]model.UpdateUser]
// @Router       /users [get]
// @Security     BearerAuth
func Get(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[[]model.User]()

	role := c.Query("role")
	query := reposity.NewQuery[model.User, model.User]()

	if role != "" {
		query.AddConditionOfTextField("AND", "role", "=", role)
	}

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

// UpdateUser godoc
// @Summary      Update single user by id
// @Description  Updates and returns a single user whose ID value matches the id. New data must be passed in the body.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "Update user by id"
// @Param        user body  model.UpdateUser  true  "User JSON"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateUser]
// @Failure      400  {object}  model.JsonDTORsp[model.UpdateUser]
// @Failure      404  {object}  model.JsonDTORsp[model.UpdateUser]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateUser]
// @Router       /users/{id} [put]
// @Security     BearerAuth
func UpdateUser(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateUser]()

	var dto model.UpdateUser
	if err := c.ShouldBindJSON(&dto); err != nil {
		jsonRsp.Code = statuscode.StatusBindingInputJsonFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusBadRequest, &jsonRsp)
		return
	}

	dto, err := reposity.UpdateItemByIDFromDTO[model.UpdateUser, model.User](c.Param("id"), dto)
	if err != nil {
		jsonRsp.Code = statuscode.StatusUpdateItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}

// DeleteUser godoc
// @Summary      Remove single user by id
// @Description  Deletes a single user from the repository based on id.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Delete user by id"
// @Success      204  "No Content"
// @Failure      404  {object}  model.JsonDTORsp[model.User]
// @Failure      500  {object}  model.JsonDTORsp[model.User]
// @Router       /users/{id} [delete]
// @Security     BearerAuth
func DeleteUser(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.User]()

	err := reposity.DeleteItemByID[model.User](c.Param("id"))
	if err != nil {
		jsonRsp.Code = statuscode.StatusDeleteItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	c.JSON(http.StatusNoContent, &jsonRsp)
}

// GetCurrentUser godoc
// @Summary      Get current user information from token
// @Description  Returns the user information based on the provided token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        token  query  string  true  "JWT token"
// @Success      200  {object}  model.JsonDTORsp[model.UpdateUser]
// @Failure      401  {object}  model.JsonDTORsp[model.UpdateUser]
// @Failure      500  {object}  model.JsonDTORsp[model.UpdateUser]
// @Router       /users/me [get]
// @Security     BearerAuth
func GetCurrentUser(c *gin.Context) {
	jsonRsp := model.NewJsonDTORsp[model.UpdateUser]()

	// Get token from query parameter
	token := c.Query("token")
	if token == "" {
		jsonRsp.Code = statuscode.StatusUnauthorized
		jsonRsp.Message = "Token is required"
		c.JSON(http.StatusUnauthorized, &jsonRsp)
		return
	}

	// Parse and validate token
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT.Secret), nil
	})

	if err != nil || !parsedToken.Valid {
		jsonRsp.Code = statuscode.StatusUnauthorized
		jsonRsp.Message = "Invalid token"
		c.JSON(http.StatusUnauthorized, &jsonRsp)
		return
	}

	// Get user ID from token claims
	userID, ok := claims["id"].(string)
	if !ok {
		jsonRsp.Code = statuscode.StatusInternalServerError
		jsonRsp.Message = "Invalid user ID in token"
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	// Get user information
	dto, err := reposity.ReadItemByIDIntoDTO[model.UpdateUser, model.User](userID)
	if err != nil {
		jsonRsp.Code = statuscode.StatusReadItemFailed
		jsonRsp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, &jsonRsp)
		return
	}

	jsonRsp.Data = dto
	c.JSON(http.StatusOK, &jsonRsp)
}
