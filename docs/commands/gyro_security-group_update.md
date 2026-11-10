## gyro security-group update

Rename / describe a security group

### Synopsis

Rename / describe a security group

Only the display name and description change; the group id and its rules stay.

API: PATCH /cloud/security-groups/{sg} (operationId updateSecurityGroup, scope cloud_write)

```
gyro security-group update <SG> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for update
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name
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

* [gyro security-group](gyro_security-group.md)	 - List security groups

