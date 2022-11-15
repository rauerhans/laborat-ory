package opl_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOplTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OplTest Suite")
}
