## gyro server snapshot list

List snapshots and the snapshot quota

### Synopsis

List snapshots and the snapshot quota

tr-ist-1: `quota` = number of snapshots allowed on the plan (add-on), snapshot ids are stable numbers derived by the API.

API: GET /cloud/servers/{id}/snapshots (operationId listSnapshots, scope cloud_read)

```
gyro server snapshot list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro server snapshot](gyro_server_snapshot.md)	 - List snapshots and the snapshot quota

