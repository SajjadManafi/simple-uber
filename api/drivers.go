package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/SajjadManafi/simple-uber/internal/token"
	"github.com/SajjadManafi/simple-uber/internal/util"
	"github.com/SajjadManafi/simple-uber/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) createDriver(ctx *gin.Context) {
	var req models.CreateDriverRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}

	arg := models.CreateDriverParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Gender:         req.Gender,
		Email:          req.Email,
	}

	driver, err := server.store.CreateDriver(ctx, arg)
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

	response := models.CreateDriverResponse{
		Username: driver.Username,
		FullName: driver.FullName,
		Gender:   driver.Gender,
		Balance:  driver.Balance,
		Email:    driver.Email,
		JoinedAt: driver.JoinedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getDriver(ctx *gin.Context) {
	var req models.GetDriverRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	driver, err := server.store.GetDriver(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if driver.Username != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated driver")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}

	ctx.JSON(http.StatusOK, driver)
}

func (server *Server) driverWithdraw(ctx *gin.Context) {
	var req models.DriverBalanceWithdrawRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	driver, err := server.store.GetDriver(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if driver.Username != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated driver")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}

	if driver.Balance == 0 {
		ctx.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("you don't have enough balance to withdraw")))
		return
	}

	arg := models.AddDriverBalanceParams{
		ID:     req.ID,
		Amount: -driver.Balance,
	}

	_, err = server.store.AddDriverBalance(ctx, arg)

	response := models.DriverBalanceWithdrawResponse{
		Username:             driver.Username,
		FullName:             driver.FullName,
		Balance:              driver.Balance,
		BalanceAfterWithdraw: 0,
	}

	ctx.JSON(http.StatusOK, response)

}

func (server *Server) setCab(ctx *gin.Context) {
	var req models.SetCabRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	driver, err := server.store.GetDriverByUsername(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := models.UpdateDriverCurrentCabParams{
		ID:           driver.ID,
		CurrentCabID: req.CabID,
	}

	driver, err = server.store.UpdateDriverCurrentCab(ctx, arg)
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation", "foreign_key_violation":
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, driver)

}
