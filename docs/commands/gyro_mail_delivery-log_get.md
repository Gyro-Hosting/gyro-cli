## gyro mail delivery-log get

GET delivery-log (GyroMail)

### Synopsis

GET delivery-log (GyroMail)

Typed shortcut of the generic bridge: tool `delivery-log`. Body fields are the tool's own fields.

API: GET /cloud/mail/{id}/delivery-log (operationId typed_get_mail__id__delivery_log, scope services_read)

```
gyro mail delivery-log get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro mail delivery-log](gyro_mail_delivery-log.md)	 - GET delivery-log (GyroMail)

