package cache

import (
	"coindesk/models"
	"golang.org/x/net/context"
)

type ICache interface {
	SetPrice(ctx context.Context, crypto models.Crypto) (bool, error)
	GetPrice(ctx context.Context, cryptoIdentifier string) (models.Crypto, error)
}
