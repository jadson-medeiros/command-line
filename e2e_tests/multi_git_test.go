package e2e_tests

import (
	"os"

	. "github.com/jadson-medeiros/command-line/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const baseDir = "/tmp/multi-git"

var _ = Describe("multi-git e2e tests", func() {
	var err error

	removeAll := func() {
		err = os.RemoveAll(baseDir)
		Ω(err).Should(BeNil())
	}

	BeforeEach(func() {
		removeAll()
		err = CreateDir(baseDir, "", false)
		Ω(err).Should(BeNil())
	})

	AfterSuite(removeAll)
})
