## gyro service metrics

Get service metrics (path: {id})

```
gyro service metrics [flags]
```

### Options

```
  -h, --help   help for metrics
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

* [gyro service](gyro_service.md)	 - All services of every product line (generic bridge)
* [gyro service metrics billing](gyro_service_metrics_billing.md)	 - List service metric billing (path: {id})
* [gyro service metrics list](gyro_service_metrics_list.md)	 - Get service metrics (path: {id})
* [gyro service metrics update](gyro_service_metrics_update.md)	 - Update service metric (path: {id}, {key})

