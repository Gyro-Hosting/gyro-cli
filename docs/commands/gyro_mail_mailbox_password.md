## gyro mail mailbox password

POST mailboxes:password (GyroMail)

### Synopsis

POST mailboxes:password (GyroMail)

Typed shortcut of the generic bridge: tool `mailboxes`, action `password`. Body fields are the tool's own fields.

API: POST /cloud/mail/{id}/mailboxes/{local}/password (operationId typed_post_mail__id__mailboxes__local__password, scope services_write)

```
gyro mail mailbox password <ID> <LOCAL> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for password
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

* [gyro mail mailbox](gyro_mail_mailbox.md)	 - Mailboxes

