## gyro container service autoscale

Autoscaling / alert threshold of a service

### Synopsis

Autoscaling / alert threshold of a service

Body `{"autoscaling": {"enabled": true, "min": 1, "max": 5, "cpu_percent": 70, "memory_percent": 0}, "alerts": {"min_ready": 2}}`. Replicas are added when the average use exceeds the target (at most double per 30 s) and removed one per minute after 3 calm minutes. Max cannot exceed the plan replica limit; replicas that do not fit in the node pool wait (`warnings: max_exceeds_capacity`, alert `capacity`). Not a revision. Switching off keeps the current replica count (or `replicas`). Refused with 422 `invalid_replicas` (`details.max` = your plan's limit) when the count is outside 0..max copies per service.

API: PUT /cloud/containers/{id}/services/{name}/autoscaling (operationId typed_put_containers__id__services__name__autoscaling, scope services_write)

```
gyro container service autoscale <ID> <NAME> [flags]
```

### Options

```
      --alerts string               Alert threshold: notify when fewer replicas are ready (0 = off). Channels: PUT /cloud/containers/{id}/notifications. (JSON)
      --autoscaling enabled:false   Autoscaling (horizontal): replicas follow the average CPU (and optionally memory) use of the reserved size. enabled:false = off (fixed replicas). (JSON)
      --data string                 request body as JSON, @file or @- (stdin)
      --dry-run                     show the price quote / request without changing anything
  -h, --help                        help for autoscale
      --idempotency-key string      idempotency key (automatic for money operations)
      --replicas int                Fixed replicas when switching autoscaling off
      --set stringArray             extra body field key=value (repeatable)
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

