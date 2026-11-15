## gyro container schedule

List scheduled jobs with their next and last run

```
gyro container schedule [flags]
```

### Options

```
  -h, --help   help for schedule
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
* [gyro container schedule create](gyro_container_schedule_create.md)	 - Create a scheduled job (cron, every N minutes or once)
* [gyro container schedule delete](gyro_container_schedule_delete.md)	 - Delete a scheduled job
* [gyro container schedule get](gyro_container_schedule_get.md)	 - Get a scheduled job
* [gyro container schedule list](gyro_container_schedule_list.md)	 - List scheduled jobs with their next and last run
* [gyro container schedule pause](gyro_container_schedule_pause.md)	 - Pause a scheduled job
* [gyro container schedule preview](gyro_container_schedule_preview.md)	 - Next run times of a timing
* [gyro container schedule resume](gyro_container_schedule_resume.md)	 - Resume a scheduled job
* [gyro container schedule run-now](gyro_container_schedule_run-now.md)	 - Run a scheduled job now
* [gyro container schedule update](gyro_container_schedule_update.md)	 - Change a scheduled job (only the given flags change)

