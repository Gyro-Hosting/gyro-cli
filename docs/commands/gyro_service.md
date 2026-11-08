## gyro service

All services of every product line (generic bridge)

```
gyro service [flags]
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro service addon](gyro_service_addon.md)	 - Get service addons
* [gyro service cancellation](gyro_service_cancellation.md)	 - Create cancellation
* [gyro service dashboard](gyro_service_dashboard.md)	 - Get service dashboard
* [gyro service get](gyro_service_get.md)	 - Get a service (renewal, cancellation, access details)
* [gyro service invoice](gyro_service_invoice.md)	 - Get service invoices
* [gyro service list](gyro_service_list.md)	 - List services of every product line
* [gyro service method](gyro_service_method.md)	 - Call a client-callable module method (e.g. license change_ip)
* [gyro service metrics](gyro_service_metrics.md)	 - Get service metrics
* [gyro service module-method](gyro_service_module-method.md)	 - Use service module method
* [gyro service password](gyro_service_password.md)	 - Change service password
* [gyro service renew](gyro_service_renew.md)	 - Renew service
* [gyro service sso](gyro_service_sso.md)	 - Get service sso
* [gyro service tool](gyro_service_tool.md)	 - Discover tools, actions and safety policy of a service
* [gyro service update](gyro_service_update.md)	 - Update service

