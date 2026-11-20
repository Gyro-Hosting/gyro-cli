## gyro backup-space credentials

SFTP connection with the password

### Synopsis

SFTP connection with the password

Host, port, user, path and the SFTP password (`Cache-Control: no-store`). Needs `services_write` because it reveals a secret. 409 `not_ready` while the space is being set up.

API: GET /cloud/backup-spaces/{id}/credentials (operationId getBackupSpaceCredentials, scope services_write)

```
gyro backup-space credentials <ID> [flags]
```

### Options

```
  -h, --help   help for credentials
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

