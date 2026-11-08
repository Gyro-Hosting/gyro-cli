## gyro reseller report

Usage and billing report per end-customer

```
gyro reseller report [flags]
```

### Options

```
      --from string   
  -h, --help          help for report
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

* [gyro reseller](gyro_reseller.md)	 - Reseller: end-customers, assignments, keys, report
* [gyro reseller report get](gyro_reseller_report_get.md)	 - Usage and billing report per end-customer

