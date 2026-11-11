## gyro container nodes

Node pool

### Synopsis

Node pool

Your dedicated nodes (aliases node-1..n), pool capacity and usage, ingress IPs for apex A records, private registry (user only).

API: GET /cloud/containers/{id}/nodes (operationId typed_get_containers__id__nodes, scope services_read)

```
gyro container nodes <ID> [flags]
```

### Options

```
  -h, --help   help for nodes
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

* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains

