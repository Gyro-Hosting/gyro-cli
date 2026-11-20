## gyro backup-space plans

List backup space plans

### Synopsis

List backup space plans

Orderable sizes (50 GB to 1 TB) with prices per billing cycle (fixed EUR, VAT excluded; `list_amount` = strike-through list price) and `servers`: the servers of the account a space can be bound to. Empty `plans` = not orderable for this account yet. `available: false` = no capacity right now.

API: GET /cloud/backup-spaces/plans (operationId listBackupSpacePlans, scope cloud_read)

```
gyro backup-space plans [flags]
```

### Options

```
  -h, --help   help for plans
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

