## gyro container pipeline redeploy

Deploy this run's image again (rollback)

### Synopsis

Deploy this run's image again (rollback)

Starts a new run that deploys the image of an earlier run to the target services (with health watch). Needs `"confirm": true`.

API: POST /cloud/containers/{id}/pipeline-runs/{run}/redeploy (operationId pipeline_redeploy, scope services_write)

```
gyro container pipeline redeploy <ID> <RUN> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for redeploy
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
      --targets strings          targets
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

* [gyro container pipeline](gyro_container_pipeline.md)	 - List pipelines

