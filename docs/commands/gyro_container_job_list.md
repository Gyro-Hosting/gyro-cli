## gyro container job list

List one-off jobs

### Synopsis

List one-off jobs

Last 50 jobs, newest first, with `active` (running now) and `limit` (parallel jobs). `?service=<name>` filters.

API: GET /cloud/containers/{id}/jobs (operationId listContainerJobs, scope services_read)

```
gyro container job list <ID> [flags]
```

### Options

```
  -h, --help             help for list
      --service string   only the jobs of this service
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

* [gyro container job](gyro_container_job.md)	 - List one-off jobs

