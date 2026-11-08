## gyro affiliate withdraw

Withdraw affiliate

### Synopsis

Withdraw affiliate

Withdraw affiliate. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/satis-ortakligi-kazanci).

API: POST /affiliate/withdraw (operationId coreWithdrawAffiliate, scope Affiliate/WithdrawAffiliate)

```
gyro affiliate withdraw [flags]
```

### Options

```
      --amount float             İstenen tutar. Alt sınır ile kullanılabilir bakiye arasında olmalıdır.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --gateway string           Ödeme yolu. Operatörün sunduklarından biri olmalıdır.
      --gateway-info string      Hedef bilgisi. Kayıtlı bir hedef varsa atlanabilir.
  -h, --help                     help for withdraw
      --idempotency-key string   idempotency key (automatic for money operations)
      --save-default             Bu hedefi sonraki talepler için kaydet.
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

* [gyro affiliate](gyro_affiliate.md)	 - Get affiliate

