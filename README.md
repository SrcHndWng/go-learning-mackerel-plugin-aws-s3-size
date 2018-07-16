# go-learning-mackerel-plugin-aws-s3-size

This is a sample mackerel-plugin for learning.
Get S3 object size for mackerel.io agent.

## Example of mackerel-agent.conf

```
[plugin.metrics.aws-s3-size]
command = "/home/ec2-user/golang/go-learning-mackerel-plugin-aws-s3-size --bucket your-bucket --key key-path-under-bucket --key key-path-under-bucket --region your-region"
```

Or

```
[plugin.metrics.aws-s3-size]
command = "/home/ec2-user/golang/go-learning-mackerel-plugin-aws-s3-size --bucket your-bucket1 --key key-path-under-bucket1 -bucket your-bucket2 --key key-path-under-bucket2 --region your-region"
```
