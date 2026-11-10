## gyro lb service list

List services (ports) of a load balancer

### Synopsis

List services (ports) of a load balancer

Services with health check and HTTP options, plus the certificates of this load balancer and its limits.

API: GET /cloud/load-balancers/{id}/services (operationId typed_get_load_balancers__id__services, scope services_read)

```
gyro lb service list <ID> [flags]
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

* [gyro lb service](gyro_lb_service.md)	 - List services (ports) of a load balancer

