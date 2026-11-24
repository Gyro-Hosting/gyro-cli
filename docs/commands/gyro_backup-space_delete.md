## gyro backup-space delete

Cancel (delete) a backup space

### Synopsis

Cancel (delete) a backup space

`confirm=true` is required. `immediate=true` cancels now: SFTP access stops and the space with EVERY file in it is deleted permanently within about a minute (no snapshot, no grace period; cannot be undone). Without `immediate` a cancellation request is filed and the space keeps working until it is processed. 403 `immediate_disabled` when the operator turned immediate deletion off.

API: DELETE /cloud/backup-spaces/{id} (operationId deleteBackupSpace, scope cloud_delete)

```
gyro backup-space delete <ID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
      --immediate                delete the space and every file now (cannot be undone)
```

### Options inherited from parent commands

```
      --account gyro account list   act in another account you are a panel user of (id from gyro account list; "self" = your own; env GYRO_ACCOUNT, profile default: `gyro account use`)
      --api-url string              API base URL (env GYRO_API_URL)
      --debug                       log HTTP requests to stderr (secrets masked)
      --lang string                 message language: tr or en (default from locale)
      --no-color                    disable colours (also NO_COLOR)
      --no-header                   omit the table header
  -o, --output string               output format: table, json, yaml, text
  -p, --profile string              profile to use (env GYRO_PROFILE)
      --query string                JMESPath query applied to the result
  -q, --quiet                       print only ids
      --timeout duration            overall timeout, e.g. 30s, 10m
      --wide                        show all table columns
  -y, --yes                         do not ask for confirmation
```

### SEE ALSO

* [gyro backup-space](gyro_backup-space.md)	 - Backup Space: SFTP backup space bound to one of your servers (tr-ist-1)

