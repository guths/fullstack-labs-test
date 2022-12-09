package challenge1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChallenge1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Challenge1 Suite")
}
