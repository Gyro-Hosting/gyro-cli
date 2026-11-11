## gyro server firewall list

Inbound firewall rules

### Synopsis

Inbound firewall rules

Read-only in v1. An empty list means no Gyro firewall is applied (all inbound traffic reaches the server). tr-ist-1: rules come from the network firewall add-on (security groups); without the add-on the list is empty and writes answer 409 `firewall_not_enabled`.

API: GET /cloud/servers/{id}/firewall (operationId getFirewall, scope cloud_read)

```
gyro server firewall list <ID> [flags]
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

* [gyro server firewall](gyro_server_firewall.md)	 - Inbound firewall rules

