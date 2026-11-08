## gyro reseller report get

Usage and billing report per end-customer

### Synopsis

Usage and billing report per end-customer

API: GET /cloud/reseller/report (operationId resellerReport, scope reseller)

```
gyro reseller report get [flags]
```

### Options

```
      --from string   
  -h, --help          help for get
      --to string     
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

* [gyro reseller report](gyro_reseller_report.md)	 - Usage and billing report per end-customer

