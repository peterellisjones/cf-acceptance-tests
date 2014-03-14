package apps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/vito/cmdtest/matchers"

	. "github.com/peterellisjones/cf-acceptance-tests/helpers"
	. "github.com/peterellisjones/cf-test-helpers/cf"
	. "github.com/peterellisjones/cf-test-helpers/generator"
	"time"
)

var _ = PDescribe("loggregator", func() {
	var appName string

	BeforeEach(func() {
		appName = RandomName()

		Expect(Cf("push", appName, "-p", NewAssets().Dora)).To(SayWithTimeout("App started", time.Minute*2))
	})

	AfterEach(func() {
		Expect(Cf("delete", appName, "-f")).To(SayWithTimeout("OK", time.Minute*2))
	})

	Context("gcf logs", func() {
		PIt("blocks and exercises basic loggregator behavior", func() {
			logs := Cf("logs", appName)

			Expect(logs).To(SayWithTimeout("Connected, tailing logs for app", time.Second*15))

			Eventually(Curling(appName, "/", LoadConfig().AppsDomain)).Should(Say("Hi, I'm Dora!"))

			Expect(logs).To(SayWithTimeout("OUT "+appName+"."+LoadConfig().AppsDomain, time.Second*15))
		})
	})

	Context("gcf logs --recent", func() {
		It("makes loggregator buffer and dump log messages", func() {
			logs := Cf("logs", appName, "--recent")

			Expect(logs).To(SayWithTimeout("Connected, dumping recent logs for app", time.Second*15))

			Expect(logs).To(SayWithTimeout("OUT Created app", time.Second*15))
			Expect(logs).To(SayWithTimeout("OUT Starting app instance", time.Second*15))
		})
	})
})
