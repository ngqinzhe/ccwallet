package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/ccwallet/internal/api/service"
	"github.com/ngqinzhe/ccwallet/internal/model"
	"github.com/ngqinzhe/ccwallet/internal/util"
)

func (a *AccountController) Deposit(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &service.DepositRequest{}
		log.Printf("[AccountController][Deposit] req: %v", util.SafeJsonDump(req))
		if err := c.BindJSON(&req); err != nil {
			log.Printf("[AccountController][Deposit] failed to bind json, err: %v", err)
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				ErrorMsg: "invalid request params",
			})
			return
		}

		resp, err := a.AccountService.Deposit(ctx, req)
		if err != nil {
			log.Printf("[AccountController][Deposit] failed to deposit, err: %v", err)
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				ErrorMsg: "internal server error",
			})
			return
		}
		
		log.Printf("[AccountController][Deposit] success resp: %v", util.SafeJsonDump(resp))
		c.JSON(http.StatusOK, resp)
	}
}
