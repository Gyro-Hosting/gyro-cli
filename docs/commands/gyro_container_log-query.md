## gyro container log-query

Search the logs (JSON body)

### Synopsis

Search the logs (JSON body)

Same as the GET search with the parameters in the body (`where` as a list). Logs are what your containers write to stdout / stderr. Every query is built by the platform with your environment fixed; there is no raw query language. **Free package**: 7 days, 5 GB per month. Paid packages (30 / 90 / 365 days x 10 / 50 / 100 / 300 GB per month, prepaid hourly; `GET /cloud/containers/{id}/logs/plans`) keep logs longer. Over the monthly size there is no extra charge: the stored volume is capped at GB x days / 30, the oldest logs are deleted earlier and you get the `log_quota` notification at 80 % and 100 %. When the environment is cancelled its logs are deleted at once. Best effort service, no compensation for lost log lines.

API: POST /cloud/containers/{id}/logs/query (operationId queryContainerLogs, scope services_write)

```
gyro container log-query <ID> [flags]
```

### Options

```
      --after tail_cursor        live tail cursor (tail_cursor / `cursor`): lines after it, oldest first; poll every 2-3 s
      --before next_cursor       page cursor (next_cursor of the previous page): older lines
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --end string               RFC3339 end (default now)
      --exclude string           line does not contain this text
      --filter string            line contains this text (case-insensitive, max 200)
  -h, --help                     help for log-query
      --idempotency-key string   idempotency key (automatic for money operations)
      --level string             error | warn | info | debug (level word in the line)
      --param-limit int          lines (default 300, max 5000)
      --parse where              json (default when where is given) | logfmt
      --service string           only this service (default: all services of the environment)
      --set stringArray          extra body field key=value (repeatable)
      --since-minutes int        window ending now (default 60); limited to the package retention
      --start string             RFC3339 start (instead of since_minutes); one query spans at most 31 days
      --where ;                  field conditions on JSON / logfmt lines, separated by ;: `status>=500; path~/api; method=POST` (ops = != > >= < <= =~ ~)
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

