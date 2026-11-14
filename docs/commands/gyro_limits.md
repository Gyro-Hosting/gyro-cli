## gyro limits

Account limits and usage; how to raise them (support ticket)

### Synopsis

Account limits: every limit of your account with the current usage (new-account limit of the prepaid cloud,
Gyro Cloud Istanbul servers, templates, additional IPs and disks per server, API order caps, keys, webhooks,
firewall rules and security groups, storage buckets and keys, DNS zones, container plans, Gyro Pipeline,
uptime monitors, projects).

Limits are raised on request: open a support ticket.
  gyro ticket create --department-id 3 --subject "Limit increase: <key>" --message "Requested limit: <N>. Reason: ..."
Plan limits (buckets, zones, container services ...) also grow with a bigger plan or an add-on.
When a command is refused because of a limit, the error shows the limit, the usage and the ready ticket command.

```
gyro limits [flags]
```

### Examples

```
  gyro limits
  gyro limits -o json --query "limits[?key=='compute.max_vms']"
  gyro ticket create --department-id 3 --subject "Limit increase: new_account.servers" --message "Requested limit: 10. Reason: ..."
```

### Options

```
      --fresh   read again (answers are cached for 60 s)
  -h, --help    help for limits
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

