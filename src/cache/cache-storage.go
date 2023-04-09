package cache

import (
	"coindesk/constants"
	"coindesk/models"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"time"
)

type CacheStorage struct {
	storage *RedisClient
}

var logger, _ = zap.NewProduction()

func NewCacheStorage(client *RedisClient) CacheStorage {
	return CacheStorage{
		storage: client,
	}
}

func (c *CacheStorage) SetPrice(ctx context.Context, crypto models.Crypto) (bool, error) {

	logger.Info("Setting value in redis ")

	priceJson, _ := json.Marshal(crypto.Price)

	res := c.storage.SetValue(ctx, crypto.CryptoName, priceJson, time.Duration(constants.EXPIRY)*time.Second)

	if res != nil {
		logger.Error("unable to set crypto price in cache")
		return false, errors.New("unable to set crypto price in cache")
	}

	logger.Info("Successfully set the value in redis")
	return true, nil
}

func (c *CacheStorage) GetPrice(ctx context.Context, cryptoName string) (models.Crypto, error) {

	logger.Info("Getting cached value from redis: ", zap.String("crypto", cryptoName))
	priceJson, err := c.storage.GetValue(ctx, cryptoName)

	if err != nil {
		logger.Error(err.Error())
		return models.Crypto{}, errors.New("unable to fetch price or value not present in cache")
	}

	var jsonMap map[string]string
	json.Unmarshal([]byte(priceJson), &jsonMap)

	return models.NewCrypto(cryptoName, map[string]string{
		constants.USD_PRICE: jsonMap[constants.USD_PRICE],
		constants.EUR_PRICE: jsonMap[constants.EUR_PRICE],
	}), err

}
