## gyro server snapshot

List snapshots and the snapshot quota

```
gyro server snapshot [flags]
```

### Options

```
  -h, --help   help for snapshot
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
* [gyro server snapshot create](gyro_server_snapshot_create.md)	 - Take a snapshot
* [gyro server snapshot delete](gyro_server_snapshot_delete.md)	 - Delete a snapshot
* [gyro server snapshot list](gyro_server_snapshot_list.md)	 - List snapshots and the snapshot quota
* [gyro server snapshot restore](gyro_server_snapshot_restore.md)	 - Restore a snapshot onto the server (overwrites the disk)

