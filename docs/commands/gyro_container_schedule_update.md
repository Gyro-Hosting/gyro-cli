## gyro container schedule update

Change a scheduled job (only the given flags change)

```
gyro container schedule update <ENVIRONMENT_ID> <NAME|ID> [flags] [-- COMMAND...]
```

### Examples

```
  gyro container schedule update 17000 nightly-report --cron "30 2 * * *" --timezone UTC
  gyro container schedule update 17000 nightly-report --service api -- npm run report:full
  gyro container schedule update 17000 ping-app --header "Authorization: Bearer $NEW"
```

### Options

```
      --at string               run once at this time: 2026-12-01T03:00 (in --timezone) or RFC 3339
      --body string             HTTP body (text, or @file)
      --clear-headers           remove the stored HTTP headers
      --cpu string              vCPU of a run (default: same as the service)
      --cron string             cron expression: minute hour day month weekday, e.g. "0 3 * * *" (daily 03:00), "*/15 * * * *", "0 9 * * 1-5"; @daily / @hourly work too
      --description string      description
      --every duration          run every N minutes instead of a cron expression, e.g. 15m, 2h (at least 1m; Serverless 5m)
      --header stringArray      HTTP header "Name: value" (repeatable; values are kept secret)
  -h, --help                    help for update
      --http-timeout duration   time limit of the HTTP call (1s-120s, default 30s)
      --memory int              memory in MB of a run (default: same as the service)
      --method string           HTTP method (default GET)
      --retries int             retries after a failed run, 0-3 (30 s × attempt apart)
      --service string          run the command after "--" with this service's image, variables and secrets
      --time-limit duration     stop a run after this time (10s-6h, default 30m)
      --timezone string         IANA time zone of the schedule (default Europe/Istanbul), e.g. UTC
      --url string              call this URL instead (http:// or https://)
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

