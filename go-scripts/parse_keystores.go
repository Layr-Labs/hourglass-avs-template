package main

import (
	"fmt"
	"os"
	"log"
	"github.com/BurntSushi/toml"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/signing/bn254"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/signing/keystore"
)

const CONTRACTS_REGISTRY = "0x5FbDB2315678afecb367f032d93F642f64180aa3"

type BLSKeystore struct {
	Path     string `toml:"path"`
	Password string `toml:"password"`
}

type BLSConfig struct {
	Keystores []BLSKeystore `toml:"keystores"`
}

type OperatorConfig struct {
	Image        string   `toml:"image"`
	Keys         []string `toml:"keys"`
	TotalStake   string   `toml:"total_stake"`
	BLS          BLSConfig `toml:"bls"`
}

type EigenConfig struct {
	Operator OperatorConfig `toml:"operator"`
}

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <eigen.toml>", os.Args[0])
	}
	tomlPath := os.Args[1]

	var cfg EigenConfig
	if _, err := toml.DecodeFile(tomlPath, &cfg); err != nil {
		log.Fatalf("Error decoding TOML: %v", err)
	}

	keystores := cfg.Operator.BLS.Keystores
	fmt.Printf("export KEYSTORE_COUNT=%d\n", len(keystores))
	for i, ks := range keystores {


		scheme := bn254.NewScheme()
		keystoreData, err := keystore.LoadKeystoreFile(ks.Path)
		if err != nil{
			log.Printf("failed to get keystore data %s",err)
		}
		privateKeyData, err := keystoreData.GetPrivateKey(ks.Password, scheme)
		if err != nil {
			log.Printf("failed to extract key %s",err)
		}
	

		pvtKey, err := bls.NewPrivateKey(string(privateKeyData.Bytes()))

		keyPair := bls.NewKeyPair(pvtKey)
		g1_x := keyPair.GetPubKeyG1().X
		g1_y := keyPair.GetPubKeyG1().Y
		g2_x_0 := keyPair.GetPubKeyG2().X.A0
		g2_x_1 := keyPair.GetPubKeyG2().X.A1
		g2_y_0 := keyPair.GetPubKeyG2().Y.A0
		g2_y_1 := keyPair.GetPubKeyG2().Y.A1


		blsInt := privateKeyData.Bytes()
		fmt.Printf("export OPERATOR_%d_BLSPRIVATE_KEY=%s\n", i, string(blsInt))
		fmt.Printf("export OPERATOR_%d_G1_X=%s\n", i, g1_x.String())
		fmt.Printf("export OPERATOR_%d_G1_Y=%s\n", i, g1_y.String())
		fmt.Printf("export OPERATOR_%d_G2_X0=%s\n", i, g2_x_0.String())
		fmt.Printf("export OPERATOR_%d_G2_X1=%s\n", i, g2_x_1.String())
		fmt.Printf("export OPERATOR_%d_G2_Y0=%s\n", i, g2_y_0.String())
		fmt.Printf("export OPERATOR_%d_G2_Y1=%s\n", i, g2_y_1.String())
	}
}


