## gyro container logs

Service logs

### Synopsis

Service logs

Lines `{time, instance, text}` oldest first. `since_minutes` (default 60, max 10080), `filter` (plain text), `limit` (max 2000), `previous=1` = last crashed instance.

API: GET /cloud/containers/{id}/services/{name}/logs (operationId typed_get_containers__id__services__name__logs, scope services_read)

```
gyro container logs <ID> <NAME> [flags]
```

### Options

```
      --filter string          Plain text filter
  -h, --help                   help for logs
      --param-limit string     Max lines
      --previous string        1 = previous (crashed) instance
      --since-minutes string   Time range in minutes
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

