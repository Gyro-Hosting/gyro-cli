## gyro template create

Capture a server as a template

### Synopsis

Capture a server as a template

Billed 0.05 EUR per GB-month of the space the template really uses, hourly from the prepaid balance (through the account's template storage service, opened with the first template). Deleting a template stops its charge. At most 10 templates; visible only in your account. Only the root disk is copied. `mode: snapshot` (default) keeps the server running (file system frozen briefly via a temporary snapshot); `mode: shutdown` stops the server, copies it and starts it again. Servers created from the template get a new host name, IP, root password, SSH host keys and machine id (cloud-init); your users, files and settings stay. Linux servers only. Ready in about a minute (`status` creating -> ready).

API: POST /cloud/servers/{id}/templates (operationId createTemplate, scope cloud_write)

```
gyro template create <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --mode string              mode (snapshot|shutdown)
      --name string              name
      --set stringArray          extra body field key=value (repeatable)
      --wait                     wait until the server is no longer busy
      --wait-timeout duration    maximum time for --wait (default 15m0s)
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

