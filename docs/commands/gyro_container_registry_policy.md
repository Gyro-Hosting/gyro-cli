## gyro container registry policy

Image clean-up policy

### Synopsis

Image clean-up policy

`enabled` (default true), `keep` 1-50 (images of the last N revisions of every service are kept, default 10), `untagged_days` 1-90 (images used by no revision are deleted after, default 7).

API: PUT /cloud/containers/{id}/registry/policy (operationId typed_put_containers__id__registry_policy, scope services_write)

```
gyro container registry policy <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --enabled                  enabled
  -h, --help                     help for policy
      --idempotency-key string   idempotency key (automatic for money operations)
      --keep int                 keep
      --set stringArray          extra body field key=value (repeatable)
      --untagged-days int        untagged_days
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

* [gyro container registry](gyro_container_registry.md)	 - Gyro image registry of the environment

