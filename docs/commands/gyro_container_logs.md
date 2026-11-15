## gyro container logs

Search and follow the logs of an environment or one service

### Synopsis

Shows what your containers wrote to stdout / stderr, oldest first. Without a service every service of the
environment is searched. Filters: --filter (line contains, case-insensitive), --exclude, --level (error|warn|info|debug)
and --where field conditions on JSON / logfmt lines (status>=500, method=POST, path~/api; repeat the flag for more).
--follow keeps printing new lines (polls every 2 s, Ctrl-C stops). --all pages back through the whole window
(up to 100,000 lines). How far back you can go depends on the log package (free: 7 days; gyro container log-usage).

```
gyro container logs <ENVIRONMENT_ID> [SERVICE] [flags]
```

### Examples

```
  gyro container logs 17000 api --since 30m
  gyro container logs 17000 --level error --since 24h
  gyro container logs 17000 api --where 'status>=500' --where 'path~/api' --since 2h
  gyro container logs 17000 api -f
```

### Options

```
      --all                 page back through the whole window (max 100,000 lines)
      --end string          end time (RFC3339 or a duration ago)
      --exclude string      line does not contain this text
      --filter string       line contains this text (case-insensitive)
  -f, --follow              keep printing new lines (Ctrl-C stops)
  -h, --help                help for logs
      --level string        error | warn | info | debug
      --limit int           lines per page (default 300, max 5000)
      --parse string        json | logfmt
      --previous            output of the last crashed replica of the service
      --service string      only this service (same as the second argument)
      --since string        window ending now: 30m, 6h, 7d (default 1h; 15m with --follow)
      --start string        start time (RFC3339 or a duration ago, e.g. 3d); one query spans at most 31 days
      --where stringArray   field condition on JSON / logfmt lines: status>=500, method=POST, path~/api (repeatable)
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

