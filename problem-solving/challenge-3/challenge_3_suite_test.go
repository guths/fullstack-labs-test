package challenge3

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChallenge3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Challenge3 Suite")
}
