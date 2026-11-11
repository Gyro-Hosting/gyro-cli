## gyro container service rollback

Roll back to an earlier revision

### Synopsis

Roll back to an earlier revision

Restores image, size, command and variables of `revision` (default: the previous one) as a new revision. Replicas and secrets stay as they are. Needs `"confirm": true`.

API: POST /cloud/containers/{id}/services/{name}/rollback (operationId typed_post_containers__id__services__name__rollback, scope services_write)

```
gyro container service rollback <ID> <NAME> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for rollback
      --idempotency-key string   idempotency key (automatic for money operations)
      --revision int             Revision number (empty = previous)
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

