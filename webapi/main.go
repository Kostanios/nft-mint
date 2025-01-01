package main

import (
	"github.com/gorilla/mux"
	nftmint "github.com/kostanios/nft-mint/webapi/src/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/nft-mint", nftmint.MintNFTHandler).Methods("POST")
	//r.HandleFunc("/strict-nft-mint", MintNFT2Handler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
