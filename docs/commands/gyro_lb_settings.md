## gyro lb settings

Algorithm, reverse DNS, private network

### Synopsis

Algorithm, reverse DNS, private network

Current algorithm, PTR records, your private networks in the same region (attached or not) and limits.

API: GET /cloud/load-balancers/{id}/settings (operationId typed_get_load_balancers__id__settings, scope services_read)

```
gyro lb settings <ID> [flags]
```

### Options

```
  -h, --help   help for settings
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

* [gyro lb](gyro_lb.md)	 - List load balancers

