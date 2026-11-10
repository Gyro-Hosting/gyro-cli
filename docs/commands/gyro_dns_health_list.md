## gyro dns health list

List health-checked records

### Synopsis

List health-checked records

Checks with primary/backup IP, which one is active and the last state (gmon majority of 3 locations).

API: GET /cloud/dns/{id}/health-checks (operationId typed_get_dns__id__health_checks, scope services_read)

```
gyro dns health list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro dns health](gyro_dns_health.md)	 - List health-checked records

