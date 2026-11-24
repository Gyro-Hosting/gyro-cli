## gyro container job logs

Output of a one-off job

### Synopsis

Output of a one-off job

Live while the job runs (`live: true`), saved afterwards (last 2,000 lines, max 512 KB; secret values of the service masked). `?tail=N` last N lines. Poll every few seconds to follow.

API: GET /cloud/containers/{id}/jobs/{job}/logs (operationId getContainerJobLogs, scope services_read)

```
gyro container job logs <ID> <JOB> [flags]
```

### Options

```
  -h, --help       help for logs
      --tail int   last N lines (default 2000)
```

### Options inherited from parent commands

```
      --account gyro account list   act in another account you are a panel user of (id from gyro account list; "self" = your own; env GYRO_ACCOUNT, profile default: `gyro account use`)
      --api-url string              API base URL (env GYRO_API_URL)
      --debug                       log HTTP requests to stderr (secrets masked)
      --lang string                 message language: tr or en (default from locale)
      --no-color                    disable colours (also NO_COLOR)
      --no-header                   omit the table header
  -o, --output string               output format: table, json, yaml, text
  -p, --profile string              profile to use (env GYRO_PROFILE)
      --query string                JMESPath query applied to the result
  -q, --quiet                       print only ids
      --timeout duration            overall timeout, e.g. 30s, 10m
      --wide                        show all table columns
  -y, --yes                         do not ask for confirmation
```

### SEE ALSO

* [gyro container job](gyro_container_job.md)	 - List one-off jobs

