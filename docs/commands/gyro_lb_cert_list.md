## gyro lb cert list

List certificates

### Synopsis

List certificates

Uploaded and Let's Encrypt certificates (Let's Encrypt ones renew automatically), a pending Let's Encrypt request and the validation redirect base URL.

API: GET /cloud/load-balancers/{id}/certificates (operationId typed_get_load_balancers__id__certificates, scope services_read)

```
gyro lb cert list <ID> [flags]
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

* [gyro lb cert](gyro_lb_cert.md)	 - List certificates

