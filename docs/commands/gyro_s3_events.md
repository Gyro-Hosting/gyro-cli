## gyro s3 events

List the event webhooks of a bucket

```
gyro s3 events [flags]
```

### Options

```
  -h, --help   help for events
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

* [gyro s3](gyro_s3.md)	 - S3 object storage: buckets and access keys
* [gyro s3 events add](gyro_s3_events_add.md)	 - Add an event webhook to a bucket
* [gyro s3 events delete](gyro_s3_events_delete.md)	 - Delete an event webhook
* [gyro s3 events deliveries](gyro_s3_events_deliveries.md)	 - Recent deliveries of a webhook
* [gyro s3 events get](gyro_s3_events_get.md)	 - Get an event webhook
* [gyro s3 events list](gyro_s3_events_list.md)	 - List the event webhooks of a bucket
* [gyro s3 events list-all](gyro_s3_events_list-all.md)	 - List all event webhooks of the service
* [gyro s3 events rotate-secret](gyro_s3_events_rotate-secret.md)	 - Rotate the signing secret
* [gyro s3 events test](gyro_s3_events_test.md)	 - Send a test event
* [gyro s3 events update](gyro_s3_events_update.md)	 - Change an event webhook
* [gyro s3 events usage](gyro_s3_events_usage.md)	 - Event allowance of this month

