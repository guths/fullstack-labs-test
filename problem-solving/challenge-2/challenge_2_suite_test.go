package challenge2

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChallenge2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Challenge2 Suite")
}
