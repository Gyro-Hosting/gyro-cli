## gyro lb metrics

Load balancer metrics

### Synopsis

Load balancer metrics

Open connections, new connections/s, requests/s, bandwidth in/out (Mbit/s).

API: GET /cloud/load-balancers/{id}/metrics (operationId typed_get_load_balancers__id__metrics, scope services_read)

```
gyro lb metrics <ID> [flags]
```

### Options

```
  -h, --help            help for metrics
      --period string   24h | 7d | 30d
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

