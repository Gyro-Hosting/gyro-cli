## gyro server storage-mount create

Mount a bucket on the server

### Synopsis

Mount a bucket on the server

Mount a bucket of your Gyro Object Storage service on a Linux server (systemd) as a folder (rclone FUSE mount with a local write cache, starts on boot, reconnects). Each mount gets its own access key limited to that one bucket (read-only or read/write); it does not use your key limit and is deleted with the mount. Gyro Cloud Istanbul (tr-ist-1) servers apply it automatically (`automatic: true`); other servers run `install_command` once. Free; S3 traffic counts towards the storage service. Supported servers: Gyro Cloud (Istanbul and EU), Turkey VDS, dedicated servers. The answer carries `install_command` (single use, 60 minutes). **Protected Storage (2.20.0):** supported; read-only unless `read_only: false` together with `accept_versioning: true`.

API: POST /cloud/servers/{id}/storage-mounts (operationId createStorageMount, scope cloud_write)

```
gyro server storage-mount create <ID> [flags]
```

### Options

```
      --accept-versioning        Protected Storage read/write only: confirms that every write/overwrite creates a new version, deletes only add a delete marker, locked files cannot be deleted and used space/cost grow. Without it: 422 versioning_ack_required.
      --bucket string            bucket name in that service
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --mount-point string       absolute path under /mnt, /srv, /data, /media, /opt, /home or /var/www; default /mnt/<bucket>. Must be empty or missing on the server.
      --read-only                Default: false on Object Storage, **true on Protected Storage**.
      --set stringArray          extra body field key=value (repeatable)
      --storage-service-id int   your active Object Storage service
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
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

