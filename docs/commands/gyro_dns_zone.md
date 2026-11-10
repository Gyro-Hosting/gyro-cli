## gyro dns zone

List DNS zones

```
gyro dns zone [flags]
```

### Options

```
  -h, --help   help for zone
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

* [gyro dns](gyro_dns.md)	 - DNS records of your domains
* [gyro dns zone create](gyro_dns_zone_create.md)	 - Add a zone
* [gyro dns zone delete](gyro_dns_zone_delete.md)	 - Delete a zone
* [gyro dns zone export](gyro_dns_zone_export.md)	 - Export the zone file
* [gyro dns zone get](gyro_dns_zone_get.md)	 - Get a zone with its records
* [gyro dns zone import](gyro_dns_zone_import.md)	 - Import a BIND zone file
* [gyro dns zone list](gyro_dns_zone_list.md)	 - List DNS zones
* [gyro dns zone verify](gyro_dns_zone_verify.md)	 - Verify ownership now

