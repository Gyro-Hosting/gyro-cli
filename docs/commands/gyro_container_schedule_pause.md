## gyro container schedule pause

Pause a scheduled job

### Synopsis

Pause a scheduled job

No run starts until resumed (a run in progress finishes).

API: POST /cloud/containers/{id}/schedules/{sched}/pause (operationId pauseContainerSchedule, scope services_write)

```
gyro container schedule pause <ID> <SCHED> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for pause
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro container schedule](gyro_container_schedule.md)	 - List scheduled jobs with their next and last run

