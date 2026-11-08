## gyro domain whois

Get whois contacts (path: {domain})

```
gyro domain whois [flags]
```

### Options

```
  -h, --help   help for whois
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
* [gyro domain whois apply](gyro_domain_whois_apply.md)	 - Apply whois profile (path: {domain})
* [gyro domain whois get](gyro_domain_whois_get.md)	 - Get whois contacts (path: {domain})
* [gyro domain whois set](gyro_domain_whois_set.md)	 - Update whois contacts (path: {domain})

