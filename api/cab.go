package api

import (
	"net/http"

	"github.com/SajjadManafi/simple-uber/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) createCab(ctx *gin.Context) {
	var req models.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := models.CreateCabParams{
		DriverID: req.DriverID,
		Brand:    req.Brand,
		Model:    req.Model,
		Color:    req.Color,
		Plate:    req.Plate,
	}
	cab, err := server.store.CreateCab(ctx, arg)
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation", "foreign_key_violation":
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, cab)

}
