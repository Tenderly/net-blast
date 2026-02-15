default: shim op-geth

op-geth:
	go build -o op-geth "github.com/tenderly/net-blast/blast-geth/cmd/geth"
.PHONY: op-geth

shim: main.go
	go build -o shim .
