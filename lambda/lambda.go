package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var invokeCount = 0

func LambdaHandler() (int, error) {
	invokeCount = invokeCount + 1
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	// Create S3 service client
	/*
		svc := s3.New(sess)

			input := &s3.ListBucketsInput{}
			result, err := svc.ListBuckets(input)
			if err != nil {
				log.Fatal(err)
				return 0, err
			}
			log.Print(result)
	*/

	/*
		downloader := s3manager.NewDownloader(sess)

		file, err := os.Create("1750.txt")
		if err != nil {
			exitErrorf("Unable to open file %q, %v", "1750.txt", err)
		}

		defer file.Close()
		numBytes, err := downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String("source-gutenberg"),
				Key:    aws.String("1/7/5/1750/1750.txt"),
			})

		if err != nil {
			exitErrorf("Unable to download item %q, %v", "1750.txt", err)
		}

		log.Println("Downloaded", file.Name(), numBytes, "bytes")
	*/

	svc := s3.New(sess)

	params := &s3.GetObjectInput{
		Bucket: aws.String("source-gutenberg"),
		Key:    aws.String("1/7/5/1750/1750.txt"),
	}

	result, err := svc.GetObject(params)
	if err != nil {
		exitErrorf("bork: %v", err)
	}

	buf := make([]byte, 100)
	for i := 0; i < 20; i++ {
		_, err := result.Body.Read(buf)
		if err != nil {
			exitErrorf("error reading: %v", err)
		}
		log.Printf("Chunk %v: %v", i+1, string(buf))
	}

	return invokeCount, nil
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	lambda.Start(LambdaHandler)
}
