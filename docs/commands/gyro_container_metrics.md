## gyro container metrics

Service CPU / memory

### Synopsis

Service CPU / memory

`source=history`: series (cpu_m, memory_mi, restarts) for `hours`; `source=current`: current usage only.

API: GET /cloud/containers/{id}/services/{name}/metrics (operationId typed_get_containers__id__services__name__metrics, scope services_read)

```
gyro container metrics <ID> <NAME> [flags]
```

### Options

```
  -h, --help           help for metrics
      --hours string   History length in hours (default 24)
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

