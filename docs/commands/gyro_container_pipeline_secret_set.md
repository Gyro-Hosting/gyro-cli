## gyro container pipeline secret set

Set build secrets

### Synopsis

Set build secrets

`set`: {NAME: value} (A-Z0-9_). Given only to `commands` as environment variables, masked as **** in logs, never returned.

API: PUT /cloud/containers/{id}/pipelines/{pipeline}/secrets (operationId pipeline_secret_set, scope services_write)

```
gyro container pipeline secret set <ID> <PIPELINE> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --param-set string         set (JSON)
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

* [gyro container pipeline secret](gyro_container_pipeline_secret.md)	 - Delete a build secret

