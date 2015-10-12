package bitcoin-identity

import "fmt"
import "encoding/hex"
import "crypto/sha256"

import "github.com/btcsuite/btcd/btcec"
import "golang.org/x/crypto/ripemd160"
import "github.com/btcsuite/btcutil/base58"

type Secret struct {
	PrivKey, PubKey []byte
	SIN             string
}

func GenerateKey() Secret {
	keys, _ := btcec.NewPrivateKey(btcec.S256())

	pub := keys.PubKey().SerializeCompressed()
	priv := keys.Serialize()
	sin := GenerateSin(pub)
	return Secret{pub, priv, sin}

}

func GenerateSin(pubkey []byte) string {

	// FIRST STEP - COMPLETE
	sha_256 := sha256.New()
	sha_256.Write(pubkey)

	// SECOND STEP - COMPLETE
	rip := ripemd160.New()
	rip.Write(sha_256.Sum(nil))

	// THIRD STEP - COMPLETE
	sin_version, _ := hex.DecodeString("0f02")
	pubPrefixed := append(sin_version, rip.Sum(nil)...)

	// FOURTH STEP - COMPLETE
	sha2nd := sha256.New()
	sha2nd.Write([]byte(pubPrefixed))

	sha3rd := sha256.New()
	sha3rd.Write([]byte(sha2nd.Sum(nil)))

	// // FIFTH STEP - COMPLETE
	checksum := sha3rd.Sum(nil)[0:4]

	// SIXTH STEP - COMPLETE
	pubWithChecksum := append(pubPrefixed, checksum...)

	// SIN
	sin := base58.Encode(pubWithChecksum)

	return sin
}
