## gyro mail deleted get

GET deleted-mail (GyroMail)

### Synopsis

GET deleted-mail (GyroMail)

Typed shortcut of the generic bridge: tool `deleted-mail`. Body fields are the tool's own fields.

API: GET /cloud/mail/{id}/deleted (operationId typed_get_mail__id__deleted, scope services_read)

```
gyro mail deleted get <ID> [flags]
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

* [gyro mail deleted](gyro_mail_deleted.md)	 - GET deleted-mail (GyroMail)

