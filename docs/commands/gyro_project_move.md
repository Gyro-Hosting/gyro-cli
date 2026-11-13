## gyro project move

Move services into a project

### Synopsis

Move services into a project

All ids must be your services within the key's scope, else 404 `service_not_found` and nothing moves. Up to 500 per call; id 0 = default project. Sub-account keys need the `project:move` action. Writes: 60 per minute per key (429 `rate_limited`).

API: POST /cloud/projects/{id}/services (operationId moveServicesToProject, scope services_write)

```
gyro project move <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for move
      --idempotency-key string   idempotency key (automatic for money operations)
      --service-ids int64Slice   service_ids (default [])
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

* [gyro project](gyro_project.md)	 - List projects

