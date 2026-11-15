## gyro s3 events add

Add an event webhook to a bucket

### Synopsis

Add an event webhook to a bucket

**Bucket events** send a signed HTTPS POST to your address when an object is created (upload, copy, multipart upload), deleted, or a bulk delete runs. Body: `{"id", "event", "created_at", "data": {service_id, bucket, key, size, etag, content_type, version_id (Protected Storage), method (put / multipart / delete / bulk), source (api / panel), access_key_id, event_time}}` - no object content. Headers: `X-Gyro-Event`, `X-Gyro-Delivery` (unique id, use it to drop duplicates), `X-Gyro-Timestamp`, `X-Gyro-Signature: sha256=HMAC-SHA256(secret, timestamp + "." + raw body)` (same format as account webhooks). Delivery is at least once with no ordering guarantee; a non-2xx answer or no answer within 6 s is retried (8 attempts, about 2 hours); 20 failures in a row disable the webhook (`disabled_reason` = `auto_failures`). Requests come from 45.87.174.151, 45.87.174.153, 45.87.174.154. Targets: https only, port 443 or 8443, public internet (private, loopback and Gyro platform addresses are refused, the name is resolved again on every delivery). Filters: `prefix` (key starts with) and `suffix` (key ends with one of, at most 10). Limits: 5 webhooks per bucket, 20 per service (`s3.max_event_hooks`, raised with a support ticket). Price: a monthly free allowance per bucket (see `GET /cloud/storage/{id}/events/usage`) plus prepaid event packs valid in the month of purchase; no overage invoice: when the allowance is used up new events are not sent. Known gaps: a bulk delete does not list the keys; lifecycle deletions produce no event. The answer carries `secret` once. A refusal is 422 `tool_action_failed` with a message in the account language (invalid / private / unresolvable URL, port, event type, limits).

API: POST /cloud/storage/{id}/buckets/{bucket}/events (operationId createBucketEventWebhook, scope services_write)

```
gyro s3 events add <ID> <BUCKET> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --description string       description
      --dry-run                  show the price quote / request without changing anything
      --events strings           default: all three. Short names created / deleted / deleted_bulk are accepted
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --prefix string            prefix
      --set stringArray          extra body field key=value (repeatable)
      --suffix strings           a comma separated text is accepted too
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

* [gyro s3 events](gyro_s3_events.md)	 - List the event webhooks of a bucket

