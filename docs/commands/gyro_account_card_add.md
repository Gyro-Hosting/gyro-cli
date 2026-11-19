## gyro account card add

Add a card (not possible over the API)

### Synopsis

Add a card (not possible over the API)

Always 409 `card_add_requires_3d` with `details.add_url`: cards are added in the browser with 3-D Secure, card numbers are never accepted by the API.

API: POST /cloud/account/cards (operationId addSavedCard, scope billing_pay)

```
gyro account card add [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
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

