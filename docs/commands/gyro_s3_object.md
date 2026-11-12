## gyro s3 object

List objects in a bucket folder

```
gyro s3 object [flags]
```

### Options

```
  -h, --help   help for object
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
* [gyro s3 object delete](gyro_s3_object_delete.md)	 - Delete objects
* [gyro s3 object download-url](gyro_s3_object_download-url.md)	 - Presigned download URL
* [gyro s3 object legal-hold](gyro_s3_object_legal-hold.md)	 - Legal hold of an object version
* [gyro s3 object list](gyro_s3_object_list.md)	 - List objects in a bucket folder
* [gyro s3 object retention](gyro_s3_object_retention.md)	 - Retention of an object version
* [gyro s3 object upload-url](gyro_s3_object_upload-url.md)	 - Presigned upload URLs

