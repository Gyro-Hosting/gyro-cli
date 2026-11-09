## gyro s3 object list

List objects in a bucket folder

### Synopsis

List objects in a bucket folder

One page (200) of a folder: `folders` (common prefixes) and `files` (key, size, modified). Pass `next` from the previous page to continue.

API: GET /cloud/storage/{id}/buckets/{bucket}/objects (operationId typed_get_storage__id__buckets__bucket__objects, scope services_read)

```
gyro s3 object list <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help            help for list
      --next string     Continuation token
      --prefix string   Folder, ends with / (empty = root)
      --search string   Name prefix inside the folder
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

* [gyro s3 object](gyro_s3_object.md)	 - List objects in a bucket folder

