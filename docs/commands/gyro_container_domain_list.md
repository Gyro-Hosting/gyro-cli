## gyro container domain list

Domains of a service

### Synopsis

Domains of a service

Status `active` or `pending` (waiting for DNS).

API: GET /cloud/containers/{id}/services/{name}/domains (operationId typed_get_containers__id__services__name__domains, scope services_read)

```
gyro container domain list <ID> <NAME> [flags]
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

* [gyro container domain](gyro_container_domain.md)	 - Domains of a service

