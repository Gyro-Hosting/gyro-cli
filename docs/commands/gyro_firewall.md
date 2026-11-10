## gyro firewall

Network firewall policy (rules, default action, status)

```
gyro firewall [flags]
```

### Options

```
  -h, --help   help for firewall
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
* [gyro firewall confirm](gyro_firewall_confirm.md)	 - Make the running firewall trial permanent
* [gyro firewall delete](gyro_firewall_delete.md)	 - Remove your firewall rules
* [gyro firewall get](gyro_firewall_get.md)	 - Network firewall policy (rules, default action, status)
* [gyro firewall set](gyro_firewall_set.md)	 - Replace the network firewall policy

