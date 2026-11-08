## gyro mail dns

Required DNS records and their status

```
gyro mail dns [flags]
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

* [gyro mail](gyro_mail.md)	 - Corporate mail: mailboxes, aliases, forwarding ...
* [gyro mail dns create](gyro_mail_dns_create.md)	 - create (fields with --set key=value)
* [gyro mail dns delete](gyro_mail_dns_delete.md)	 - delete (fields with --set key=value)
* [gyro mail dns list](gyro_mail_dns_list.md)	 - Show
* [gyro mail dns update](gyro_mail_dns_update.md)	 - update (fields with --set key=value)

