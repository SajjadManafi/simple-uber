package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/SajjadManafi/simple-uber/internal/util"
	"github.com/SajjadManafi/simple-uber/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) login(ctx *gin.Context) {
	var req models.Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response := models.LoginResponse{}
	var HashedPassword string

	switch req.Type {
	case "driver":
		driver, err := server.store.GetDriverByUsername(ctx, req.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		HashedPassword = driver.HashedPassword
		response.User = mapToLoginUserResponse(models.User{}, driver)
	case "rider":
		user, err := server.store.GetUserByUsername(ctx, req.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		HashedPassword = user.HashedPassword
		response.User = mapToLoginUserResponse(user, models.Driver{})
	default:
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("type must be rider or driver")))
		return
	}

	err := util.CheckPassword(req.Password, HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	accessToken, err := server.tokenMaker.CreateToken(
		req.Username,
		req.Type,
		server.Config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response.AccessToken = accessToken
	ctx.JSON(http.StatusOK, response)

}

func mapToLoginUserResponse(user models.User, driver models.Driver) models.LoginUserResponse {
	var response models.LoginUserResponse
	if user.Username == "" && driver.Username != "" {
		response.Username = driver.Username
		response.FullName = driver.FullName
		response.Gender = driver.Gender
		response.Balance = driver.Balance
		response.Email = driver.Email
		response.JoinedAt = driver.JoinedAt
	} else if user.Username != "" && driver.Username == "" {
		response.Username = user.Username
		response.FullName = user.FullName
		response.Gender = user.Gender
		response.Balance = user.Balance
		response.Email = user.Email
		response.JoinedAt = user.JoinedAt
	}
	return response
}
