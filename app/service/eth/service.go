package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/futuramacoder/uniswap-calculator/app/client/eth"
)

type Service struct {
	client *eth.Client
}

func NewService(client *eth.Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetERC20Decimals(ctx context.Context, address common.Address) (uint8, error) {
	return s.client.GetERC20Decimal(ctx, address)
}
