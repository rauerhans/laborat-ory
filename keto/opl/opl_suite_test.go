package opl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opl Suite")
}
