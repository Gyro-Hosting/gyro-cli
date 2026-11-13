## gyro container pipeline list

List pipelines

### Synopsis

List pipelines

Pipelines of the environment with the last run, plan limits (`concurrent`, `monthly_minutes`, `max_timeout`, `max_cpu_m`, `max_memory_mi`, `max_pipelines`) and this month's build minutes.

API: GET /cloud/containers/{id}/pipelines (operationId pipeline_list, scope services_read)

```
gyro container pipeline list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

