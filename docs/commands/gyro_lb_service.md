## gyro lb service

List services (ports) of a load balancer

```
gyro lb service [flags]
```

### Options

```
  -h, --help   help for service
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
* [gyro lb service add](gyro_lb_service_add.md)	 - Add a service
* [gyro lb service delete](gyro_lb_service_delete.md)	 - Delete a service
* [gyro lb service list](gyro_lb_service_list.md)	 - List services (ports) of a load balancer
* [gyro lb service update](gyro_lb_service_update.md)	 - Update a service

