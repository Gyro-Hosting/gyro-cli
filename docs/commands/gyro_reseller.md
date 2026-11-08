## gyro reseller

Reseller: end-customers, assignments, keys, report

```
gyro reseller [flags]
```

### Options

```
  -h, --help   help for reseller
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
* [gyro reseller assign](gyro_reseller_assign.md)	 - Assign a service to an end-customer (customer_id 0 = unassign)
* [gyro reseller customer](gyro_reseller_customer.md)	 - List end-customers
* [gyro reseller report](gyro_reseller_report.md)	 - Usage and billing report per end-customer
* [gyro reseller setting](gyro_reseller_setting.md)	 - Reseller settings

