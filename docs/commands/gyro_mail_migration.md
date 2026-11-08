## gyro mail migration

IMAP migration jobs

```
gyro mail migration [flags]
```

### Options

```
  -h, --help   help for migration
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

* [gyro mail](gyro_mail.md)	 - Corporate mail: mailboxes, aliases, forwarding ...
* [gyro mail migration create](gyro_mail_migration_create.md)	 - create (fields with --set key=value)
* [gyro mail migration delete](gyro_mail_migration_delete.md)	 - delete (fields with --set key=value)
* [gyro mail migration list](gyro_mail_migration_list.md)	 - Show
* [gyro mail migration update](gyro_mail_migration_update.md)	 - update (fields with --set key=value)

