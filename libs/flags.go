package mpawss3size

import (
	"errors"
	"fmt"
)

type bucketFlags []string
type keyFlags []string

func (s *bucketFlags) String() string {
	return fmt.Sprintf("%v", buckets)
}

func (s *bucketFlags) Set(v string) error {
	*s = append(*s, v)
	return nil
}

func (s *keyFlags) String() string {
	return fmt.Sprintf("%v", keys)
}

func (s *keyFlags) Set(v string) error {
	*s = append(*s, v)
	return nil
}

var buckets bucketFlags
var keys keyFlags
var region string

func validateParams() error {
	if len(buckets) == 0 {
		return errors.New("bucket is not set")
	}
	if len(keys) == 0 {
		return errors.New("keys is not set")
	}
	if len(buckets) > 1 && (len(buckets) != len(keys)) {
		return errors.New("buckets count is not equal to keys count")
	}
	return nil
}
