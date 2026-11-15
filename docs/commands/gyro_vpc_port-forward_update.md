## gyro vpc port-forward update

Change who may use a port forward (source allow list)

### Synopsis

Only the given public IPv4 networks may use the forward (a single IP becomes /32; write the network
address, e.g. 198.51.100.0/24; at most 20). The list replaces the old one; --any opens it to everyone.

```
gyro vpc port-forward update <VPC_ID> <PORT_FORWARD_ID> --source <IP|CIDR> ... | --any [flags]
```

### Examples

```
  gyro vpc port-forward update vpc-1a2b3c4d pf-9c8b7a6d --source 203.0.113.10 --source 198.51.100.0/24
  gyro vpc port-forward update vpc-1a2b3c4d pf-9c8b7a6d --any
```

### Options

```
      --any                  allow everyone (empty list)
      --dry-run              show the request without changing anything
  -h, --help                 help for update
      --source stringArray   allowed source IPv4 or CIDR (repeatable or comma separated)
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

* [gyro vpc port-forward](gyro_vpc_port-forward.md)	 - List port forwards

