## gyro s3 key

Access keys (secret shown once)

```
gyro s3 key [flags]
```

### Options

```
  -h, --help   help for key
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
* [gyro s3 key create](gyro_s3_key_create.md)	 - create (fields with --set key=value)
* [gyro s3 key delete](gyro_s3_key_delete.md)	 - delete (fields with --set key=value)
* [gyro s3 key list](gyro_s3_key_list.md)	 - Show
* [gyro s3 key rotate](gyro_s3_key_rotate.md)	 - POST access-keys:rotate (GyroDepolama)
* [gyro s3 key update](gyro_s3_key_update.md)	 - update (fields with --set key=value)

