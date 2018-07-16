package mpawss3size

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Key struct {
	Bucket string
	Key    string
	Title  string
}

func getS3Keys() []s3Key {
	var s3Keys []s3Key
	for i, k := range keys {
		var bucket string
		if len(buckets) == 1 {
			bucket = buckets[0]
		} else {
			bucket = buckets[i]
		}
		t := fmt.Sprintf("%s-%s", bucket, k)
		t = strings.Replace(t, "/", "-", -1)
		t = strings.Replace(t, ".", "-", -1)
		s3Keys = append(s3Keys, s3Key{Bucket: bucket, Key: k, Title: t})
	}
	return s3Keys
}

func getS3KeySize(bucket string, key string, region string) (float64, error) {
	svc := s3.New(session.New(), &aws.Config{
		Region: aws.String(region),
	})
	params := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	result, err := svc.HeadObject(params)
	if err != nil {
		return 0, err
	}
	return float64(*result.ContentLength), nil
}
