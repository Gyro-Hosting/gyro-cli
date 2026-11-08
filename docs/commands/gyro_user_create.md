## gyro user create

Invite a sub-user

### Synopsis

Invite a sub-user

API: POST /cloud/users (operationId createPanelUser, scope users)

```
gyro user create [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --email string             email
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --label string             label
      --permissions strings      e.g. view_services, view_invoices, pay_invoices, view_tickets, open_tickets ...
      --send-invite              send_invite
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

* [gyro user](gyro_user.md)	 - List sub-users

