## gyro security-group create

Create a security group

### Synopsis

Create a security group

Name unique per account (2-64 characters). Without `rules` the group has one outbound rule (all traffic) and no inbound rule.

API: POST /cloud/security-groups (operationId createSecurityGroup, scope cloud_write)

```
gyro security-group create [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name
      --rules strings            rules
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

