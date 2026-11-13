## gyro project delete

Delete a project

### Synopsis

Delete a project

Only the project is deleted: its services go back to the default project, nothing else changes. Writes: 60 per minute per key (429 `rate_limited`). Sub-account keys cannot create, rename or delete projects (403 `subkey_forbidden`).

API: DELETE /cloud/projects/{id} (operationId deleteProject, scope services_write)

```
gyro project delete <ID> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
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

