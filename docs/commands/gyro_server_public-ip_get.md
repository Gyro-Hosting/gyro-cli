## gyro server public-ip get

Public IPv4 state

### Synopsis

Public IPv4 state

Turning the public IPv4 off makes the server reachable only from its VPC, through a gateway port forward or with the web console (which always works). Confirm with the server's current IP (`confirm_ip`). Needs a ready private card (409 `no_private_network`) and no additional IPv4 (409 `extra_ips`). If the VPC gateway (NAT) is off the server loses internet access: refused with 409 `no_egress` unless `allow_no_egress: true`. The address stays reserved for the server; turning it on brings the same IP.

API: GET /cloud/servers/{id}/public-ip (operationId getServerPublicIp, scope cloud_read)

```
gyro server public-ip get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro server public-ip](gyro_server_public-ip.md)	 - Public IPv4 state

