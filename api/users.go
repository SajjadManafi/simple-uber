package api

import (
	"net/http"

	"github.com/SajjadManafi/simple-uber/internal/util"
	"github.com/SajjadManafi/simple-uber/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := models.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Gender:         req.Gender,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := models.CreateUserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Gender:   user.Gender,
		Balance:  user.Balance,
		Email:    user.Email,
		JoinedAt: user.JoinedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
