## gyro billing topup-url

Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)

```
gyro billing topup-url [flags]
```

### Options

```
  -h, --help   help for topup-url
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
* [gyro billing topup-url get](gyro_billing_topup-url_get.md)	 - Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)

