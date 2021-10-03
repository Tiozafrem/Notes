package handler

import (
	"net/http"
	"notes/model"

	"github.com/gin-gonic/gin"
)

// Struct for parse json username and password
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Struct for output tokens and binding in json
type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Struct for input refresh and binding out json
type refreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// @Summary SignUp
// @Tags auth
// @Description create account, registry
// @ID registry
// @Accept json
// @Produce json
// @Param input body model.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.usecases.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @Id login
// @Accept json
// @Produce json
// @Param input body signInInput true "credentials"
// @Success 200 {object} tokenResponse "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.usecases.Authorization.GenerateToken(input.Username, input.Password, c.ClientIP())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	})
}

// @Summary Refresh
// @Tags auth
// @Description refresh token
// @Id refresh
// @Accept json
// @Produce json
// @Param input body refreshInput true "credentials"
// @Success 200 {object} tokenResponse "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/refresh [post]
func (h *Handler) refreshToken(c *gin.Context) {
	var input refreshInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokenResponse, err := h.usecases.Authorization.RefreshToken(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}
