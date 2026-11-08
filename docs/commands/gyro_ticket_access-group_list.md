## gyro ticket access-group list

List ticket access groups

### Synopsis

List ticket access groups

List ticket access groups. Support ticket of the calling account. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/destek-isteme).

API: GET /tickets/access-groups (operationId coreGetTicketAccessGroups, scope Tickets/GetTicketAccessGroups)

```
gyro ticket access-group list [flags]
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

* [gyro ticket access-group](gyro_ticket_access-group.md)	 - List ticket access groups

