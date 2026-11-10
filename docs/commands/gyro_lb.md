## gyro lb

List load balancers

```
gyro lb [flags]
```

### Options

```
  -h, --help   help for lb
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
* [gyro lb algorithm](gyro_lb_algorithm.md)	 - Change the balancing algorithm
* [gyro lb cert](gyro_lb_cert.md)	 - List certificates
* [gyro lb create](gyro_lb_create.md)	 - Order a load balancer
* [gyro lb delete](gyro_lb_delete.md)	 - Delete (cancel) a load balancer
* [gyro lb get](gyro_lb_get.md)	 - Get a load balancer
* [gyro lb list](gyro_lb_list.md)	 - List load balancers
* [gyro lb metrics](gyro_lb_metrics.md)	 - Load balancer metrics
* [gyro lb network](gyro_lb_network.md)	 - Attach to a private network
* [gyro lb rdns](gyro_lb_rdns.md)	 - Set reverse DNS
* [gyro lb service](gyro_lb_service.md)	 - List services (ports) of a load balancer
* [gyro lb settings](gyro_lb_settings.md)	 - Algorithm, reverse DNS, private network
* [gyro lb target](gyro_lb_target.md)	 - List targets with health
* [gyro lb types](gyro_lb_types.md)	 - List load balancer sizes and prices

