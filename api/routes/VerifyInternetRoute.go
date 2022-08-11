package api

import (
	"fmt"
	queries "isc/api/queries"
	"isc/database"
	"isc/encryption"
	internet "isc/internetconnectionchecker"
	"net/http"
	"strconv"
)

func VerifyInternetRouteHandler(writer http.ResponseWriter, request *http.Request) {
	query := queries.ParseRawQuery(request.URL.RawQuery)
	attempts, err := strconv.Atoi(query.GetQuery("attempts"))

	if err != nil {
		attempts = 2
	}

	database := database.KeyDatabase{}
	result := internet.VerifyInternetConnection(attempts)                  // possible open-closed principle architecture violation
	messageEncryptor := encryption.MessageEncryptor{KeyDatabase: database} // possible open-closed principle architecture violation
	encryptedMessage := messageEncryptor.EncryptMessage(strconv.FormatBool(result), "")
	fmt.Fprintf(writer, encryptedMessage)
}
