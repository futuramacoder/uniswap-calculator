# CLI Application for getting output amount in Uniswap V2

### Commands

#### Command for getting output amount
Name: **outputAmount**

Parameters:
* inputToken - ERC20 address
* outputToken - ERC20 address
* pair - Address to pair
* inputAmount - Input amount for swap
* format - Boolean value if set to true you can view output amount with formatting to decimal

### How to run

```
COPY .env.example .env

go run cmd/main.go outputAmount -inputToken={address} -outputToken={address} -pair={address} -inputAmount={amountInWei} -format=true
```

### Example

```
❯ go run cmd/main.go outputAmount -inputToken=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2 -outputToken=0x2260fac5e5542a773aa44fbcfedf7c193bc2c599 -pair=0xbb2b8038a1640196fbe3e38816f3e67cba72d940 -inputAmount=996613480541082505 -format=true
INFO[0000] Output amount (in wei): 6367937              
INFO[0000] Output amount (to decimal): 0.06367937
```
