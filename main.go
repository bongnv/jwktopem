package main

import (
	"bufio"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/mendsley/gojwk"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter JWK: ")
	if scanner.Scan() {
		key, _ := gojwk.Unmarshal([]byte(scanner.Text()))
		pubKey, _ := key.DecodePublicKey()
		bytes, _ := x509.MarshalPKIXPublicKey(pubKey)
		pemdata := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: bytes,
			},
		)
		fmt.Println(string(pemdata))
	}
}
