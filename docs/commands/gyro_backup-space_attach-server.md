## gyro backup-space attach-server

Bind the space to another server

### Synopsis

Bind the space to another server

Binds the space to `server_id` (one of your active servers; optional when the account has one). From then on only that server's IPv4 addresses can connect; the previous server loses access at once. Files are not touched. When the bound server's IP changes or the server is cancelled, access follows automatically.

API: PUT /cloud/backup-spaces/{id}/server (operationId attachBackupSpaceServer, scope services_write)

```
gyro backup-space attach-server <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for attach-server
      --idempotency-key string   idempotency key (automatic for money operations)
      --server-id int            server id
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro backup-space](gyro_backup-space.md)	 - Backup Space: SFTP backup space bound to one of your servers (tr-ist-1)

