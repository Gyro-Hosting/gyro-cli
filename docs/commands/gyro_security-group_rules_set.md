## gyro security-group rules set

Replace the rules of a security group

### Synopsis

Replace the rules of a security group

Full replace (inbound + outbound), max 60 rules. Applied to every attached server. If the change removes or narrows SSH (22) / RDP (3389) access for any source on an attached server, or cuts the calling address, it runs as a 5 minute trial (`trial: true`, `trial_seconds_left`) - the same rule as the panel, CLI and Terraform: confirm with POST .../confirm, otherwise the previous rules return. A server's rules (own + all attached groups) must fit its allowance (422 err_sg_quota).

API: PUT /cloud/security-groups/{sg}/rules (operationId putSecurityGroupRules, scope cloud_write)

```
gyro security-group rules set <SG> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro security-group rules](gyro_security-group_rules.md)	 - Replace the rules of a security group

