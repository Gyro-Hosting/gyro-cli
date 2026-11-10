## gyro security-group server set

Set the exact list of attached groups

### Synopsis

Set the exact list of attached groups

Attaches / detaches in one change (order = evaluation order, max 5). Answers 422 err_busy while the previous change of the server is being applied: retry after a few seconds.

API: PUT /cloud/servers/{id}/security-groups (operationId setServerSecurityGroups, scope cloud_write)

```
gyro security-group server set <ID> [flags]
```

### Options

```
      --data string               request body as JSON, @file or @- (stdin)
      --dry-run                   show the price quote / request without changing anything
  -h, --help                      help for set
      --idempotency-key string    idempotency key (automatic for money operations)
      --security-groups strings   security_groups
      --set stringArray           extra body field key=value (repeatable)
      --wait                      wait until the server is no longer busy
      --wait-timeout duration     maximum time for --wait (default 15m0s)
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

* [gyro security-group server](gyro_security-group_server.md)	 - Security groups attached to a server

