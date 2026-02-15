module claim

go 1.21

toolchain go1.21.1

require github.com/tenderly/net-blast/blast-optimism v0.0.0

require (
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
)

replace github.com/tenderly/net-blast/blast-optimism v0.0.0 => ../../..
