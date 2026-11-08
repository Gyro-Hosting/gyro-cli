## gyro mail migration run

POST sync-jobs:run (GyroMail)

### Synopsis

POST sync-jobs:run (GyroMail)

Typed shortcut of the generic bridge: tool `sync-jobs`, action `run`. Body fields are the tool's own fields. Needs `"confirm": true`.

API: POST /cloud/mail/{id}/migrations/{jid}/run (operationId typed_post_mail__id__migrations__jid__run, scope services_write)

```
gyro mail migration run <ID> <JID> [flags]
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

* [gyro mail migration](gyro_mail_migration.md)	 - IMAP migration jobs

