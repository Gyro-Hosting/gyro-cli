## gyro account card list

List saved cards

### Synopsis

List saved cards

Masked cards only (brand, last 4, expiry). The card is kept by the bank (VakıfBank PanCode); Gyro stores no card number. `add_url` opens the client area page where a card is added with 3-D Secure.

API: GET /cloud/account/cards (operationId listSavedCards, scope billing_read)

```
gyro account card list [flags]
```

### Options

```
  -h, --help   help for list
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

