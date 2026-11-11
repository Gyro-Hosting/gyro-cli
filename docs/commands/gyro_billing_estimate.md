## gyro billing estimate

Hourly, daily and monthly (capped) cost of plans; how long your balance lasts

```
gyro billing estimate [flags]
```

### Examples

```
  gyro billing estimate --plan standart-2-4 --region tr-ist-1 --count 3
  gyro billing estimate --item standart-2-4:2 --item standart-4-8:1
```

### Options

```
      --count int          how many (default 1)
  -h, --help               help for estimate
      --item stringArray   plan:count, repeatable
      --plan string        plan name, e.g. standart-2-4
      --product-id int     or the product id
      --region string      region (prepaid: tr-ist-1) (default "tr-ist-1")
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

