## gyro template rebuild

Reinstall a server from a template

### Synopsis

Reinstall a server from a template

Overwrites the root disk of one of your tr-ist-1 servers with the template (its disk must be at least `min_disk_gb`); the server's SSH keys are installed again and a new root password is returned (withheld when an Idempotency-Key is sent). Needs `"confirm": true`.

API: POST /cloud/templates/{id}/rebuild (operationId rebuildFromTemplate, scope cloud_write)

```
gyro template rebuild <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for rebuild
      --idempotency-key string   idempotency key (automatic for money operations)
      --server-id int            server_id
      --set stringArray          extra body field key=value (repeatable)
      --user-data string         user_data
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

* [gyro template](gyro_template.md)	 - List your templates

