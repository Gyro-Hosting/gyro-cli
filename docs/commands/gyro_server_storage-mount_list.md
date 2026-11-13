## gyro server storage-mount list

List object storage mounts of a server

### Synopsis

List object storage mounts of a server

Mount a bucket of your Gyro Object Storage service on a Linux server (systemd) as a folder (rclone FUSE mount with a local write cache, starts on boot, reconnects). Each mount gets its own access key limited to that one bucket (read-only or read/write); it does not use your key limit and is deleted with the mount. Gyro Cloud Istanbul (tr-ist-1) servers apply it automatically (`automatic: true`); other servers run `install_command` once. Free; S3 traffic counts towards the storage service. Protected Storage buckets: read-only by default (2.20.0). Supported servers: Gyro Cloud (Istanbul and EU), Turkey VDS, dedicated servers.

API: GET /cloud/servers/{id}/storage-mounts (operationId listStorageMounts, scope cloud_read)

```
gyro server storage-mount list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro server storage-mount](gyro_server_storage-mount.md)	 - List object storage mounts of a server

