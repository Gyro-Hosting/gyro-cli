## gyro container service list

List container services

### Synopsis

List container services

Services of the environment with state (running, deploying, failing, pending, stopped, suspended), ready/desired instances, size and public URL. Each item carries `internal_address` (services with a port); the list carries `namespace` and `cluster_dns`.

API: GET /cloud/containers/{id}/services (operationId typed_get_containers__id__services, scope services_read)

```
gyro container service list <ID> [flags]
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

* [gyro container service](gyro_container_service.md)	 - List container services

