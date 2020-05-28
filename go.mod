module github.com/cosmos/peggy

go 1.13

require (
	github.com/aristanetworks/goarista v0.0.0-20190528200627-2e9fd846018e // indirect
	github.com/btcsuite/btcd v0.0.0-20190824003749-130ea5bddde3 // indirect
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cosmos/cosmos-sdk v0.37.11
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.9.8
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/huin/goupnp v1.0.0 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/miguelmota/go-solidity-sha3 v0.1.0
	github.com/onsi/ginkgo v1.9.0 // indirect
	github.com/onsi/gomega v1.6.0 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rcrowley/go-metrics v0.0.0-20190706150252-9beb055b7962 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.0
	github.com/tendermint/tm-db v0.2.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/fetchai/cosmos-sdk v0.4.1
	github.com/tendermint/tendermint => github.com/fetchai/cosmos-consensus v0.4.1
)
