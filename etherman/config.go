package etherman

import "github.com/ethereum/go-ethereum/common"

// Config represents the configuration of the etherman
type Config struct {
	URL        string
	PoEAddress common.Address

	PrivateKeyPath     string
	PrivateKeyPassword string
}