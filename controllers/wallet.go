package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/guyi02/bpc-api/utils"
)

type WalletBalance struct {
	Balance *big.Int `json:"balance"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	strTest := "teste de requisição"
	json.NewEncoder(w).Encode(strTest)
}

func GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var res WalletBalance
	ctx := context.Background()
	conn, err := ethclient.DialContext(ctx, utils.GetEnv("TEST_NET_BINANCE"))
	if err != nil {
		log.Fatal("Whoops something went wrong!", err)
	}

	result, err := conn.BalanceAt(ctx, common.HexToAddress(params["id"]), nil)
	if err != nil {
		fmt.Println("Cannot find this wallet balance")
	}
	res.Balance = result

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
