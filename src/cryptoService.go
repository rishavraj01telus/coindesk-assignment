package main

import (
	"coindesk/cache"
	"coindesk/constants"
	"coindesk/models"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
)

type CryptoService struct {
	storageCache cache.ICache
}

var logger, _ = zap.NewProduction()

func NewCryptoPriceService(storageCache *cache.RedisClient) *CryptoService {
	return &CryptoService{
		storageCache: cache.NewCacheStorage(storageCache),
	}
}

func NewCryptoPriceServiceTest(storageCache cache.ICache) *CryptoService {
	return &CryptoService{
		storageCache: storageCache,
	}
}

func (cs CryptoService) CryptoPriceService() (models.Crypto, error) {
	cryptoPrice, err := cs.GetPriceFromCache(constants.BITCOIN)
	return cryptoPrice, err
}

func (cs CryptoService) GetPriceFromCache(cryptoName string) (models.Crypto, error) {

	storedCryptoPrice, err := cs.getPriceFromCacheUtil(cryptoName)
	if err == nil {
		return storedCryptoPrice, err
	}
	cryptoLivePrice, err := cs.getLiveCryptoPrice("btc")
	if err == nil {
		return cryptoLivePrice, err
	}
	logger.Error(err.Error())
	return models.Crypto{}, err
}

func (cs CryptoService) getPriceFromCacheUtil(cryptoName string) (models.Crypto, error) {

	logger.Info("Getting price from cache....")
	cachePrice, err := cs.storageCache.GetPrice(cryptoName)

	if err != nil {
		logger.Error(err.Error())
		return models.Crypto{}, err
	}

	return models.Crypto{
		CryptoName: cryptoName,
		Price: map[string]string{
			constants.USD_PRICE: cachePrice.GetPrice(constants.USD_PRICE),
			constants.EUR_PRICE: cachePrice.GetPrice(constants.EUR_PRICE),
		},
	}, nil
}

func (cs CryptoService) getLiveCryptoPrice(cryptoName string) (models.Crypto, error) {

	logger.Info("Getting live crypto price: ", zap.String("crypto", cryptoName))
	response, err := http.Get(constants.COINDESK_ENDPOINT)

	if err != nil || response.StatusCode != 200 {
		logger.Error(err.Error())
		return models.Crypto{}, err
	}

	var coinDeskresponse models.CoinDeskResponse
	var crypto models.Crypto

	err = json.NewDecoder(response.Body).Decode(&coinDeskresponse)

	if err != nil {
		logger.Error("error while decoding coinDesk Response\"")
		return models.Crypto{}, errors.New("error while decoding coinDesk Response")
	}

	price := map[string]string{
		constants.USD_PRICE: coinDeskresponse.GetPrice(constants.USD_PRICE),
		constants.EUR_PRICE: coinDeskresponse.GetPrice(constants.EUR_PRICE),
	}
	crypto = models.NewCrypto(constants.BITCOIN, price)
	cs.storageCache.SetPrice(crypto)

	//if err != nil || isValueSet == false {
	//	logger.Error(err.Error())
	//	return models.Crypto{}, errors.New("error while setting value in cache")
	//}
	return crypto, err
}
