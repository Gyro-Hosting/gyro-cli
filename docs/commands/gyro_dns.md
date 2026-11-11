## gyro dns

DNS records of your domains

```
gyro dns [flags]
```

### Options

```
  -h, --help   help for dns
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro dns acme-challenge](gyro_dns_acme-challenge.md)	 - Add an ACME challenge TXT value
* [gyro dns acme-key](gyro_dns_acme-key.md)	 - List ACME keys
* [gyro dns dnssec](gyro_dns_dnssec.md)	 - Disable DNSSEC
* [gyro dns health](gyro_dns_health.md)	 - List health-checked records
* [gyro dns record](gyro_dns_record.md)	 - List, add, change and delete DNS records
* [gyro dns rrset](gyro_dns_rrset.md)	 - List record sets
* [gyro dns zone](gyro_dns_zone.md)	 - List DNS zones

