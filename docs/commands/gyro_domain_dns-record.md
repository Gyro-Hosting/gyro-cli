## gyro domain dns-record

Get dns records

```
gyro domain dns-record [flags]
```

### Options

```
  -h, --help   help for dns-record
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

* [gyro domain](gyro_domain.md)	 - Domains: search, register, transfer, nameservers, DNS, WHOIS
* [gyro domain dns-record create](gyro_domain_dns-record_create.md)	 - Create dns record
* [gyro domain dns-record delete](gyro_domain_dns-record_delete.md)	 - Delete dns record
* [gyro domain dns-record list](gyro_domain_dns-record_list.md)	 - Get dns records
* [gyro domain dns-record set](gyro_domain_dns-record_set.md)	 - Update dns record

