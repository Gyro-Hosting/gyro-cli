## gyro billing usage

Hour-by-hour usage (or daily/monthly totals)

```
gyro billing usage [flags]
```

### Examples

```
  gyro billing usage --from 2026-11-01 --to 2026-11-30
  gyro billing usage --service 17501 --daily
  gyro billing usage --monthly -o json
  gyro billing usage --csv kullanim.csv
```

### Options

```
      --all             fetch every page
      --csv string      save the hours as CSV (file, '.' = server name, '-' = stdout)
      --daily           daily totals
      --from string     first day YYYY-MM-DD (default: 1st of this month)
  -h, --help            help for usage
      --monthly         monthly totals
      --service int     only this service id
      --status string   ucretli|tavan|kismi|odenmedi|ucretsiz
      --to string       last day YYYY-MM-DD (default: today)
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

* [gyro billing](gyro_billing.md)	 - Prepaid hourly usage (Istanbul cloud, containers): balance, usage, estimate, statement

