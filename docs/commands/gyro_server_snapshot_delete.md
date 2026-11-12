## gyro server snapshot delete

Delete a snapshot

### Synopsis

Delete a snapshot

**Billing (tr-ist-1, prepaid):** the first snapshot is free; every additional snapshot is its own hourly item and is billed **at least 1 hour**, even if it is deleted within minutes (the oldest snapshot counts as the free one).

API: DELETE /cloud/servers/{id}/snapshots/{sid} (operationId deleteSnapshot, scope cloud_write)

```
gyro server snapshot delete <ID> <SID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
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

