## gyro project rename

Rename / recolour a project

### Synopsis

Rename / recolour a project

Send any of name, description, color, icon. The default project cannot be changed (409 `default_project`). Writes: 60 per minute per key (429 `rate_limited`). Sub-account keys cannot create, rename or delete projects (403 `subkey_forbidden`).

API: PATCH /cloud/projects/{id} (operationId updateProject, scope services_write)

```
gyro project rename <ID> [flags]
```

### Options

```
      --color string             color (gray|blue|green|red|orange|yellow|purple|pink|teal|indigo)
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for rename
      --icon string              icon (folder|briefcase|building|cart3|cloud|code-slash|controller|globe2|heart|house|kanban|lightning|people|rocket-takeoff|shield-lock|shop|star|tools|truck|box-seam|mortarboard|camera)
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              Letter or digit first; then letters, digits, space and . _ - / + # & ( ). Unique per account (case-insensitive). Reserved: Genel, General, Default.
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

