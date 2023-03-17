package cache

import "coindesk/models"

type ICache interface {
	SetPrice(crypto models.Crypto) (bool, error)
	GetPrice(cryptoIdentifier string) (models.Crypto, error)
}
