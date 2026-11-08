## gyro dns record

List, add, change and delete DNS records

```
gyro dns record [flags]
```

### Options

```
  -h, --help   help for record
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
* [gyro dns record add](gyro_dns_record_add.md)	 - Create dns record (path: {domain})
* [gyro dns record delete](gyro_dns_record_delete.md)	 - Delete dns record (path: {domain})
* [gyro dns record list](gyro_dns_record_list.md)	 - List dns records (path: {domain})
* [gyro dns record update](gyro_dns_record_update.md)	 - Update dns record (path: {domain})

