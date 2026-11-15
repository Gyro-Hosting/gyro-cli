## gyro container log-rules list

Log alert rules

### Synopsis

Log alert rules

At most 10 rules per environment. A rule fires when the text / level shows up at least `threshold` times within `window_minutes`; the `log_rule` notification goes to the Alerts channels (2 checks in a row, repeat suppression, resolved message).

API: GET /cloud/containers/{id}/log-rules (operationId listContainerLogRules, scope services_read)

```
gyro container log-rules list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro container log-rules](gyro_container_log-rules.md)	 - Log alert rules

