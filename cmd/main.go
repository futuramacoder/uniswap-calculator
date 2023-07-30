package main

import (
	"context"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"

	"github.com/futuramacoder/uniswap-calculator/app/client/eth"
	"github.com/futuramacoder/uniswap-calculator/app/config"
	ethService "github.com/futuramacoder/uniswap-calculator/app/service/eth"
	"github.com/futuramacoder/uniswap-calculator/app/service/uniswap"
	"github.com/futuramacoder/uniswap-calculator/app/utils"
)

func init() {
	_ = godotenv.Load()
}
func main() {
	envConfig, err := config.LoadEnvConfig()
	if err != nil {
		log.WithError(err).Fatal("parse config")
	}
	ctx := context.Background()
	ethClient, err := ethclient.DialContext(ctx, envConfig.EthNodeUrl)
	if err != nil {
		log.WithError(err).Fatalf("failed to init ethereum client")
	}
	client := eth.NewClient(ethClient)
	uniswapService := uniswap.NewService(client)
	ethService := ethService.NewService(client)

	app := &cli.App{
		Name:        envConfig.AppName,
		Description: "Test application for calculation output amount",
		Commands: []*cli.Command{
			{
				Name:  "outputAmount",
				Usage: "Get output amount for uniswap V2",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "inputToken",
					},
					&cli.StringFlag{
						Name: "outputToken",
					},
					&cli.StringFlag{
						Name: "pair",
					},
					&cli.Int64Flag{
						Name: "inputAmount",
					},
					&cli.BoolFlag{
						Name:     "format",
						Required: false,
					},
				},
				Action: func(c *cli.Context) error {
					token0 := common.HexToAddress(c.String("inputToken"))
					token1 := common.HexToAddress(c.String("outputToken"))
					pair := common.HexToAddress(c.String("pair"))
					amount := big.NewInt(c.Int64("inputAmount"))

					outputAmount, err := uniswapService.GetOutputAmount(ctx, token0, token1, pair, amount)
					if err != nil {
						return err
					}
					log.Infof("Output amount (in wei): %s", outputAmount.String())
					if c.Bool("format") {
						token1Decimal, err := ethService.GetERC20Decimals(c.Context, token1)
						if err != nil {
							return err
						}
						
						log.Infof("Output amount (to decimal): %s", utils.ToDecimal(outputAmount, int(token1Decimal)))
					}
					return nil
				},
			},
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
