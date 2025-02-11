package authenticate_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuthenticate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authenticate Suite")
}
