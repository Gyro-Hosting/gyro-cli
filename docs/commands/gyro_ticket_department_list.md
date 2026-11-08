## gyro ticket department list

List ticket departments

### Synopsis

List ticket departments

List ticket departments. Support ticket of the calling account. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/destek-isteme).

API: GET /tickets/departments (operationId coreGetTicketDepartments, scope Tickets/GetTicketDepartments)

```
gyro ticket department list [flags]
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

* [gyro ticket department](gyro_ticket_department.md)	 - List ticket departments

