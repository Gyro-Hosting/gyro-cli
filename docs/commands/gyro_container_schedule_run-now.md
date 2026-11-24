## gyro container schedule run-now

Run a scheduled job now

### Synopsis

Run a scheduled job now

Starts one run now (trigger `manual`), whether paused or not; refused with `schedule_running` while a run of the schedule is going. Returns the run and its job id.

API: POST /cloud/containers/{id}/schedules/{sched}/run (operationId runContainerScheduleNow, scope services_write)

```
gyro container schedule run-now <ID> <SCHED> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for run-now
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro container schedule](gyro_container_schedule.md)	 - List scheduled jobs with their next and last run

