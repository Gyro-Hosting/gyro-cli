## gyro webhook

Webhooks (signed event callbacks)

```
gyro webhook [flags]
```

### Options

```
  -h, --help   help for webhook
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

* [gyro](gyro.md)	 - Gyro Hosting command-line interface
* [gyro webhook create](gyro_webhook_create.md)	 - Create a webhook (secret returned once)
* [gyro webhook delete](gyro_webhook_delete.md)	 - Delete a webhook
* [gyro webhook delivery](gyro_webhook_delivery.md)	 - Last 100 deliveries
* [gyro webhook get](gyro_webhook_get.md)	 - Get a webhook
* [gyro webhook list](gyro_webhook_list.md)	 - List webhooks
* [gyro webhook test](gyro_webhook_test.md)	 - Send an api.test event now
* [gyro webhook update](gyro_webhook_update.md)	 - Update url, events, description or status

