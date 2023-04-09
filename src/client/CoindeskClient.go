package client

import (
	"coindesk/models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type CoindeskClient struct {
	api string
}

func NewCoindeskClient(api string) *CoindeskClient {
	return &CoindeskClient{
		api: api,
	}
}

var logger, _ = zap.NewProduction()

func (coinDeskCLient CoindeskClient) GetCurrentPrice() (models.CoinDeskResponse, error) {
	response, err := http.Get(coinDeskCLient.api)

	if err != nil || response.StatusCode != 200 {
		logger.Error("Error while fetching crypto price from CoinDesk", zap.String("error", err.Error()))
		return models.CoinDeskResponse{}, err
	}

	cryptoResponse, err := coinDeskCLient.parseResponse(response)

	return cryptoResponse, err
}

func (cd *CoindeskClient) parseResponse(response *http.Response) (models.CoinDeskResponse, error) {

	var coinDeskresponse models.CoinDeskResponse

	err := json.NewDecoder(response.Body).Decode(&coinDeskresponse)

	if err != nil {
		logger.Error("error while decoding coinDesk Response", zap.String("error", err.Error()))
		return models.CoinDeskResponse{}, err
	}

	return coinDeskresponse, err
}
