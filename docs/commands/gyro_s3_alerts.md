## gyro s3 alerts

Fill alerts of all buckets

```
gyro s3 alerts [flags]
```

### Options

```
  -h, --help   help for alerts
```

### Options inherited from parent commands

```
      --api-url string     API base URL (env GYRO_API_URL)
      --debug              log HTTP requests to stderr (secrets masked)
      --lang string        message language: tr or en (default from locale)
      --no-color           disable colours (also NO_COLOR)
      --no-header          omit the table header
  -o, --output string      output format: table, json, yaml, text
  -p, --profile string     profile to use (env GYRO_PROFILE)
      --query string       JMESPath query applied to the result
  -q, --quiet              print only ids
      --timeout duration   overall timeout, e.g. 30s, 10m
      --wide               show all table columns
  -y, --yes                do not ask for confirmation
```

### SEE ALSO

* [gyro s3](gyro_s3.md)	 - S3 object storage: buckets and access keys
* [gyro s3 alerts delete](gyro_s3_alerts_delete.md)	 - Remove the fill alert of a bucket
* [gyro s3 alerts get](gyro_s3_alerts_get.md)	 - Fill alert of a bucket
* [gyro s3 alerts list](gyro_s3_alerts_list.md)	 - Fill alerts of all buckets
* [gyro s3 alerts set](gyro_s3_alerts_set.md)	 - Set the fill alert of a bucket

