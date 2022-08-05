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

	result := internet.VerifyInternetConnection(attempts)
	messageEncryptor := encryption.MessageEncryptor{database.GetPublicKey()}
	encryptedMessage := messageEncryptor.EncryptMessage(strconv.FormatBool(result), "")
	fmt.Fprintf(writer, encryptedMessage)
}
