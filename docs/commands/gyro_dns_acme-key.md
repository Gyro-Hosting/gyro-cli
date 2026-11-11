## gyro dns acme-key

List ACME keys

```
gyro dns acme-key [flags]
```

### Options

```
  -h, --help   help for acme-key
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
* [gyro dns acme-key create](gyro_dns_acme-key_create.md)	 - Create an ACME key
* [gyro dns acme-key delete](gyro_dns_acme-key_delete.md)	 - Delete an ACME key
* [gyro dns acme-key list](gyro_dns_acme-key_list.md)	 - List ACME keys

