package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	log.Println("Hello world")

	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "playground")}))

	//(&aws.Config{
	//	Region:      aws.String("us-west-2"),
	//	Credentials: credentials.NewSharedCredentials("", "test-account"),
	//})

	ss3Ep := s3.New(sess)
	buckets, err := ss3Ep.ListBuckets(nil)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	log.Printf("Buckerts: %s", buckets)

}
