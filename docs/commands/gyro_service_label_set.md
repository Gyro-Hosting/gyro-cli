## gyro service label set

Set the nickname and/or tags of a service

### Synopsis

Set the nickname and/or tags of a service

Send `nickname`, `tags` or both; a field that is not sent stays unchanged. `nickname` "" or null clears it; `tags` replaces all tags and `[]` clears them. A tag is `"name"` or `{"name", "color"}`; an unknown or missing colour gets a stable default. Duplicate names (case-insensitive) are merged. Limit: 30 label writes per minute per key (429 `rate_limited`). Sub-account keys need the `label:edit` action. Errors: 422 `nothing_to_update`, `invalid_nickname`, `invalid_tags`; 404 `service_not_found` (also for services of other accounts) or `feature_unavailable`.

API: PUT /cloud/services/{id}/label (operationId updateServiceLabel, scope services_write)

```
gyro service label set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --nickname string          "" or null clears the nickname. No control characters.
      --set stringArray          extra body field key=value (repeatable)
      --tags strings             Replaces all tags; [] or null clears them. Item: a tag name, or {"name", "color"} (CLI: use --data for colours).
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

* [gyro service label](gyro_service_label.md)	 - Get the nickname and tags of a service

