## gyro container log-plans

Log packages

### Synopsis

Log packages

Packages with monthly price (EUR, TL at the current rate), list price and discount. The free package (7 days, 5 GB) is always available.

API: GET /cloud/containers/{id}/logs/plans (operationId listContainerLogPlans, scope services_read)

```
gyro container log-plans <ID> [flags]
```

### Options

```
  -h, --help   help for log-plans
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

