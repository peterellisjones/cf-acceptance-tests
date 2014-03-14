package helpers

import (
	"os"

	"github.com/peterellisjones/cf-test-helpers/cf"
)

var AdminUserContext = cf.NewUserContext(os.Getenv("API_ENDPOINT"),
	os.Getenv("ADMIN_USER"),
	os.Getenv("ADMIN_PASSWORD"),
	os.Getenv("CF_ORG"),
	os.Getenv("CF_SPACE"))

var RegularUserContext = cf.NewUserContext(os.Getenv("API_ENDPOINT"),
	os.Getenv("CF_USER"),
	os.Getenv("CF_USER_PASSWORD"),
	os.Getenv("CF_ORG"),
	os.Getenv("CF_SPACE"))

