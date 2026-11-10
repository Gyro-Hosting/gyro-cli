## gyro dns rrset

List record sets

```
gyro dns rrset [flags]
```

### Options

```
  -h, --help   help for rrset
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
* [gyro dns rrset delete](gyro_dns_rrset_delete.md)	 - Delete a record set
* [gyro dns rrset list](gyro_dns_rrset_list.md)	 - List record sets
* [gyro dns rrset set](gyro_dns_rrset_set.md)	 - Create or replace a record set

