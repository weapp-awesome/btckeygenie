/* gobtcaddr v1.0
 * vsergeev
 * https://github.com/vsergeev/gobtcaddr
 * MIT Licensed
 */

package main

import (
	"fmt"
	"github.com/vsergeev/gobtcaddr/btckey"
	"log"
	"os"
)

func main() {
	/* Redirect fatal errors to stderr */
	log.SetOutput(os.Stderr)

	/* Generate a new ECDSA keypair */
	priv, err := btckey.GenerateKey()
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	/* Convert the public key to a bitcoin network address */
	address := priv.ToAddress(0x00)

	/* Convert the private key to a WIF string */
	wif := priv.ToWIF()

	fmt.Println("Address:", address)
	fmt.Println("    WIF:", wif)
}
