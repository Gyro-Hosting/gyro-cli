## gyro server snapshot list

List snapshots and the snapshot quota

### Synopsis

List snapshots and the snapshot quota

**Billing (tr-ist-1, prepaid):** the first snapshot is free; every additional snapshot is its own hourly item and is billed **at least 1 hour**, even if it is deleted within minutes (the oldest snapshot counts as the free one). tr-ist-1: `quota` = number of snapshots allowed on the plan (add-on), snapshot ids are stable numbers derived by the API. Prepaid tr-ist-1 servers also return `pricing` {billing, free, extra_monthly_eur, extra_hourly_eur, billable}: the first snapshot(s) are free, each extra snapshot is charged hourly from the balance while it exists. 2.14.0 (Gyro Compute): snapshot ids are the platform's own ids; `name` is returned too.

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

