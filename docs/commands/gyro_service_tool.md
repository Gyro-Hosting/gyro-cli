## gyro service tool

Discover tools, actions and safety policy of a service

```
gyro service tool [flags]
```

### Options

```
  -h, --help   help for tool
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

* [gyro service](gyro_service.md)	 - All services of every product line (generic bridge)
* [gyro service tool get](gyro_service_tool_get.md)	 - Read a tool (list mailboxes, snapshots, DNS records, ...)
* [gyro service tool list](gyro_service_tool_list.md)	 - Discover tools, actions and safety policy of a service
* [gyro service tool run](gyro_service_tool_run.md)	 - Run a tool action (create a mailbox, restore a snapshot, ...)

