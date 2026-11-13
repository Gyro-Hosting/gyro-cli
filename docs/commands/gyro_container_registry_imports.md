## gyro container registry imports

Image imports (progress and history)

### Synopsis

Image imports (progress and history)

`active` = the running import (or null) with `phase`, `bytes_done`/`bytes_total`, `layers_done`/`layers`; `jobs` = the last 20 imports (newest first) with `state` and, when failed, `error` + `message`; `limits`.

API: GET /cloud/containers/{id}/registry/imports (operationId listContainerImageImports, scope services_read)

```
gyro container registry imports <ID> [flags]
```

### Options

```
  -h, --help   help for imports
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

