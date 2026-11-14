## gyro container notifications set

Set notification channels

### Synopsis

Set notification channels

Send only the fields to change. A webhook secret is generated when a webhook URL is first set (POST .../notifications/webhook-secret returns it). No WhatsApp.

API: PUT /cloud/containers/{id}/notifications (operationId typed_put_containers__id__notifications, scope services_write)

```
gyro container notifications set <ID> [flags]
```

### Options

```
      --cooldown int             Minutes between two announcements of the same service + event (5-1440, default 30)
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --email                    Account e-mail (default on)
      --events string            Comma separated: down, crashloop, replicas_low, scale_max, capacity, egress
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --resolved                 Send a 'resolved' message (default on)
      --set stringArray          extra body field key=value (repeatable)
      --sms                      Account phone
      --telegram string          Telegram chat id (send /start to the platform bot, e.g. @gyrouptime_bot); empty = off
      --webhook-url string       https URL of a public server; JSON POST signed with X-Gyro-Signature: sha256=HMAC(secret, X-Gyro-Timestamp + '.' + body)
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

* [gyro container notifications](gyro_container_notifications.md)	 - Notification settings and current alerts

