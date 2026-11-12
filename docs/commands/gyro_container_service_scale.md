## gyro container service scale

Scale a service

### Synopsis

Scale a service

Sets the number of instances (0 = stop, settings kept). Does not create a revision. Refused with 409 `autoscaling_enabled` while autoscaling is on (change min/max instead).

API: POST /cloud/containers/{id}/services/{name}/scale (operationId typed_post_containers__id__services__name__scale, scope services_write)

```
gyro container service scale <ID> <NAME> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for scale
      --idempotency-key string   idempotency key (automatic for money operations)
      --replicas int             replicas
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

* [gyro container service](gyro_container_service.md)	 - List container services

