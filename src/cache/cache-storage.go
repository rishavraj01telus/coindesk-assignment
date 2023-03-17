package cache

import (
	"coindesk/constants"
	"coindesk/models"
	"errors"
	"go.uber.org/zap"
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

func (c CacheStorage) SetPrice(crypto models.Crypto) (bool, error) {

	logger.Info("Setting value in redis ")
	usd := c.storage.SetValue(crypto.CryptoName+"_"+constants.USD_PRICE, crypto.GetPrice(constants.USD_PRICE),
		time.Duration(constants.EXPIRY)*time.Second)
	if usd != nil {
		logger.Error("unable to set usd crypto price in cache")
		return false, errors.New("unable to set usd crypto price in cache")
	}

	euro := c.storage.SetValue(crypto.CryptoName+"_"+constants.EUR_PRICE, crypto.GetPrice(constants.EUR_PRICE),
		time.Duration(constants.EXPIRY)*time.Second)

	if euro != nil {
		logger.Error("unable to set eur crypto price in cache")
		return false, errors.New("unable to set eur crypto price in cache")
	}

	logger.Info("Successfully set the value in redis")
	return true, nil
}

func (c CacheStorage) GetPrice(cryptoName string) (models.Crypto, error) {

	logger.Info("Getting cached value from redis: ", zap.String("crypto", cryptoName))
	usdRate, err := c.storage.GetValue(cryptoName + "_" + constants.USD_PRICE)

	if err != nil {
		logger.Error(err.Error())
		return models.Crypto{}, errors.New("unable to fetch usd price or value not present in cache")
	}

	eurRate, err := c.storage.GetValue(cryptoName + "_" + constants.EUR_PRICE)

	if err != nil {
		logger.Error(err.Error())
		return models.Crypto{}, errors.New("unable to fetch eur price from cache")
	}

	logger.Info("Rates fetched from cached success ")

	return models.NewCrypto(cryptoName, map[string]string{
		constants.USD_PRICE: usdRate,
		constants.EUR_PRICE: eurRate,
	}), err

}
