## gyro server firewall rule

Add an inbound firewall rule

```
gyro server firewall rule [flags]
```

### Options

```
  -h, --help   help for rule
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

* [gyro server firewall](gyro_server_firewall.md)	 - Inbound firewall rules
* [gyro server firewall rule add](gyro_server_firewall_rule_add.md)	 - Add an inbound firewall rule
* [gyro server firewall rule delete](gyro_server_firewall_rule_delete.md)	 - Remove a firewall rule (the last one needs confirm)

