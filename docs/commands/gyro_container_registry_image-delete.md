## gyro container registry image-delete

Delete an image tag

### Synopsis

Delete an image tag

Query `repo` (without the code prefix) and `tag`. Images used by a service are refused (409 image_in_use). Needs `"confirm": true`.

API: DELETE /cloud/containers/{id}/registry/images (operationId typed_delete_containers__id__registry_images, scope services_write)

```
gyro container registry image-delete <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for image-delete
      --idempotency-key string   idempotency key (automatic for money operations)
      --repo string              Repository, e.g. web
      --set stringArray          extra body field key=value (repeatable)
      --tag string               Tag, e.g. 1.4
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

