package account

import (
	"fmt"

	"crypto/rand"

	"github.com/iost-official/Go-IOS-Protocol/common"
)

var (
	MainAccount    Account
	GenesisAccount = map[string]int64{
		"IOST7j5ynm6UXmJqY3GJp4RchiEWxjoPvifqgQVVAUaKV9sPLpx719": 3400000000,
		"IOST5QWmFgVpXwz7vgQmPd4qUHDbexKeK523PtcMDh8geyM131niP2": 3200000000,
		"IOST7koa9yKFh4j45t97xwMBzVKrJw1q8mnixRmsdJx4JftnLo52tH": 3100000000,
		"IOST5dFQcKgmU3ZaKDgfVuX2vVCgm3W9W1GyJWyhEtgfibK6nM6Hsn": 3000000000,
		"IOST7mo771oBeAuYNp7b8VLhkP7FHxcub4HeWx3VzqsLvmNKxatQ8c": 2900000000,
		"IOST5XxHPKQ5gdvEocak4rjeDKgLh3uYumrT1iUFoR82vrHkMWNz4p": 2800000000,
		"IOST8SJYAr9ig81JPvjnyeDsPX9fbDuNGB4ECWKTedvMBN4fzZiG7F": 2600000000,
	}
	// local net
	//GenesisAccount = map[string]float64{
	//	"iWgLQj3VTPN4dZnomuJMMCggv22LFw4nAkA6bmrVsmCo":  13400000000,
	//	"281pWKbjMYGWKf2QHXUKDy4rVULbF61WGCZoi4PiKhbEk": 13200000000,
	//	"bj38rN9xdqBa4eiMi1vPjcUwdMyZmQhvYbVA6cnHyQCH":  13100000000,
	//}
)

type Account struct {
	ID     string
	Pubkey []byte
	Seckey []byte
}

func NewAccount(seckey []byte) (Account, error) {
	var m Account
	if seckey == nil {
		seckey = randomSeckey()
	}
	if len(seckey) != 32 {
		return Account{}, fmt.Errorf("seckey length error")
	}

	m.Seckey = seckey
	m.Pubkey = makePubkey(seckey)
	m.ID = GetIDByPubkey(m.Pubkey)
	return m, nil
}

func randomSeckey() []byte {
	seckey := make([]byte, 32)
	_, err := rand.Read(seckey)
	if err != nil {
		return nil
	}
	return seckey
}

func makePubkey(seckey []byte) []byte {
	return common.CalcPubkeyInSecp256k1(seckey)
}

func GetIDByPubkey(pubkey []byte) string {
	if len(pubkey) != 33 {
		panic("illegal pubkey")
	}
	return "IOST" + common.Base58Encode(append(pubkey, common.Parity(pubkey)...))
}

func GetPubkeyByID(ID string) []byte {
	b := common.Base58Decode(ID[4:])
	return b[:len(b)-4]
}
