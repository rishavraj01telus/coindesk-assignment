package service

import (
	"coindesk/cache"
	"coindesk/client"
	"coindesk/constants"
	"coindesk/models"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type CryptoService struct {
	storageCache cache.ICache
	cryptoClient client.ICryptoClient
}

var logger, _ = zap.NewProduction()

func NewCryptoPriceService(storageCache cache.ICache, cryptoClient client.ICryptoClient) *CryptoService {
	return &CryptoService{
		storageCache: storageCache,
		cryptoClient: cryptoClient,
	}
}

func NewCryptoPriceServiceTest(storageCache cache.ICache, cryptoClient client.ICryptoClient) *CryptoService {
	return &CryptoService{
		storageCache: storageCache,
		cryptoClient: cryptoClient,
	}
}

func (cs *CryptoService) CryptoPriceService(ctx context.Context) (models.Crypto, error) {
	cryptoPrice, err := cs.GetPriceFromCache(ctx, constants.BITCOIN)
	return cryptoPrice, err
}

func (cs *CryptoService) GetPriceFromCache(ctx context.Context, cryptoName string) (models.Crypto, error) {

	storedCryptoPrice, err := cs.getPriceFromCacheUtil(ctx, cryptoName)
	if err == nil {
		return storedCryptoPrice, err
	}
	cryptoLivePrice, err := cs.GetLiveCryptoPrice(ctx, cryptoName)
	if err == nil {
		return cryptoLivePrice, err
	}
	logger.Error(err.Error())
	return models.Crypto{}, err
}

func (cs *CryptoService) getPriceFromCacheUtil(ctx context.Context, cryptoName string) (models.Crypto, error) {

	logger.Info("Getting price from cache....")
	cachePrice, err := cs.storageCache.GetPrice(ctx, cryptoName)

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

func (cs *CryptoService) GetLiveCryptoPrice(ctx context.Context, cryptoName string) (models.Crypto, error) {

	logger.Info("Getting live crypto price: ", zap.String("crypto", cryptoName))
	response, err := cs.cryptoClient.GetCurrentPrice()

	if err != nil {
		logger.Error(err.Error())
		return models.Crypto{}, err
	}

	var crypto models.Crypto
	crypto = cs.parseResponse(response)
	cs.setCryptoPrice(ctx, crypto)

	return crypto, err
}

func (cs *CryptoService) setCryptoPrice(ctx context.Context, crypto models.Crypto) bool {

	isPriceSet, err := cs.storageCache.SetPrice(ctx, crypto)

	if err != nil {
		logger.Error(err.Error())
		return false
	}
	return isPriceSet
}

func (cs *CryptoService) parseResponse(coinDeskResponse models.CoinDeskResponse) models.Crypto {

	price := map[string]string{
		constants.USD_PRICE: coinDeskResponse.GetPrice(constants.USD_PRICE),
		constants.EUR_PRICE: coinDeskResponse.GetPrice(constants.EUR_PRICE),
	}
	return models.NewCrypto(constants.BITCOIN, price)

}
