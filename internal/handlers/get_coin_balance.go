package handlers

import(
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/RaiyanKhan19-41453-3/goapi/api"
	"github.com/RaiyanKhan19-41453-3/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request){
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil{
		log.Error(err)
		fmt.Println("found at decode")
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		fmt.Println("found at db")
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil{
		log.Error("nothing found")
		fmt.Println("not found")
		api.InternalErrorHandler(w)
		return
	}

	log.Error("sneaky")
	fmt.Println("sneakyPrint")
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}
	

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil{
		log.Error(err)
		fmt.Println("found error by end")
		api.InternalErrorHandler(w)
		return
	}
}
