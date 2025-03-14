package handler

import "github.com/ngqinzhe/ccwallet/internal/api/service"

type AccountController struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) *AccountController {
	return &AccountController{
		AccountService: accountService,
	}
}
