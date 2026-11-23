## gyro container log-retention set

Change the log package

### Synopsis

Change the log package

`days` 7 | 30 | 90 | 365 and `gb` (monthly size) from `GET .../logs/plans`; `{"free": true}` goes back to the free package. A retention SHORTER than the current one deletes the older logs permanently (`data_loss`, `deletes_logs_older_than_days` and a `warning` text in the answer; `GET .../logs/plans` marks such packages `data_loss`). Every switch from free to a paid package is billed at least 6 hours (`minimum_hours`); going back to free earlier charges the remaining hours at once (`minimum_completion_hours`, shown as "minimum completion" in the hourly detail and CSV, like an additional IPv4). A paid package is billed hourly from the prepaid balance (billing line "Log retention") and needs `"confirm": true`. 409/422 when the platform has no room for the package. Logs are what your containers write to stdout / stderr. Every query is built by the platform with your environment fixed; there is no raw query language. **Free package**: 7 days, 5 GB per month. Paid packages (30 / 90 / 365 days x 10 / 50 / 100 / 300 GB per month, prepaid hourly; `GET /cloud/containers/{id}/logs/plans`) keep logs longer. Over the monthly size there is no extra charge: the stored volume is capped at GB x days / 30, the oldest logs are deleted earlier and you get the `log_quota` notification at 80 % and 100 %. When the environment is cancelled its logs are deleted at once. Best effort service, no compensation for lost log lines.

API: PUT /cloud/containers/{id}/logs/retention (operationId setContainerLogPackage, scope services_write)

```
gyro container log-retention set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --days int                 days (7|30|90|365)
      --dry-run                  show the price quote / request without changing anything
      --free                     free
      --gb int                   gb
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro container log-retention](gyro_container_log-retention.md)	 - Change the log package

