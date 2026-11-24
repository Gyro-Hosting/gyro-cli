## gyro container schedule create

Create a scheduled job (cron, every N minutes or once)

### Synopsis

Runs a command of a service (as a one-off job: the image of the running version, its variables and secrets) or
calls a URL on a schedule. Times follow --timezone (default Europe/Istanbul). A run never starts while the previous
run of the same schedule is still going. Failed runs can be retried (--retries 0-3); after the last attempt you get
a schedule_failed notification (and schedule_missed when a run could not start in time). Free on node plans; on
Serverless each run is billed like a task. Limits: shortest interval 1 minute (Serverless 5), schedules per
environment 5-10 by plan (see "gyro limits").

```
gyro container schedule create <ENVIRONMENT_ID> <NAME> (--cron EXPR | --every DURATION | --at TIME) (--service NAME [-- COMMAND...] | --url URL) [flags]
```

### Examples

```
  gyro container schedule create 17000 nightly-report --cron "0 3 * * *" --service api -- npm run report
  gyro container schedule create 17000 queue-clean --every 15m --service worker --retries 2 -- php artisan queue:prune-failed
  gyro container schedule create 17000 ping-app --every 5m --url https://example.com/cron --method POST --header "Authorization: Bearer $TOKEN"
  gyro container schedule create 17000 campaign --at 2026-12-01T09:00 --service api -- node send.js
```

### Options

```
      --at string               run once at this time: 2026-12-01T03:00 (in --timezone) or RFC 3339
      --body string             HTTP body (text, or @file)
      --cpu string              vCPU of a run (default: same as the service)
      --cron string             cron expression: minute hour day month weekday, e.g. "0 3 * * *" (daily 03:00), "*/15 * * * *", "0 9 * * 1-5"; @daily / @hourly work too
      --description string      description
      --every duration          run every N minutes instead of a cron expression, e.g. 15m, 2h (at least 1m; Serverless 5m)
      --header stringArray      HTTP header "Name: value" (repeatable; values are kept secret)
  -h, --help                    help for create
      --http-timeout duration   time limit of the HTTP call (1s-120s, default 30s)
      --memory int              memory in MB of a run (default: same as the service)
      --method string           HTTP method (default GET)
      --paused                  create it paused
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

