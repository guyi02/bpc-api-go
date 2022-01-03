package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guyi02/bpc-api/controllers"
)

func main() {

	// walletAddressEncoded := wallet.MakeWallet("testegui")
	// walletSavedPath, err := wallet.FindWalletByAddressEncoded(walletAddressEncoded)
	// if err != nil {
	// 	fmt.Println("Cannot find this wallet")
	// }

	// walletPublicAddress := wallet.GetWalletAddress(walletSavedPath, "testegui")
	// fmt.Println(walletPublicAddress)

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/wallet/{id}", controllers.GetWalletBalance).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
