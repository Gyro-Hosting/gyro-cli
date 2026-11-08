## gyro mail quarantine get

GET quarantine (GyroMail)

### Synopsis

GET quarantine (GyroMail)

Typed shortcut of the generic bridge: tool `quarantine`. Body fields are the tool's own fields.

API: GET /cloud/mail/{id}/quarantine (operationId typed_get_mail__id__quarantine, scope services_read)

```
gyro mail quarantine get <ID> [flags]
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

* [gyro mail quarantine](gyro_mail_quarantine.md)	 - GET quarantine (GyroMail)

