## gyro container pipeline webhook

Webhook address and secret

### Synopsis

Webhook address and secret

Returns `webhook_url` and `webhook_secret`.

API: GET /cloud/containers/{id}/pipelines/{pipeline}/webhook (operationId pipeline_webhook, scope services_read)

```
gyro container pipeline webhook <ID> <PIPELINE> [flags]
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

* [gyro container pipeline](gyro_container_pipeline.md)	 - List pipelines

