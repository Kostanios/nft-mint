package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	nftmint "github.com/kostanios/nft-mint/webapi/src/routes"
	services "github.com/kostanios/nft-mint/webapi/src/services/blockchain"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/nft-mint", nftmint.MintNFTHandler).Methods("POST")
	r.HandleFunc("/strict-nft-mint", nftmint.StrictMintNFTHandler).Methods("POST")

	port, portExists := os.LookupEnv("PORT")
	if !portExists {
		port = "8080"
	}

	go services.ListenToMintEvents()

	_, chainURLExists := os.LookupEnv("CHAIN_API_URL")
	if !chainURLExists {
		log.Fatalf("ChainURL is incorrect")
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
