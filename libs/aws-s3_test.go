package mpawss3size

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetS3Keys(t *testing.T) {
	test := func(result []s3Key, expect []s3Key) {
		if !reflect.DeepEqual(result, expect) {
			fmt.Printf("result = %v\n", result)
			t.Error("Invalid result.")
		}
	}

	var expect []s3Key

	buckets.Set("test-bucket")
	keys.Set("test-object1.txt")
	keys.Set("test-folder/test-object2")

	result := getS3Keys()
	expect = append(expect, s3Key{Bucket: "test-bucket", Key: "test-object1.txt", Title: "test-bucket-test-object1-txt"})
	expect = append(expect, s3Key{Bucket: "test-bucket", Key: "test-folder/test-object2", Title: "test-bucket-test-folder-test-object2"})
	test(result, expect)

	buckets.Set("test-bucket2")
	result = getS3Keys()
	expect = nil
	expect = append(expect, s3Key{Bucket: "test-bucket", Key: "test-object1.txt", Title: "test-bucket-test-object1-txt"})
	expect = append(expect, s3Key{Bucket: "test-bucket2", Key: "test-folder/test-object2", Title: "test-bucket2-test-folder-test-object2"})
	test(result, expect)
}

func TestGetS3KeySize(t *testing.T) {
	// Get object size that is a sample data for Redshift.
	// https://docs.aws.amazon.com/ja_jp/redshift/latest/gsg/rs-gsg-create-sample-db.html
	bucket := "awssampledbuswest2"
	key := "tickit/allusers_pipe.txt"
	region := "us-west-2"
	size, err := getS3KeySize(bucket, key, region)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("size = %v\n", size)
}
