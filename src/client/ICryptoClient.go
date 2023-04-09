package client

import (
	"coindesk/models"
)

type ICryptoClient interface {
	GetCurrentPrice() (models.CoinDeskResponse, error)
}
