## gyro security-group

List security groups

```
gyro security-group [flags]
```

### Options

```
  -h, --help   help for security-group
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
* [gyro security-group attach](gyro_security-group_attach.md)	 - Attach a security group to a server
* [gyro security-group confirm](gyro_security-group_confirm.md)	 - Make the running rule trial permanent
* [gyro security-group create](gyro_security-group_create.md)	 - Create a security group
* [gyro security-group delete](gyro_security-group_delete.md)	 - Delete a security group
* [gyro security-group detach](gyro_security-group_detach.md)	 - Detach a security group from a server
* [gyro security-group get](gyro_security-group_get.md)	 - Get a security group
* [gyro security-group list](gyro_security-group_list.md)	 - List security groups
* [gyro security-group rules](gyro_security-group_rules.md)	 - Replace the rules of a security group
* [gyro security-group server](gyro_security-group_server.md)	 - Security groups attached to a server
* [gyro security-group update](gyro_security-group_update.md)	 - Rename / describe a security group

