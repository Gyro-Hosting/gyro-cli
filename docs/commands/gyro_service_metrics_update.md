## gyro service metrics update

Update service metric (path: {id}, {key})

### Synopsis

Update service metric (path: {id}, {key})

Update service metric (path: {id}, {key}). Service of the calling account (any product line). Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/sahip-olunan-hizmetler).

API: PATCH /services/{id}/metrics/{key} (operationId coreUpdateServiceMetric, scope Services/UpdateServiceMetric)

```
gyro service metrics update <ID> <KEY> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --enabled                  İstenen durum.
  -h, --help                     help for update
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro service metrics](gyro_service_metrics.md)	 - Get service metrics (path: {id})

