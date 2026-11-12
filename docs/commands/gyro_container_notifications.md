## gyro container notifications

Notification settings and current alerts

```
gyro container notifications [flags]
```

### Options

```
  -h, --help   help for notifications
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

* [gyro container](gyro_container.md)	 - Gyro Containers (ECS), Istanbul: services, versions, logs, domains
* [gyro container notifications get](gyro_container_notifications_get.md)	 - Notification settings and current alerts
* [gyro container notifications set](gyro_container_notifications_set.md)	 - Set notification channels
* [gyro container notifications test](gyro_container_notifications_test.md)	 - Send a test notification
* [gyro container notifications webhook-secret](gyro_container_notifications_webhook-secret.md)	 - Webhook signing secret

