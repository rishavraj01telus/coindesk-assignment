package main

import (
	"coindesk/constants"
	mock "coindesk/mocks"
	"coindesk/models"
	"coindesk/service"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"testing"
)

func TestCryptoService(t *testing.T) {

	ctx := context.Background()
	t.Run("Get price from cache", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockInterface := mock.NewMockICache(mockCtrl)
		mockCryptoClientInterface := mock.NewMockICryptoClient(mockCtrl)
		mockService := service.NewCryptoPriceServiceTest(mockInterface, mockCryptoClientInterface)

		mockCryptoPrice := map[string]string{
			"USD": "10",
			"EUR": "11",
		}

		response := &models.Crypto{
			CryptoName: constants.BITCOIN,
			Price:      mockCryptoPrice,
		}

		mockInterface.EXPECT().GetPrice(ctx, constants.BITCOIN).Return(models.NewCrypto(
			constants.BITCOIN, mockCryptoPrice), nil)

		got, err := mockService.GetPriceFromCache(ctx, constants.BITCOIN)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("Get live crypto price ", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockInterface := mock.NewMockICache(mockCtrl)
		mockCryptoClientInterface := mock.NewMockICryptoClient(mockCtrl)
		mockService := service.NewCryptoPriceServiceTest(mockInterface, mockCryptoClientInterface)

		mockCryptoPrice := map[string]string{
			"USD": "10",
			"EUR": "11",
		}

		response := &models.Crypto{
			CryptoName: constants.BITCOIN,
			Price:      mockCryptoPrice,
		}

		mockInterface.EXPECT().GetPrice(ctx, constants.BITCOIN).Return(models.NewCrypto(
			constants.BITCOIN, mockCryptoPrice), nil)

		mockInterface.EXPECT().SetPrice(ctx, gomock.Any()).Times(1).Return(false, nil)

		got, err := mockService.GetLiveCryptoPrice(ctx, constants.BITCOIN)

		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("Get live crypto price ", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockInterface := mock.NewMockICache(mockCtrl)
		mockCryptoClientInterface := mock.NewMockICryptoClient(mockCtrl)
		mockService := service.NewCryptoPriceServiceTest(mockInterface, mockCryptoClientInterface)

		mockInterface.EXPECT().GetPrice(ctx, constants.BITCOIN).Return(models.Crypto{}, fmt.Errorf("Throwing error"))

		mockInterface.EXPECT().SetPrice(ctx, gomock.Any()).Times(1).Return(false, nil)

		got, err := mockService.GetLiveCryptoPrice(ctx, constants.BITCOIN)

		assert.Equal(t, models.Crypto{}, got)
		assert.Equal(t, err, fmt.Errorf("Throwing error"))
	})

}
