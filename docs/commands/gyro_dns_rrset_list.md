## gyro dns rrset list

List record sets

### Synopsis

List record sets

All record sets (`rrsets`: name, type, ttl, records, read_only, health_check). `source=cpanel` reads a zone of your cPanel hosting (read only).

API: GET /cloud/dns/{id}/zones/{zone}/records (operationId typed_get_dns__id__zones__zone__records, scope services_read)

```
gyro dns rrset list <ID> <ZONE> [flags]
```

### Options

```
  -h, --help            help for list
      --source string   gyrodns (default) | cpanel
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

* [gyro dns rrset](gyro_dns_rrset.md)	 - List record sets

