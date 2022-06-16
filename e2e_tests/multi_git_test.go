package e2e_tests

import (
	"fmt"
	"os"

	. "github.com/jadson-medeiros/command-line/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const baseDir = "/tmp/multi-git"

var repoList string

var _ = Describe("multi-git e2e tests", func() {
	var err error

	fmt.Println("*** e2e_tests starting")
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

	Context("Tests for empty/undefined environment failure cases", func() {
		It("Should fail with invalid base dir", func() {
			output, err := RunMultiGit("status", false, "/no-such-dir", repoList)
			Ω(err).ShouldNot(BeNil())
			suffix := "base dir: '/no-such-dir/' doesn't exist\n"
			Ω(output).Should(HaveSuffix(suffix))
		})

		It("Should fail with empty repo list", func() {
			output, err := RunMultiGit("status", false, baseDir, repoList)
			Ω(err).ShouldNot(BeNil())
			Ω(output).Should(ContainSubstring("repo list can't be empty"))
		})
	})
})
