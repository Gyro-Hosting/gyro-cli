## gyro container log-rules

Log alert rules

```
gyro container log-rules [flags]
```

### Options

```
  -h, --help   help for log-rules
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
* [gyro container log-rules list](gyro_container_log-rules_list.md)	 - Log alert rules
* [gyro container log-rules set](gyro_container_log-rules_set.md)	 - Replace the log alert rules
* [gyro container log-rules test](gyro_container_log-rules_test.md)	 - Try a log alert rule

