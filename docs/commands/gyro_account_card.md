## gyro account card

List saved cards

```
gyro account card [flags]
```

### Options

```
  -h, --help   help for card
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

* [gyro account](gyro_account.md)	 - Automatic balance top-up setting
* [gyro account card add](gyro_account_card_add.md)	 - Add a card (not possible over the API)
* [gyro account card delete](gyro_account_card_delete.md)	 - Delete a saved card (also at the bank)
* [gyro account card list](gyro_account_card_list.md)	 - List saved cards
* [gyro account card update](gyro_account_card_update.md)	 - Use a card for auto-payment / make it the default

