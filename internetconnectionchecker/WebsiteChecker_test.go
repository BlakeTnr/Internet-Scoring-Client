package internetchecker

import (
	"testing"
)

func TestVerificationWithGoogle(test *testing.T) {
	google := Website{"google.com"}
	respond := google.VerifyConnection()

	if !respond {
		test.FailNow()
	}
}

func TestNonExistentVerification(test *testing.T) {
	nonExistentWebsite := Website{"asdadvlkavoejlaksdnvnacvlkjajdpovfbvnadfjaeisadnfjasdoifjalkdsvnadkmcvnadjvea.com"}
	respond := nonExistentWebsite.VerifyConnection()

	if respond {
		test.FailNow()
	}
}

func TestRandomWebsite(test *testing.T) {
	var websites []Website

	numberOfRandomGeneration := 3
	for i := 0; i < numberOfRandomGeneration; i++ {
		randomWebsite := RandomWebsite()
		websites = append(websites, randomWebsite)
	}

	allEqual := true
	for i := 0; i < len(websites); i++ {
		firstWebsite := websites[0]
		website := websites[i]

		if firstWebsite.Domain != website.Domain {
			allEqual = false
		}
	}

	if allEqual {
		test.Errorf("Likely not random")
	}
}
