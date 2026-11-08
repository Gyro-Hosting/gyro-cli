## gyro domain transfer

Transfer domain

### Synopsis

Transfer domain

Transfer domain. Domain operations run against the registrar module of the domain. **Spends money / creates an invoice.** Money rules apply (AK-40): `"confirm": true` in the body, an `Idempotency-Key` header, account balance only (no stored card), daily limit per account; violations answer 422 with `[confirm_required]`, `[idempotency_key_required]`, `[balance_only]` or `[daily_limit]` in the message. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/alan-adi-edinme).

API: POST /domains/transfer (operationId coreTransferDomain, scope Domains/TransferDomain)

```
gyro domain transfer [flags]
```

### Options

```
      --accept-premium string    accept_premium
      --addons string            Birlikte alınacak ek hizmetler. Alan adı yönetimi, kayıt gizliliği ve yönlendirme; yalnız uzantı sunuyorsa.
      --auth-code string         Mevcut sağlayıcıdan alınan transfer kodu. Uzantı işaretliyse zorunludur ve her durumda hizmete yazılır.
      --billing-profile-id int   Fatura profilinin numarası.
      --card-id string           card_id
      --contacts string          Rol bazında kişiler. Her rol ya kayıtlı bir profil numarası ya da satır içi bir kişi taşır ve o rol için profili ezer. (JSON)
      --coupons string           Kupon kodları.
      --data string              request body as JSON, @file or @- (stdin)
      --domain string            Tam alan adı.
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for transfer
      --idempotency-key string   idempotency key (automatic for money operations)
      --method string            method
      --nameservers string       Ad sunucuları. İkiden dörde; verilmezse hesabın varsayılanları uygulanır.
      --notes string             Sipariş notu.
      --payment string           Ödeme kaynağı. Hesap bakiyesi ya da kayıtlı kart. (JSON)
      --set stringArray          extra body field key=value (repeatable)
      --terms                    Şartların kabulü. Doğru gönderilmelidir.
      --whois-profile-id int     Kayıtlı kayıt sahibi profili. Dört rolü birden doldurur.
      --years string             years
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

* [gyro domain](gyro_domain.md)	 - Domains: search, register, transfer, nameservers, DNS, WHOIS

