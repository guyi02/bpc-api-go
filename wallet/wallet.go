package wallet

import (
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func DecodeWalletAddress(walletEncoded string) string {
	decodedWalletByte, err := hex.DecodeString(walletEncoded)
	if err != nil {
		log.Fatal(err)
	}

	return strings.ReplaceAll(string(decodedWalletByte), "0x", "")
}

// MakeWallet cria um arquivo seguro dentro do projeto contendo o endereço do user (endereço da carteira)
// e retorna uma string com a carteira criptografada
func MakeWallet(walletPassword string) string {
	newKey := keystore.NewKeyStore("./usersWallets", keystore.StandardScryptN, keystore.StandardScryptP)
	result, err := newKey.NewAccount(walletPassword)
	if err != nil {
		log.Fatal(err)
	}
	encodeAddress := hex.EncodeToString([]byte(result.Address.Hex()))
	return encodeAddress
}

// ̣ FindWalletByAddressEncoded busca carteira salva através do parâmetro id da mesma quando criada
func FindWalletByAddressEncoded(walletEncoded string) (string, error) {
	var (
		walletRef string
		notFound  error
	)

	files, err := ioutil.ReadDir("./usersWallets")
	if err != nil {
		notFound = err
	}

	walletAddress := DecodeWalletAddress(walletEncoded)

	for _, f := range files {
		if strings.Contains(f.Name(), strings.ToLower(walletAddress)) {
			walletRef = f.Name()
		}
	}

	if walletRef == "" {
		notFound = errors.New("not found")
	}

	return walletRef, notFound
}

func GetWalletAddress(walletId string, walletPassword string) string {

	walletByte, err := ioutil.ReadFile("./usersWallets/" + walletId)
	if err != nil {
		log.Fatal(err)
	}

	walletDecrypted, err := keystore.DecryptKey(walletByte, walletPassword)
	if err != nil {
		log.Fatal(err)
	}
	addresGenerate := crypto.PubkeyToAddress(walletDecrypted.PrivateKey.PublicKey).Hex()
	return addresGenerate
}
