## gyro backup-space

Backup Space: SFTP backup space bound to one of your servers (tr-ist-1)

```
gyro backup-space [flags]
```

### Options

```
  -h, --help   help for backup-space
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro backup-space attach-server](gyro_backup-space_attach-server.md)	 - Bind the space to another server
* [gyro backup-space create](gyro_backup-space_create.md)	 - Order a backup space
* [gyro backup-space credentials](gyro_backup-space_credentials.md)	 - SFTP connection with the password
* [gyro backup-space delete](gyro_backup-space_delete.md)	 - Cancel (delete) a backup space
* [gyro backup-space get](gyro_backup-space_get.md)	 - Get a backup space
* [gyro backup-space list](gyro_backup-space_list.md)	 - List backup spaces
* [gyro backup-space plans](gyro_backup-space_plans.md)	 - List backup space plans
* [gyro backup-space reset-password](gyro_backup-space_reset-password.md)	 - Reset the SFTP password

