## gyro container schedule preview

Next run times of a timing

### Synopsis

Next run times of a timing

Checks `cron` / `rate_minutes` / `at` + `timezone` and returns the next `count` (1-20, default 5) run times in that time zone. Nothing is stored.

API: POST /cloud/containers/{id}/schedules/preview (operationId previewContainerSchedule, scope services_write)

```
gyro container schedule preview <ID> [flags]
```

### Options

```
      --at timezone              RFC 3339, or local date-time in timezone
      --count int                count
      --cron string              cron
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for preview
      --idempotency-key string   idempotency key (automatic for money operations)
      --rate-minutes int         rate_minutes
      --set stringArray          extra body field key=value (repeatable)
      --timezone string          timezone
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

