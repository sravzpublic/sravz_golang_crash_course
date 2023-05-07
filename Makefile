run: 
	go mod tidy
	go mod vendor
	.bin/air -- -symbols "BTCUSDC,ETHBTC,ETHUSDC"