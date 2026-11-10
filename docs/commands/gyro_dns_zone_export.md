## gyro dns zone export

Export the zone file

### Synopsis

Export the zone file

Returns the zone in BIND format (`bind`).

API: GET /cloud/dns/{id}/zones/{zone}/export (operationId typed_get_dns__id__zones__zone__export, scope services_read)

```
gyro dns zone export <ID> <ZONE> [flags]
```

### Options

```
  -h, --help   help for export
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

