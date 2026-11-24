## gyro backup-space list

List backup spaces

### Synopsis

List backup spaces

Backup spaces of the account (same filters and paging as GET /cloud/services; `product_line=backup_space`).

API: GET /cloud/backup-spaces (operationId listBackupSpaces, scope services_read)

```
gyro backup-space list [flags]
```

### Options

```
      --all             fetch every page
  -h, --help            help for list
      --limit int       items per page (max 100)
      --page int        page number (default 1)
      --status string   service status filter (active, suspended, cancelled …)
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

