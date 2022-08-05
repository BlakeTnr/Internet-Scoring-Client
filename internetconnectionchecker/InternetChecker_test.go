package internetchecker

import (
	"testing"
)

func TestInternetChecker(test *testing.T) {
	internetConnection := VerifyInternetConnection(2)

	if !internetConnection {
		test.FailNow()
	}
}
