## gyro container domain add

Attach a domain

### Synopsis

Attach a domain

Point a CNAME to the service host (`<name>-<code>.gyroapp.net`) first; the domain is activated (with a free certificate) once its DNS points to us, checked hourly or with .../domains/verify.

API: POST /cloud/containers/{id}/services/{name}/domains (operationId typed_post_containers__id__services__name__domains, scope services_write)

```
gyro container domain add <ID> <NAME> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --domain string            domain
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
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

