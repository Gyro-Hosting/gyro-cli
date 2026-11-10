## gyro security-group server

Security groups attached to a server

```
gyro security-group server [flags]
```

### Options

```
  -h, --help   help for server
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

* [gyro security-group](gyro_security-group.md)	 - List security groups
* [gyro security-group server list](gyro_security-group_server_list.md)	 - Security groups attached to a server
* [gyro security-group server set](gyro_security-group_server_set.md)	 - Set the exact list of attached groups

