## gyro mail autoresponder set

PUT autoresponder:edit (GyroMail)

### Synopsis

PUT autoresponder:edit (GyroMail)

Typed shortcut of the generic bridge: tool `autoresponder`, action `edit`. Body fields are the tool's own fields.

API: PUT /cloud/mail/{id}/autoresponders/{local} (operationId typed_put_mail__id__autoresponders__local, scope services_write)

```
gyro mail autoresponder set <ID> <LOCAL> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
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

* [gyro mail autoresponder](gyro_mail_autoresponder.md)	 - Auto-responders

