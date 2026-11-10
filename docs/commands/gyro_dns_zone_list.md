## gyro dns zone list

List DNS zones

### Synopsis

List DNS zones

Zones of this DNS service (`items`, state active|pending with the TXT verification token), read-only zones copied from your cPanel hosting (`cpanel`), plan limits and our name servers.

API: GET /cloud/dns/{id}/zones (operationId typed_get_dns__id__zones, scope services_read)

```
gyro dns zone list <ID> [flags]
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

* [gyro dns zone](gyro_dns_zone.md)	 - List DNS zones

