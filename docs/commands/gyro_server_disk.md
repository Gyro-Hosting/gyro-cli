## gyro server disk

List the server's additional disks

```
gyro server disk [flags]
```

### Options

```
  -h, --help   help for disk
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

* [gyro server](gyro_server.md)	 - Cloud servers: create, power, snapshots, SSH ...
* [gyro server disk add](gyro_server_disk_add.md)	 - Add an additional disk
* [gyro server disk list](gyro_server_disk_list.md)	 - List the server's additional disks
* [gyro server disk remove](gyro_server_disk_remove.md)	 - Remove an additional disk (with its data)
* [gyro server disk resize](gyro_server_disk_resize.md)	 - Grow an additional disk

