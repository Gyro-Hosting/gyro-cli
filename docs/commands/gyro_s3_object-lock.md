## gyro s3 object-lock

Default retention of a bucket

```
gyro s3 object-lock [flags]
```

### Options

```
  -h, --help   help for object-lock
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
* [gyro s3 object-lock get](gyro_s3_object-lock_get.md)	 - Default retention of a bucket
* [gyro s3 object-lock scan](gyro_s3_object-lock_scan.md)	 - Scan locked versions now
* [gyro s3 object-lock set](gyro_s3_object-lock_set.md)	 - Set the default retention of a bucket
* [gyro s3 object-lock status](gyro_s3_object-lock_status.md)	 - Protected Storage: object lock summary

