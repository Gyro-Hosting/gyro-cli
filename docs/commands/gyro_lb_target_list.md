## gyro lb target list

List targets with health

### Synopsis

List targets with health

Targets, per-service health (healthy/unhealthy/unknown), and the cloud servers / dedicated IPs that can be added.

API: GET /cloud/load-balancers/{id}/targets (operationId typed_get_load_balancers__id__targets, scope services_read)

```
gyro lb target list <ID> [flags]
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

* [gyro lb target](gyro_lb_target.md)	 - List targets with health

