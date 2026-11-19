## gyro account card update

Use a card for auto-payment / make it the default

### Synopsis

Use a card for auto-payment / make it the default

`auto_pay` true adds the card to the end of the auto-payment chain (renewal invoices are paid from balance first, then from these cards in order), false removes it. `default: true` makes it the default card. Expired cards are refused (422 card_expired).

API: PATCH /cloud/account/cards/{id} (operationId updateSavedCard, scope billing_pay)

```
gyro account card update <ID> [flags]
```

### Options

```
      --auto-pay                 auto_pay
      --data string              request body as JSON, @file or @- (stdin)
      --default                  default
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for update
      --idempotency-key string   idempotency key (automatic for money operations)
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro account card](gyro_account_card.md)	 - List saved cards

