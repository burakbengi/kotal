package ethereum2

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEthereum2Client(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ethereum 2.0 Clients Suite")
}
