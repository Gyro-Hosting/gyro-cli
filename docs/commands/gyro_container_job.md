## gyro container job

List one-off jobs

```
gyro container job [flags]
```

### Options

```
  -h, --help   help for job
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
* [gyro container job cancel](gyro_container_job_cancel.md)	 - Stop a one-off job
* [gyro container job get](gyro_container_job_get.md)	 - Get a one-off job
* [gyro container job list](gyro_container_job_list.md)	 - List one-off jobs
* [gyro container job logs](gyro_container_job_logs.md)	 - Output of a one-off job

