## gyro container registry get

Gyro image registry of the environment

### Synopsis

Gyro image registry of the environment

Host, image prefix `<host>/<code>/`, username, usage (`used_gb`, `free_gb` = 2, `billable_gb`, `cap_gb` = 50), `push_allowed` and the clean-up policy. Only this environment can push/pull its repositories. Above 2 GB: €0.05 per GB per month, charged hourly from the prepaid balance. The password is returned only by POST .../registry/credentials.

API: GET /cloud/containers/{id}/registry (operationId typed_get_containers__id__registry, scope services_read)

```
gyro container registry get <ID> [flags]
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

* [gyro container registry](gyro_container_registry.md)	 - Gyro image registry of the environment

