package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	Models "tokenBalance/models"
	Modules "tokenBalance/modules"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// ClientHandler ethereum client instance
type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get parameter from url request
	vars := mux.Vars(r)
	module := vars["module"]

	// Get the query parameters from url request
	address := r.URL.Query().Get("address")

	// Set our response header
	w.Header().Set("Content-Type", "application/json")

	// Handle each request using the module parameter:
	switch module {

	case "get-balance":
		if address == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}

		balance, err := Modules.GetAddressBalance(*client.Client, address)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&Models.BalanceResponse{
			Address: address,
			Balance: balance,
			Symbol:  "Ether",
			Units:   "Wei",
		})

	}

}