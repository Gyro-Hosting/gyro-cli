## gyro container service revisions

Revision history

### Synopsis

Revision history

Last 15 revisions, newest first, with image, size, variable names and reason.

API: GET /cloud/containers/{id}/services/{name}/revisions (operationId typed_get_containers__id__services__name__revisions, scope services_read)

```
gyro container service revisions <ID> <NAME> [flags]
```

### Options

```
  -h, --help   help for revisions
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

