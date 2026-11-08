## gyro domain auth-code

Send auth code (path: {domain})

```
gyro domain auth-code [flags]
```

### Options

```
  -h, --help   help for auth-code
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
* [gyro domain auth-code send](gyro_domain_auth-code_send.md)	 - Send auth code (path: {domain})
* [gyro domain auth-code set](gyro_domain_auth-code_set.md)	 - Save auth code (path: {domain})

