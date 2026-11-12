## gyro billing

Prepaid hourly usage (Istanbul cloud, containers): balance, usage, estimate, statement

```
gyro billing [flags]
```

### Options

```
  -h, --help   help for billing
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro billing balance](gyro_billing_balance.md)	 - Balance, hourly cost and estimated time left
* [gyro billing estimate](gyro_billing_estimate.md)	 - Hourly, daily and monthly cost of plans and serverless containers; how long your balance lasts
* [gyro billing statement](gyro_billing_statement.md)	 - Monthly usage statement (not an invoice); --pdf / --csv to download
* [gyro billing topup-url](gyro_billing_topup-url.md)	 - Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)
* [gyro billing usage](gyro_billing_usage.md)	 - Hour-by-hour usage (or daily/monthly totals)

