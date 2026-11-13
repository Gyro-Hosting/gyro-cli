## gyro server disk list

List the server's additional disks

### Synopsis

List the server's additional disks

Two classes: `ssd` (fast) and `hdd` (large storage for backups, archives and media; do not run databases on it). Additional disks are billed per GB hourly from the prepaid balance (`classes.<class>.per_gb_monthly_eur`, monthly / 672), from the hour they are added until they are removed; each disk at least `pricing.minimum_hours`. Up to 4 disks per server and `classes.<class>.max_gb` GB per class; sizes are multiples of `step_gb`. A new disk is empty (partition, format and mount it inside the server). Reinstalling the server removes the additional disks.

API: GET /cloud/servers/{id}/disks (operationId listServerDisks, scope cloud_read)

```
gyro server disk list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro server disk](gyro_server_disk.md)	 - List the server's additional disks

