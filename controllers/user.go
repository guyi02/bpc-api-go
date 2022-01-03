package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/guyi02/bpc-api/wallet"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	WalletID string `json:"walletID"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	walletAddressEncoded := wallet.MakeWallet(user.Password)
	walletSavedPath, err := wallet.FindWalletByAddressEncoded(walletAddressEncoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	walletPublicAddress := wallet.GetWalletAddress(walletSavedPath, user.Password)
	user.WalletID = walletPublicAddress
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
