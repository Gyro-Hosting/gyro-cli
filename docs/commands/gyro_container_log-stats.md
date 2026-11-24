## gyro container log-stats

Log statistics

### Synopsis

Log statistics

Line counts over time (`series`, at most 60 buckets) and the top values of one grouping: `by=service` (default), `level`, or a field of JSON / logfmt lines (`by=status`). Same filters as the search; window at most 7 days.

API: GET /cloud/containers/{id}/logs/stats (operationId containerLogStats, scope services_read)

```
gyro container log-stats <ID> [flags]
```

### Options

```
      --by string           service | level | <field>
      --end string          RFC3339 end (default now)
      --exclude string      line does not contain this text
      --filter string       line contains this text (case-insensitive, max 200)
  -h, --help                help for log-stats
      --level string        error | warn | info | debug (level word in the line)
      --parse where         json (default when where is given) | logfmt
      --service string      only this service (default: all services of the environment)
      --since-minutes int   window ending now (default 60); limited to the package retention
      --start string        RFC3339 start (instead of since_minutes); one query spans at most 31 days
      --top int             groups (default 10)
      --where ;             field conditions on JSON / logfmt lines, separated by ;: `status>=500; path~/api; method=POST` (ops = != > >= < <= =~ ~)
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

* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains

