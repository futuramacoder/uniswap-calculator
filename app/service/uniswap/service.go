package uniswap

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/futuramacoder/uniswap-calculator/app/client/eth"
)

type Service struct {
	client *eth.Client
}

func NewService(client *eth.Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetOutputAmount(ctx context.Context, token0, token1, poolID common.Address, amount *big.Int) (*big.Int, error) {
	reserves, err := s.client.GetReserves(ctx, poolID)
	if err != nil {
		return nil, err
	}
	sToken0, _ := s.sortTokens(token0, token1)
	if sToken0 != token0 {
		reserves.Reserve0, reserves.Reserve1 = reserves.Reserve1, reserves.Reserve0
	}
	outputAmount := s.outputAmount(amount, reserves.Reserve0, reserves.Reserve1)
	return outputAmount, nil
}

func (s *Service) outputAmount(amountIn, reserve0, reserve1 *big.Int) *big.Int {
	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, reserve1)
	denominator := new(big.Int).Mul(reserve0, big.NewInt(1000))
	denominator = new(big.Int).Add(denominator, amountInWithFee)
	return new(big.Int).Div(numerator, denominator)
}

func (s *Service) sortTokens(tkn0, tkn1 common.Address) (common.Address, common.Address) {
	token0Rep := big.NewInt(0).SetBytes(tkn0.Bytes())
	token1Rep := big.NewInt(0).SetBytes(tkn1.Bytes())

	if token0Rep.Cmp(token1Rep) > 0 {
		tkn0, tkn1 = tkn1, tkn0
	}

	return tkn0, tkn1
}

/**
go run cmd/main.go outputAmount -inputToken=0x2260fac5e5542a773aa44fbcfedf7c193bc2c599 -outputToken=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2 -pair=0xbb2b8038a1640196fbe3e38816f3e67cba72d940 -inputAmount=6366511 -format=true
*/
