## gyro service tool run

Run a tool action (create a mailbox, restore a snapshot, ...)

### Synopsis

Run a tool action (create a mailbox, restore a snapshot, ...)

Body fields are the tool's own fields (see the tool data and the per-module table in the guide). Actions flagged `confirm` in the discovery answer need `"confirm": true`.

API: POST /cloud/services/{id}/tools/{tool}/{action} (operationId runServiceToolAction, scope services_write)

```
gyro service tool run <ID> <TOOL> <ACTION> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for run
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

* [gyro service tool](gyro_service_tool.md)	 - Discover tools, actions and safety policy of a service

