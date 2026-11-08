## gyro mail two-factor get

GET two-factor (GyroMail)

### Synopsis

GET two-factor (GyroMail)

Typed shortcut of the generic bridge: tool `two-factor`. Body fields are the tool's own fields.

API: GET /cloud/mail/{id}/two-factor (operationId typed_get_mail__id__two_factor, scope services_read)

```
gyro mail two-factor get <ID> [flags]
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

* [gyro mail two-factor](gyro_mail_two-factor.md)	 - GET two-factor (GyroMail)

