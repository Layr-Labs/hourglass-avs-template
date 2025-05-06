package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

func main() {
	index := os.Getenv("OPERATOR_INDEX")
	if index == "" {
		log.Fatal("OPERATOR_INDEX env var not set")
	}

	envVarName := fmt.Sprintf("OPERATOR_%s_BLSPRIVATE_KEY", index)
	blsprivKey := os.Getenv(envVarName)

	hashX := os.Getenv("PUBKEY_REGISTRATION_MESSAGE_HASH_X_POINT")
	hashY := os.Getenv("PUBKEY_REGISTRATION_MESSAGE_HASH_Y_POINT")
	log.Printf("bls_key %s",blsprivKey)
	log.Printf("hash_x %s",hashX)
	log.Printf("hash_y %s",hashY)
	// Convert hashX and hashY to *big.Int
	x := new(big.Int)
	y := new(big.Int)
	if _, ok := x.SetString(hashX, 10); !ok {
		log.Fatalf("failed to parse hashX: %s", hashX)
	}
	if _, ok := y.SetString(hashY, 10); !ok {
		log.Fatalf("failed to parse hashY: %s", hashY)
	}
	log.Printf("big_int_hash_x %s",x)
	g1_point_hash := bls.NewG1Point(x,y).G1Affine
	pvtKey, _ := bls.NewPrivateKey(blsprivKey)
	keyPair := bls.NewKeyPair(pvtKey)
	signature := keyPair.SignHashedToCurveMessage(g1_point_hash)

	fmt.Printf("export PUBKEY_REGISTRATION_SIGNATURE_X=%s\n", signature.X.String())
	fmt.Printf("export PUBKEY_REGISTRATION_SIGNATURE_Y=%s\n", signature.Y.String())
}


