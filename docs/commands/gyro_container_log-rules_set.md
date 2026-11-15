## gyro container log-rules set

Replace the log alert rules

### Synopsis

Replace the log alert rules

Sends the whole list (at most 10; empty list removes all). Keep each rule's `id` to update it in place. 409 `limit_log_rules`, 422 with the rule number on invalid input.

API: PUT /cloud/containers/{id}/log-rules (operationId setContainerLogRules, scope services_write)

```
gyro container log-rules set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --rules strings            rules
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

* [gyro container log-rules](gyro_container_log-rules.md)	 - Log alert rules

