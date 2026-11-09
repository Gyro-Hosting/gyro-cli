## gyro s3 folder count

Count objects in a folder

### Synopsis

Count objects in a folder

Number of objects under `prefix` (`more` = over 100000 or time limit).

API: GET /cloud/storage/{id}/buckets/{bucket}/folders/count (operationId typed_get_storage__id__buckets__bucket__folders_count, scope services_read)

```
gyro s3 folder count <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help            help for count
      --prefix string   Folder, ends with /
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

* [gyro s3 folder](gyro_s3_folder.md)	 - Count objects in a folder

