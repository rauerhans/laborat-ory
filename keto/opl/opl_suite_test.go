package opl_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	ketoclient "github.com/rauerhans/laborat-ory/keto/client"
)

func TestOpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opl Suite")
}

var wcl rts.WriteServiceClient
var rcl rts.ReadServiceClient
var ccl rts.CheckServiceClient

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

	conn, err := ketoclient.GetWriteConn(context.TODO())
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	wcl = rts.NewWriteServiceClient(conn)

	conn, err = ketoclient.GetReadConn(context.TODO())
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	rcl = rts.NewReadServiceClient(conn)

	conn, err = ketoclient.GetReadConn(context.TODO())
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	ccl = rts.NewCheckServiceClient(conn)

})
