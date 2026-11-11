## gyro container external-registry set

External private registry credentials

### Synopsis

External private registry credentials

One private registry per environment (e.g. ghcr.io + a pull-only token). The token is stored encrypted and never returned.

API: PUT /cloud/containers/{id}/external-registry (operationId typed_put_containers__id__external_registry, scope services_write)

```
gyro container external-registry set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --password string          password
      --server string            server
      --set stringArray          extra body field key=value (repeatable)
      --username string          username
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

* [gyro container external-registry](gyro_container_external-registry.md)	 - Remove external registry credentials

