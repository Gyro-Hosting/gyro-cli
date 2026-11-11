## gyro billing statement

Monthly usage statement (not an invoice); --pdf / --csv to download

```
gyro billing statement YYYY-MM [flags]
```

### Examples

```
  gyro billing statement 2026-11
  gyro billing statement 2026-11 --pdf .
  gyro billing statement 2026-11 --csv kasim.csv
```

### Options

```
      --csv string   save as CSV (file, '.' = server name)
  -h, --help         help for statement
      --pdf string   save as PDF (file, '.' = server name)
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

