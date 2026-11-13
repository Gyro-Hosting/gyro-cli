## gyro project create

Create a project

### Synopsis

Create a project

At most 50 projects per account (409 `project_limit`). Optional `service_ids` move services in at once. Writes: 60 per minute per key (429 `rate_limited`). Sub-account keys cannot create, rename or delete projects (403 `subkey_forbidden`). Errors: 422 `invalid_name`, `name_reserved`, `invalid_color`, `invalid_icon`, `invalid_description`, `invalid_service_ids`; 409 `name_taken`; 404 `service_not_found`.

API: POST /cloud/projects (operationId createProject, scope services_write)

```
gyro project create [flags]
```

### Options

```
      --color string             color (gray|blue|green|red|orange|yellow|purple|pink|teal|indigo)
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --icon string              icon (folder|briefcase|building|cart3|cloud|code-slash|controller|globe2|heart|house|kanban|lightning|people|rocket-takeoff|shield-lock|shop|star|tools|truck|box-seam|mortarboard|camera)
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              Letter or digit first; then letters, digits, space and . _ - / + # & ( ). Unique per account (case-insensitive). Reserved: Genel, General, Default.
      --service-ids int64Slice   Optional: move these services into the new project at once (default [])
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

