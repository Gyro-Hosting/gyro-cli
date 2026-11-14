## gyro container notifications get

Notification settings and current alerts

### Synopsis

Notification settings and current alerts

`settings` (channels, events, cooldown; the webhook secret is never returned), `conditions` = what is wrong right now, `history` = last 30 alerts. Events: down (no ready replica for >3 min), crashloop (keeps restarting / image pull fails / out of memory), replicas_low (below alerts.min_ready), scale_max (autoscaling wants more than max), capacity (replicas wait for room in the node pool), egress (the egress IPs of the environment changed; one message, no resolved). An event is announced after ~2 minutes, at most once per cooldown per service + event (flapping included), with a 'resolved' message.

API: GET /cloud/containers/{id}/notifications (operationId typed_get_containers__id__notifications, scope services_read)

```
gyro container notifications get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

