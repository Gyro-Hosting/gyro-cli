## gyro webhook create

Create a webhook (secret returned once)

### Synopsis

Create a webhook (secret returned once)

Deliveries are `POST` JSON with headers `X-Gyro-Event`, `X-Gyro-Delivery`, `X-Gyro-Timestamp` and `X-Gyro-Signature: sha256=HMAC_SHA256(secret, timestamp + "." + body)` (white-label resellers: `X-Webhook-*`). Reject requests whose timestamp is older than 5 minutes. Retries: exponential back-off, 8 attempts; 20 consecutive failures disable the webhook.

API: POST /cloud/webhooks (operationId createWebhook, scope webhooks)

```
gyro webhook create [flags]
```

### Options

```
      --customer-id int          Reseller main key only
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --dry-run                  show the price quote / request without changing anything
      --events strings           empty = all
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
      --url string               url
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

* [gyro webhook](gyro_webhook.md)	 - Webhooks (signed event callbacks)

