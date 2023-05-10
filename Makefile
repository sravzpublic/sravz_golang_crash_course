run: 
	go mod tidy
	go mod vendor
	.bin/air -- -symbols "BTCUSDC,ETHBTC,ETHUSDC" -cpuprofile=havlak1.prof

build: 
	go build .