# Sravz Golang Crash Course

### Install/Setup
Install golang 1.20. On Mac use brew
```
$ go version
go version go1.20.4 darwin/amd64

# Install air
$ mkdir -p $(pwd)/.bin; GOBIN=$(pwd)/.bin go install github.com/cosmtrek/air@latest 
$ GOBIN=$(pwd)/.bin; go install github.com/smartystreets/goconvey
$ ./.bin/air init

# Go mod tidy and vendor
go mod tidy
go mod vendor
```

### Run
```
$ make run
go build
./toy_poc
2023/05/04 16:29:45 Server listening at localhost:8080
2023/05/04 16:29:49 Processing symbol:  ETHBTC
{"type":"spot","base_currency":"ETH","quote_currency":"BTC","status":"working","quantity_increment":"0.0001","tick_size":"0.000001","take_rate":"0.0025","make_rate":"0.001","fee_currency":"BTC","margin_trading":true,"max_initial_leverage":"10.00"}
Update quote for:  BTC
2023/05/04 16:29:49 Processing symbol:  BTCUSDC
{"type":"spot","base_currency":"BTC","quote_currency":"USDC","status":"working","quantity_increment":"0.00001","tick_size":"0.00001","take_rate":"0.0025","make_rate":"0.001","fee_currency":"USDC","margin_trading":true,"max_initial_leverage":"10.00"}
Update quote for:  USDC
```

### Test
```
# Get all symbols
$ curl localhost:8080/currency/all
[{"type":"spot","base_currency":"ETH","quote_currency":"BTC","status":"working","quantity_increment":"0.0001","tick_size":"0.000001","take_rate":"0.0025","make_rate":"0.001","fee_currency":"BTC","margin_trading":true,"max_initial_leverage":"10.00"},{"type":"spot","base_currency":"BTC","quote_currency":"USDC","status":"working","quantity_increment":"0.00001","tick_size":"0.00001","take_rate":"0.0025","make_rate":"0.001","fee_currency":"USDC","margin_trading":true,"max_initial_leverage":"10.00"}]

# Get a specific symbol
$curl localhost:8080/currency/BTCUSDC
{"type":"spot","base_currency":"BTC","quote_currency":"USDC","status":"working","quantity_increment":"0.00001","tick_size":"0.00001","take_rate":"0.0025","make_rate":"0.001","fee_currency":"USDC","margin_trading":true,"max_initial_leverage":"10.00"}

# Get a specific symbol
sampats-mbp-2:toy_poc fd98279$ curl localhost:8080/currency/ETHBTC
{"type":"spot","base_currency":"ETH","quote_currency":"BTC","status":"working","quantity_increment":"0.0001","tick_size":"0.000001","take_rate":"0.0025","make_rate":"0.001","fee_currency":"BTC","margin_trading":true,"max_initial_leverage":"10.00"}
sampats-mbp-2:toy_poc fd98279$

# Get invalid symbol
$ curl localhost:8080/currency/TEST
"Unsupported symbol: Supported Systems: [ETHBTC BTCUSDC]"
```