package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	pieceio "github.com/filecoin-project/go-fil-markets/pieceio"
)

type CommPRequest struct {
	Region string `json:"region"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

// { "region": "us-east-2", "source-gutenberg", "1/7/5/1750/1750.txt" }

type CommPResponse struct {
	Region     string `json:"region"`
	Bucket     string `json:"bucket"`
	Key        string `json:"key"`
	CommP      string `json:"commp"`
	Size       uint64 `json:"size"`
	PaddedSize uint64 `json:"paddedSize"`
}

func LambdaHandler(ctx context.Context, request CommPRequest) (CommPResponse, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(request.Region)},
	)
	svc := s3.New(sess)

	params := &s3.GetObjectInput{
		Bucket: aws.String(request.Bucket),
		Key:    aws.String(request.Key),
	}

	result, err := svc.GetObject(params)
	if err != nil {
		exitErrorf("Could not fetch S3 object: %v", err)
	}

	size := uint64(*result.ContentLength)
	commp, psize, err := pieceio.GeneratePieceCommitment(result.Body, size)
	if err != nil {
		exitErrorf("Error generating CommP: %v", err)
	}

	return CommPResponse{
		Region:     request.Region,
		Bucket:     request.Bucket,
		Key:        request.Key,
		CommP:      hex.EncodeToString(commp),
		Size:       size,
		PaddedSize: psize,
	}, nil
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	lambda.Start(LambdaHandler)
}
