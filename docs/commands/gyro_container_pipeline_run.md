## gyro container pipeline run

Start a run

### Synopsis

Start a run

`ref`: branch (`main`), tag (`tag:v1.2.0`) or full ref; `sha`: optional 40 character commit. Without both the first branch of the rule builds. Queued runs start when the environment's parallel build limit allows.

API: POST /cloud/containers/{id}/pipelines/{pipeline}/runs (operationId pipeline_run, scope services_write)

```
gyro container pipeline run <ID> <PIPELINE> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for run
      --idempotency-key string   idempotency key (automatic for money operations)
      --ref string               ref
      --set stringArray          extra body field key=value (repeatable)
      --sha string               sha
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

