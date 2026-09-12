package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)


func SnipHashAlgoritm(Url string) (string, error) {
	n1, err := getRandom()
	if err != nil {
		return "",fmt.Errorf("SnipHashAlgoritm error: error generate cryptorand num 1 : %w\n",err)
	}

	n2,err := getRandom()
	if err != nil {
		return "",fmt.Errorf("SnipHashAlgoritm error: error generate cryptorand num 2 : %w\n",err)
	}

	n3,err := getRandom()
	if err != nil {
		return "",fmt.Errorf("SnipHashAlgoritm error: error generate cryptorand num 3 : %w\n",err)
	}

	n4,err := getRandom()
	if err != nil {
		return "",fmt.Errorf("SnipHashAlgoritm error: error generate cryptorand num 4 : %w\n",err)
	}

	saltedURL := fmt.Sprintf("%s%d%d%s%d%d",Url,n1,n2,"snipme",n3,n4)

	hashBytes := sha256.Sum256([]byte(saltedURL))
	hashString := hex.EncodeToString(hashBytes[:]) //64 byte

	mid := len(hashString)/2
	firstSymbols := hashString[0:4]
	midSymbols := hashString[mid:mid+2]
	lastSymbols := hashString[len(hashString) - 2: ]

	return fmt.Sprintf("%s%s%s",firstSymbols,midSymbols,lastSymbols),nil
}

func getRandom() (*big.Int, error) {
	return rand.Int(rand.Reader,big.NewInt(100))
}