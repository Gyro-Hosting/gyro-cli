## gyro s3 versions

List object versions

```
gyro s3 versions [flags]
```

### Options

```
  -h, --help   help for versions
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
* [gyro s3 versions delete](gyro_s3_versions_delete.md)	 - Permanently delete a version
* [gyro s3 versions list](gyro_s3_versions_list.md)	 - List object versions
* [gyro s3 versions restore](gyro_s3_versions_restore.md)	 - Restore an older version
* [gyro s3 versions undelete](gyro_s3_versions_undelete.md)	 - Remove a delete marker

