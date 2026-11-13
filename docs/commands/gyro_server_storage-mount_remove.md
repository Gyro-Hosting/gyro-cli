## gyro server storage-mount remove

Remove a storage mount

### Synopsis

Remove a storage mount

The mount key is deleted at once (the server can no longer reach the bucket; files in the bucket are kept). tr-ist-1 servers remove the unit automatically; on others run `remove_command`.

API: DELETE /cloud/servers/{id}/storage-mounts/{mid} (operationId deleteStorageMount, scope cloud_write)

```
gyro server storage-mount remove <ID> <MID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for remove
      --idempotency-key string   idempotency key (automatic for money operations)
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

