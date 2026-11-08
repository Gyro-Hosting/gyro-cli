## gyro domain forwarding

Get url forwarding (path: {domain})

```
gyro domain forwarding [flags]
```

### Options

```
  -h, --help   help for forwarding
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
* [gyro domain forwarding delete](gyro_domain_forwarding_delete.md)	 - Delete url forwarding (path: {domain})
* [gyro domain forwarding get](gyro_domain_forwarding_get.md)	 - Get url forwarding (path: {domain})
* [gyro domain forwarding set](gyro_domain_forwarding_set.md)	 - Update url forwarding (path: {domain})

