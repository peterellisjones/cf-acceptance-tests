package apps

import (
	"testing"

	"github.com/peterellisjones/cf-acceptance-tests/helpers"
)

func TestServices(t *testing.T) {
	helpers.GinkgoBootstrap(t, "Applications")
}
