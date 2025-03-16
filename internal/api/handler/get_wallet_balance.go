package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/ccwallet/internal/api/service"
	"github.com/ngqinzhe/ccwallet/internal/model"
	"github.com/ngqinzhe/ccwallet/internal/util"
)

func (w *WalletController) GetWalletBalance(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &service.GetWalletBalanceRequest{}
		if err := c.BindJSON(&req); err != nil {
			log.Printf("[WalletController][GetWalletBalance] failed to bind json to req, err: %v", err)
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				ErrorMsg: "invalid request params",
			})
			return
		}
		log.Printf("[WalletController][GetWalletBalance] req: %v", util.SafeJsonDump(req))
		resp, err := w.WalletService.GetWalletBalance(ctx, req)
		if err != nil {
			log.Printf("[WalletController][GetWalletBalance] failed to deposit, err: %v", err)
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				ErrorMsg: fmt.Sprintf("internal server error: %v", err),
			})
			return
		}

		log.Printf("[WalletController][GetWalletBalance] success resp: %v", util.SafeJsonDump(resp))
		c.JSON(http.StatusOK, resp)
	}
}
