## gyro dns zone get

Get a zone with its records

### Synopsis

Get a zone with its records

Zone state, serial, DNSSEC flag (+DS) and all record sets. Pending zones return the verification details.

API: GET /cloud/dns/{id}/zones/{zone} (operationId typed_get_dns__id__zones__zone, scope services_read)

```
gyro dns zone get <ID> <ZONE> [flags]
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

* [gyro dns zone](gyro_dns_zone.md)	 - List DNS zones

