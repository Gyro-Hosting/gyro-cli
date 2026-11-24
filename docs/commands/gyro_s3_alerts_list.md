## gyro s3 alerts list

Fill alerts of all buckets

### Synopsis

Fill alerts of all buckets

A **fill alert** sends one e-mail to the account address when a bucket's usage reaches the threshold: `percent` = a share of the package size (every bucket of the service shares the package) or `gb` = an absolute size. Checked every 5 minutes. After the e-mail the alert is `triggered`; it re-arms when usage drops below 95% of the threshold, and the next crossing sends again. The alert only informs, uploads are not stopped. The package-wide 80% / 100% e-mails stay as they are. Object Storage and Protected Storage (old versions and delete markers count as usage).

API: GET /cloud/storage/{id}/alerts (operationId listStorageBucketAlerts, scope services_read)

```
gyro s3 alerts list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

