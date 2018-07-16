package mpawss3size

import (
	"fmt"
	"testing"
)

func TestValidateParams(t *testing.T) {
	invalidTest := func(err error) {
		if err == nil {
			t.Error("validate failed.")
		}
		fmt.Println(err)
	}

	validTest := func(err error) {
		if err != nil {
			t.Error(err)
		}
	}

	keys.Set("test-object1")
	keys.Set("test-object2")
	err := validateParams()
	invalidTest(err)

	buckets.Set("test-bucket")
	keys = nil
	err = validateParams()
	invalidTest(err)

	buckets = nil
	keys = nil
	buckets.Set("test-bucket1")
	buckets.Set("test-bucket2")
	keys.Set("test-object1")
	keys.Set("test-object2")
	keys.Set("test-object3")
	err = validateParams()
	invalidTest(err)

	buckets = nil
	keys = nil
	buckets.Set("test-bucket1")
	keys.Set("test-object1")
	keys.Set("test-object2")
	err = validateParams()
	validTest(err)

	buckets = nil
	keys = nil
	buckets.Set("test-bucket1")
	buckets.Set("test-bucket2")
	buckets.Set("test-bucket3")
	keys.Set("test-object1")
	keys.Set("test-object2")
	keys.Set("test-object3")
	err = validateParams()
	validTest(err)
}
