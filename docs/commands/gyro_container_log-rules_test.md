## gyro container log-rules test

Try a log alert rule

### Synopsis

Try a log alert rule

Counts the matching lines in the rule's window now and says whether it would fire. Nothing is saved.

API: POST /cloud/containers/{id}/log-rules/test (operationId testContainerLogRule, scope services_write)

```
gyro container log-rules test <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for test
      --idempotency-key string   idempotency key (automatic for money operations)
      --rule string              rule (JSON)
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro container log-rules](gyro_container_log-rules.md)	 - Log alert rules

