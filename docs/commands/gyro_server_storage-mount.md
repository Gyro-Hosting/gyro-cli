## gyro server storage-mount

List object storage mounts of a server

```
gyro server storage-mount [flags]
```

### Options

```
  -h, --help   help for storage-mount
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
* [gyro server storage-mount command](gyro_server_storage-mount_command.md)	 - New install command (rotates the mount key)
* [gyro server storage-mount create](gyro_server_storage-mount_create.md)	 - Mount a bucket on the server
* [gyro server storage-mount get](gyro_server_storage-mount_get.md)	 - Get one storage mount
* [gyro server storage-mount list](gyro_server_storage-mount_list.md)	 - List object storage mounts of a server
* [gyro server storage-mount remove](gyro_server_storage-mount_remove.md)	 - Remove a storage mount

