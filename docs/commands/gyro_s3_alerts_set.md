## gyro s3 alerts set

Set the fill alert of a bucket

### Synopsis

Set the fill alert of a bucket

Sets or replaces the threshold (the alert re-arms). If the bucket is already above it, the e-mail is sent within ~5 minutes. `type` percent (1-100, of the package size) or gb (1 to the package size). Errors: 422 `tool_action_failed` with a message in the account language. A **fill alert** sends one e-mail to the account address when a bucket's usage reaches the threshold: `percent` = a share of the package size (every bucket of the service shares the package) or `gb` = an absolute size. Checked every 5 minutes. After the e-mail the alert is `triggered`; it re-arms when usage drops below 95% of the threshold, and the next crossing sends again. The alert only informs, uploads are not stopped. The package-wide 80% / 100% e-mails stay as they are. Object Storage and Protected Storage (old versions and delete markers count as usage).

API: PUT /cloud/storage/{id}/buckets/{bucket}/alert (operationId setStorageBucketAlert, scope services_write)

```
gyro s3 alerts set <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
      --type string              type (percent|gb)
      --value int                value
```

### Options inherited from parent commands

```
      --account gyro account list   act in another account you are a panel user of (id from gyro account list; "self" = your own; env GYRO_ACCOUNT, profile default: `gyro account use`)
      --api-url string              API base URL (env GYRO_API_URL)
      --debug                       log HTTP requests to stderr (secrets masked)
      --lang string                 message language: tr or en (default from locale)
      --no-color                    disable colours (also NO_COLOR)
      --no-header                   omit the table header
  -o, --output string               output format: table, json, yaml, text
  -p, --profile string              profile to use (env GYRO_PROFILE)
      --query string                JMESPath query applied to the result
  -q, --quiet                       print only ids
      --timeout duration            overall timeout, e.g. 30s, 10m
      --wide                        show all table columns
  -y, --yes                         do not ask for confirmation
```

### SEE ALSO

* [gyro s3 alerts](gyro_s3_alerts.md)	 - Fill alerts of all buckets

