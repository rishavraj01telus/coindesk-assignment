package models

type Crypto struct {
	CryptoName string            `json:"chartName"`
	Price      map[string]string `json:"bpi"`
}

func (crypto Crypto) GetName() string {
	return crypto.CryptoName
}

func (crypto Crypto) GetPrice(cryptoType string) string {
	return crypto.Price[cryptoType]
}

func NewCrypto(cryptoName string, price map[string]string) Crypto {
	return Crypto{
		CryptoName: cryptoName,
		Price:      price,
	}
}
