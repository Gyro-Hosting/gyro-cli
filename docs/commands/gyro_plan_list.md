## gyro plan list

List orderable plans with prices

```
gyro plan list [flags]
```

### Examples

```
  gyro plan list --location falkenstein
  gyro plan list --cycle annually -o json
```

### Options

```
      --all               include out-of-stock plans
      --cycle string      monthly|quarterly|semiannually|annually (default "monthly")
  -h, --help              help for list
      --location string   region (eu-de-1) or city
      --plan string       plan name, e.g. standart-2-4
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

* [gyro plan](gyro_plan.md)	 - Server plans and prices (catalog)

