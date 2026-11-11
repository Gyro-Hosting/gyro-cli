## gyro billing balance

Balance, hourly cost and estimated time left

```
gyro billing balance [flags]
```

### Examples

```
  gyro billing balance
  gyro billing balance -o json --query hours_left
```

### Options

```
  -h, --help   help for balance
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

