## gyro container pipeline logs

Run log

### Synopsis

Run log

Live while building (clone / commands / build sections, secrets masked), saved afterwards (max 512 KB). `tail`: last N lines.

API: GET /cloud/containers/{id}/pipeline-runs/{run}/logs (operationId pipeline_logs, scope services_read)

```
gyro container pipeline logs <ID> <RUN> [flags]
```

### Options

```
  -h, --help       help for logs
      --tail int   last N lines (default 3000)
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

