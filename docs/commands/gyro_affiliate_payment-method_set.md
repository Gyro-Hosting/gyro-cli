## gyro affiliate payment-method set

Save affiliate payout method

### Synopsis

Save affiliate payout method

Save affiliate payout method. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/satis-ortakligi-kazanci).

API: PUT /affiliate/payment-method (operationId coreSaveAffiliatePayoutMethod, scope Affiliate/SaveAffiliatePayoutMethod)

```
gyro affiliate payment-method set [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --gateway string           Ödeme yolu.
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --info string              Hedef bilgisi. Boş göndermek kaydı siler.
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

* [gyro affiliate payment-method](gyro_affiliate_payment-method.md)	 - Save affiliate payout method

