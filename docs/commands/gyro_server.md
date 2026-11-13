## gyro server

Cloud servers: create, power, snapshots, SSH ...

```
gyro server [flags]
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro server activity](gyro_server_activity.md)	 - Recent platform actions on the server
* [gyro server app](gyro_server_app.md)	 - App install status (credentials shown once when ready)
* [gyro server backup](gyro_server_backup.md)	 - List automatic backups
* [gyro server console](gyro_server_console.md)	 - Open a one-time VNC console ticket
* [gyro server create](gyro_server_create.md)	 - Order a new cloud server (paid from your balance)
* [gyro server delete](gyro_server_delete.md)	 - Delete (cancel) a server - its data is destroyed after the grace period
* [gyro server disk](gyro_server_disk.md)	 - List the server's additional disks
* [gyro server firewall](gyro_server_firewall.md)	 - Inbound firewall rules
* [gyro server get](gyro_server_get.md)	 - Show a server with its live state
* [gyro server ip](gyro_server_ip.md)	 - Print the public IPv4 (or --ipv6) of a server
* [gyro server list](gyro_server_list.md)	 - List your servers
* [gyro server metrics](gyro_server_metrics.md)	 - CPU, disk and network metrics (summary)
* [gyro server power](gyro_server_power.md)	 - Power on, off, shutdown, reboot or reset a server
* [gyro server rdns](gyro_server_rdns.md)	 - List reverse DNS entries
* [gyro server rebuild](gyro_server_rebuild.md)	 - Reinstall the OS - WIPES THE DISK
* [gyro server rename](gyro_server_rename.md)	 - Change the host name
* [gyro server rescue](gyro_server_rescue.md)	 - Rescue system (boot into a recovery Linux)
* [gyro server reset-password](gyro_server_reset-password.md)	 - Set a new random root password (shown once)
* [gyro server snapshot](gyro_server_snapshot.md)	 - List snapshots and the snapshot quota
* [gyro server ssh](gyro_server_ssh.md)	 - Open an SSH session to a server
* [gyro server storage-mount](gyro_server_storage-mount.md)	 - List object storage mounts of a server
* [gyro server update](gyro_server_update.md)	 - Change host name or auto-renew
* [gyro server wait](gyro_server_wait.md)	 - Wait until a server is provisioned and idle

