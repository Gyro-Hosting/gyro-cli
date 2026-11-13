## gyro service project set

Set the project of a service

### Synopsis

Set the project of a service

`project_id` 0 = default project. Answers the service (with `project`). Sub-account keys need the `project:move` action.

API: PUT /cloud/services/{id}/project (operationId setServiceProject, scope services_write)

```
gyro service project set <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --project-id int           project_id
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

* [gyro service project](gyro_service_project.md)	 - Set the project of a service

