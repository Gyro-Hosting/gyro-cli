## gyro dns dnssec

Disable DNSSEC

```
gyro dns dnssec [flags]
```

### Options

```
  -h, --help   help for dnssec
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
* [gyro dns dnssec disable](gyro_dns_dnssec_disable.md)	 - Disable DNSSEC
* [gyro dns dnssec enable](gyro_dns_dnssec_enable.md)	 - Enable DNSSEC
* [gyro dns dnssec status](gyro_dns_dnssec_status.md)	 - DNSSEC status and DS records

