package mpawss3size

import (
	"flag"
	"fmt"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

// AwsS3SizePlugin mackerel plugin
type AwsS3SizePlugin struct {
	Prefix string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (a AwsS3SizePlugin) MetricKeyPrefix() string {
	if a.Prefix == "" {
		a.Prefix = "aws-s3-size"
	}
	return a.Prefix
}

// GraphDefinition interface for mackerelplugin
func (a AwsS3SizePlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(a.Prefix)
	ms := make([]mp.Metrics, 0)
	s3Keys := getS3Keys()
	for _, k := range s3Keys {
		ms = append(ms, mp.Metrics{Name: k.Title, Label: k.Title})
	}

	return map[string]mp.Graphs{
		"S3Size": {
			Label:   labelPrefix,
			Unit:    mp.UnitInteger,
			Metrics: ms,
		},
	}
}

// FetchMetrics interface for mackerelplugin
func (a AwsS3SizePlugin) FetchMetrics() (map[string]float64, error) {
	err := validateParams()
	if err != nil {
		return nil, err
	}
	size := make(map[string]float64)
	s3Keys := getS3Keys()
	for _, k := range s3Keys {
		f, err := getS3KeySize(k.Bucket, k.Key, region)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		size[k.Title] = f
	}
	return size, nil
}

// Do the plugin
func Do() {
	optPrefix := flag.String("metric-key-prefix", "aws-s3-size", "Metric key prefix")

	flag.Var(&buckets, "bucket", "S3 Bucket")
	flag.Var(&keys, "key", "S3 Key")
	flag.StringVar(&region, "region", "", "S3 Region")

	flag.Parse()

	a := AwsS3SizePlugin{
		Prefix: *optPrefix,
	}
	helper := mp.NewMackerelPlugin(a)
	helper.Run()
}
