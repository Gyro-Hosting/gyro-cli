## gyro server snapshot-schedule delete

Remove the snapshot schedule

### Synopsis

Remove the snapshot schedule

The schedule is removed; existing snapshots (automatic ones too) stay. Gyro Cloud Istanbul (tr-ist-1) servers only (other servers: 422 `not_supported`). A snapshot named `auto-YYYYMMDD-HHMM` is taken every day or every week at `hour`:`minute` Türkiye time (Europe/Istanbul); before it the oldest automatic snapshots are deleted so at most `keep` of them stay. Snapshots you take yourself are never deleted, but they count toward the server's snapshot limit (`snapshot_limit`, manual + automatic): when the limit is full the run is skipped (`last_status` skipped, `last_error` snapshot_limit). Snapshots above the free allowance of a prepaid server are billed hourly like manual ones. While the server has snapshots an extra disk cannot be grown (409 `has_snapshots`): turn the schedule off and delete the snapshots first.

API: DELETE /cloud/servers/{id}/snapshot-schedule (operationId deleteSnapshotSchedule, scope cloud_write)

```
gyro server snapshot-schedule delete <ID> [flags]
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

* [gyro server snapshot-schedule](gyro_server_snapshot-schedule.md)	 - Scheduled snapshots of a server

