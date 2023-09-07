package payment

import (
	"strconv"

	"github.com/vctrthe/api-go/config"
	"github.com/vctrthe/api-go/user"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	midtrans.ServerKey = config.C.Midtrans.ServerKey
	midtrans.ClientKey = config.C.Midtrans.ClientKey
	midtrans.Environment = midtrans.Sandbox

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
			FName: user.Name,
		},
	}

	snapTrans, err := snap.CreateTransaction(snapReq)
	if err != nil {
		return "", err
	}

	redirectURL := snapTrans.RedirectURL

	return redirectURL, nil
}
