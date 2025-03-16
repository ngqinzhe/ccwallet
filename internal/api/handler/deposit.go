package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/ccwallet/internal/api/service"
	"github.com/ngqinzhe/ccwallet/internal/model"
	"github.com/ngqinzhe/ccwallet/internal/util"
)

func (w *WalletController) Deposit() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &service.DepositRequest{}
		if err := c.BindJSON(&req); err != nil {
			log.Printf("[WalletController][Deposit] failed to bind json to req, err: %v", err)
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				ErrorMsg: "invalid request params",
			})
			return
		}
		log.Printf("[WalletController][Deposit] req: %v", util.SafeJsonDump(req))
		resp, err := w.WalletService.Deposit(c.Request.Context(), req)
		if err != nil {
			log.Printf("[WalletController][Deposit] failed to deposit, err: %v", err)
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				ErrorMsg: fmt.Sprintf("internal server error: %v", err),
			})
			return
		}

		log.Printf("[WalletController][Deposit] success resp: %v", util.SafeJsonDump(resp))
		c.JSON(http.StatusOK, resp)
	}
}
