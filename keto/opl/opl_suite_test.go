package opl_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opl Suite")
}

var _ = BeforeSuite(func() {
	err := godotenv.Load("../.env") // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}

	Expect(os.Getenv("KETO_READ_REMOTE")).NotTo(BeZero(), "Please make sure KETO_READ_REMOTE is set correctly.")
	Expect(os.Getenv("KETO_WRITE_REMOTE")).NotTo(BeZero(), "Please make sure KETO_WRITE_REMOTE is set correctly.")
	Expect(os.Getenv("KETO_BEARER_TOKEN")).NotTo(BeZero(), "Please make sure KETO_WRITE_REMOTE is set correctly.")

	GinkgoWriter.Printf("KETO_READ_REMOTE: %s\n", os.Getenv("KETO_READ_REMOTE"))
	GinkgoWriter.Printf("KETO_WRITE_REMOTE: %s\n", os.Getenv("KETO_WRITE_REMOTE"))

})
