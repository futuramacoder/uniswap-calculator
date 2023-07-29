package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/futuramacoder/uniswap-calculator/app/bindings"
)

type Client struct {
	client *ethclient.Client
}

func NewClient(client *ethclient.Client) *Client {
	return &Client{client: client}
}

func (c *Client) GetReserves(ctx context.Context, pairAddress common.Address) (*PairReserves, error) {
	pair, err := bindings.NewUniswapv2pair(pairAddress, c.client)
	if err != nil {
		return nil, err
	}
	reserves, err := pair.GetReserves(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	return &PairReserves{
		Reserve0:           reserves.Reserve0,
		Reserve1:           reserves.Reserve1,
		BlockTimestampLast: reserves.BlockTimestampLast,
	}, nil
}

func (c *Client) GetERC20Decimal(ctx context.Context, address common.Address) (uint8, error) {
	erc20, err := bindings.NewErc20Caller(address, c.client)
	if err != nil {
		return 0, err
	}
	decimal, err := erc20.Decimals(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}
	return decimal, nil
}
