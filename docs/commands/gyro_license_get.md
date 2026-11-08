## gyro license get

GET method license_status (LisansAPP)

### Synopsis

GET method license_status (LisansAPP)

Typed shortcut of the generic bridge: tool `license_status`. Body fields are the tool's own fields.

API: GET /cloud/licenses/{id} (operationId typed_get_licenses__id, scope services_read)

```
gyro license get <ID> [flags]
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

* [gyro license](gyro_license.md)	 - Software licenses (cPanel, Plesk, LiteSpeed ...)

