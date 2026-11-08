## gyro domain email-forward

List email forwards (path: {domain})

```
gyro domain email-forward [flags]
```

### Options

```
  -h, --help   help for email-forward
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
* [gyro domain email-forward create](gyro_domain_email-forward_create.md)	 - Create email forward (path: {domain})
* [gyro domain email-forward delete](gyro_domain_email-forward_delete.md)	 - Delete email forward (path: {domain})
* [gyro domain email-forward list](gyro_domain_email-forward_list.md)	 - List email forwards (path: {domain})

