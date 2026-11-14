## gyro container delete

Delete a container environment (prepaid)

### Synopsis

Delete a container environment (prepaid)

Same as the panel's Spending > Delete: hourly charges stop at once and the environment is deleted right away with every service, node, registry image, pipeline and custom domain binding (no retention period). Needs `"confirm": true` (body or query) and an `Idempotency-Key` header; repeating returns 202 `deleted`. Environments billed by invoice answer 409 `not_prepaid`. Terraform: destroying `gyro_container_environment` calls this.

API: DELETE /cloud/containers/{id} (operationId deleteContainerEnvironment, scope cloud_delete)

```
gyro container delete <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
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

* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains

