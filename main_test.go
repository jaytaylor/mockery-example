package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jaytaylor/mockery-example/mocks"
	"github.com/stretchr/testify/mock"
)

func TestS3Mock(t *testing.T) {
	mockS3 := &mocks.S3API{}

	mockS3.On("ListObjects", mock.MatchedBy(func(input *s3.ListObjectsInput) bool {
		return input.Delimiter != nil && *input.Delimiter == "/" && input.Prefix == nil
	})).Return(func(input *s3.ListObjectsInput) *s3.ListObjectsOutput {
		output := &s3.ListObjectsOutput{}
		output.SetCommonPrefixes([]*s3.CommonPrefix{
			&s3.CommonPrefix{
				Prefix: aws.String("2017-01-01"),
			},
			&s3.CommonPrefix{
				Prefix: aws.String("2017-01-02"),
			},
			&s3.CommonPrefix{
				Prefix: aws.String("2017-01-03"),
			},
			&s3.CommonPrefix{
				Prefix: aws.String("2017-01-04"),
			},
			&s3.CommonPrefix{
				Prefix: aws.String("2017-01-05"),
			},
			&s3.CommonPrefix{
				Prefix: aws.String("2017-02-13"),
			},
		})
		return output
	})

	listingInput := &s3.ListObjectsInput{
		Bucket:    aws.String("foo"),
		Delimiter: aws.String("/"),
	}
	listingOutput, err := mockS3.ListObjects(listingInput)
	if err != nil {
		t.Fatalf("Error listing keys: %s", err)
	}

	for _, x := range listingOutput.Contents {
		t.Logf("%+v", *x)
	}
}
