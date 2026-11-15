## gyro server snapshot-schedule set

Create or change the snapshot schedule

### Synopsis

Create or change the snapshot schedule

Fields left out keep their value (a new schedule starts from daily 03:00, keep 7, enabled). `keep` is 1 to the server's snapshot limit; `weekday` only for weekly. Gyro Cloud Istanbul (tr-ist-1) servers only (other servers: 422 `not_supported`). A snapshot named `auto-YYYYMMDD-HHMM` is taken every day or every week at `hour`:`minute` Türkiye time (Europe/Istanbul); before it the oldest automatic snapshots are deleted so at most `keep` of them stay. Snapshots you take yourself are never deleted, but they count toward the server's snapshot limit (`snapshot_limit`, manual + automatic): when the limit is full the run is skipped (`last_status` skipped, `last_error` snapshot_limit). Snapshots above the free allowance of a prepaid server are billed hourly like manual ones. While the server has snapshots an extra disk cannot be grown (409 `has_snapshots`): turn the schedule off and delete the snapshots first.

API: PUT /cloud/servers/{id}/snapshot-schedule (operationId setSnapshotSchedule, scope cloud_write)

```
gyro server snapshot-schedule set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --enabled                  enabled
      --frequency string         frequency (daily|weekly)
  -h, --help                     help for set
      --hour int                 Türkiye time
      --idempotency-key string   idempotency key (automatic for money operations)
      --keep int                 <= snapshot_limit
      --minute int               minute
      --set stringArray          extra body field key=value (repeatable)
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
      --weekday string           weekly only (0-6 also accepted, 0 = Sunday) (sunday|monday|tuesday|wednesday|thursday|friday|saturday)
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

