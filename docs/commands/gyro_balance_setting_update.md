## gyro balance setting update

Update balance settings

### Synopsis

Update balance settings

Update balance settings. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/hesap-bakiyesi-api).

API: PATCH /balance/settings (operationId coreUpdateBalanceSettings, scope Balance/UpdateBalanceSettings)

```
gyro balance setting update [flags]
```

### Options

```
      --alert string              Uyarı bloğu. Açık ve eşik alanlarını taşır; açarken pozitif bir eşik gerekir. (JSON)
      --alert-recipients string   Uyarıyı alacak fatura adresi numaraları. Liste yerine geçer; hesabın sahibi burada yer almaz, her zaman alır.
      --auto-pay-from-balance     Yenilemeler cüzdandan tahsil edilsin mi. Bu değiştiğinde açık yenileme faturaları yeni tercihe göre yeniden ele alınır.
      --backup-card-id int        Cüzdan yetmediğinde çekilecek kart. Sıfır göndermek kart zincirini tümüyle kaldırır.
      --data string               request body as JSON, @file or @- (stdin)
      --dry-run                   show the price quote / request without changing anything
  -h, --help                      help for update
      --idempotency-key string    idempotency key (automatic for money operations)
      --set stringArray           extra body field key=value (repeatable)
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

* [gyro balance setting](gyro_balance_setting.md)	 - Update balance settings

