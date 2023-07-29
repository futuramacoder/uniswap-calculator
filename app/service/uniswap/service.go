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
	outputAmount := s.quote(amount, reserves.Reserve0, reserves.Reserve1)
	return outputAmount, nil
}

func (s *Service) quote(amount, reserve0, reserve1 *big.Int) *big.Int {
	if reserve1.Cmp(big.NewInt(0)) <= 0 ||
		reserve0.Cmp(big.NewInt(0)) <= 0 ||
		amount.Cmp(big.NewInt(0)) <= 0 {

		return new(big.Int)
	}

	multiplied := new(big.Int).Mul(amount, reserve1)
	res := new(big.Int).Div(multiplied, reserve0)
	return res
}

func (s *Service) sortTokens(tkn0, tkn1 common.Address) (common.Address, common.Address) {
	token0Rep := big.NewInt(0).SetBytes(tkn0.Bytes())
	token1Rep := big.NewInt(0).SetBytes(tkn1.Bytes())

	if token0Rep.Cmp(token1Rep) > 0 {
		tkn0, tkn1 = tkn1, tkn0
	}

	return tkn0, tkn1
}
