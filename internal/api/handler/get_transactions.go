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

func (w *WalletController) GetTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Query("user_id")
		if userId == "" {
			log.Printf("[WalletController][GetTransactions] failed to get user_id from query params")
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				ErrorMsg: "invalid request params",
			})
			return
		}
		req := &service.GetTransactionsRequest{
			UserId: userId,
		}
		log.Printf("[WalletController][GetTransactions] req: %v", util.SafeJsonDump(req))
		resp, err := w.WalletService.GetTransactions(c.Request.Context(), req)
		if err != nil {
			log.Printf("[WalletController][GetTransactions] failed to deposit, err: %v", err)
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				ErrorMsg: fmt.Sprintf("internal server error: %v", err),
			})
			return
		}

		log.Printf("[WalletController][GetTransactions] success resp: %v", util.SafeJsonDump(resp))
		c.JSON(http.StatusOK, resp)
	}
}
