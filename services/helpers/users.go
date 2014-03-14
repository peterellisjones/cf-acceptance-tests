package helpers

import (
	. "github.com/peterellisjones/cf-acceptance-tests/helpers"
	. "github.com/onsi/gomega"
	. "github.com/peterellisjones/cf-test-helpers/cf"
	. "github.com/vito/cmdtest/matchers"
)

func LoginAsAdmin() {
	Expect(Cf("login", "-u", AdminUserContext.Username, "-p", AdminUserContext.Password, "-o", AdminUserContext.Org, "-s", AdminUserContext.Space)).To(ExitWith(0))
}

func LoginAsUser() {
	Expect(Cf("login", "-u", RegularUserContext.Username, "-p", RegularUserContext.Password, "-o", RegularUserContext.Org, "-s", RegularUserContext.Space)).To(ExitWith(0))
}
