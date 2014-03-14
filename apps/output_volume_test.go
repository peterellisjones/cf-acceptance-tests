package apps

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/vito/cmdtest/matchers"

	. "github.com/peterellisjones/cf-acceptance-tests/helpers"
	. "github.com/peterellisjones/cf-test-helpers/cf"
	. "github.com/peterellisjones/cf-test-helpers/generator"
	. "github.com/peterellisjones/cf-test-helpers/runner"
)

var _ = Describe("An application printing a bunch of output", func() {
	var appName string

	BeforeEach(func() {
		appName = RandomName()

		Expect(Cf("push", appName, "-p", NewAssets().Dora)).To(Say("App started"))
	})

	AfterEach(func() {
		Expect(Cf("delete", appName, "-f")).To(Say("OK"))
	})

	It("doesn't die when printing 32MB", func() {
		beforeId := string(Curl(AppUri(appName, "/id", LoadConfig().AppsDomain)).FullOutput())

		Expect(Curl(AppUri(appName, "/logspew/33554432", LoadConfig().AppsDomain))).To(
			Say("Just wrote 33554432 random bytes to the log"),
		)

		// Give time for components (i.e. Warden) to react to the output
		// and potentially make bad decisions (like killing the app)
		time.Sleep(10 * time.Second)

		afterId := string(Curl(AppUri(appName, "/id", LoadConfig().AppsDomain)).FullOutput())

		Expect(beforeId).To(Equal(afterId))

		Expect(Curl(AppUri(appName, "/logspew/2", LoadConfig().AppsDomain))).To(
			Say("Just wrote 2 random bytes to the log"),
		)
	})
})
