## gyro domain whois-profile

List whois profiles

```
gyro domain whois-profile [flags]
```

### Options

```
  -h, --help   help for whois-profile
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
* [gyro domain whois-profile create](gyro_domain_whois-profile_create.md)	 - Create whois profile
* [gyro domain whois-profile default](gyro_domain_whois-profile_default.md)	 - Set default whois profile (path: {pid})
* [gyro domain whois-profile delete](gyro_domain_whois-profile_delete.md)	 - Delete whois profile (path: {pid})
* [gyro domain whois-profile list](gyro_domain_whois-profile_list.md)	 - List whois profiles
* [gyro domain whois-profile set](gyro_domain_whois-profile_set.md)	 - Update whois profile (path: {pid})

