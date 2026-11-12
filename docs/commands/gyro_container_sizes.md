## gyro container sizes

Serverless container sizes (vCPU / memory pairs) and prices

### Synopsis

Serverless container sizes (vCPU / memory pairs) and prices

Gyro Containers Serverless bills each running container copy per hour: `cpu_m/1000` x vCPU-hour + `memory_mi/1024` x GB-hour, from the prepaid balance, no monthly cap. Only the `cpu_m` / `memory_mi` pairs listed here are accepted on `POST /cloud/containers/{id}/services` of a serverless environment (422 `invalid_cpu` / `invalid_memory`). 404 `serverless_not_available` when the product is not offered.

API: GET /cloud/billing/container-sizes (operationId getContainerSizes, scope billing_read)

```
gyro container sizes [flags]
```

### Options

```
  -h, --help   help for sizes
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

* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains

