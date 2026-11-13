## gyro order create

Create order

### Synopsis

Create order

Create order. Orders of the calling account. **Spends money / creates an invoice.** Money rules apply (AK-40): `"confirm": true` in the body, an `Idempotency-Key` header, account balance only (no stored card), daily limit of paid API operations per account (default 50, operator can override per account/reseller, resets 00:00 Europe/Istanbul); violations answer 422 with `[confirm_required]`, `[idempotency_key_required]`, `[balance_only]` or `[daily_limit]` in the message. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/musteri-olarak-siparis).

API: POST /orders (operationId coreCreateOrder, scope Orders/CreateOrder)

```
gyro order create [flags]
```

### Options

```
      --addons string                  Seçilen ek hizmetler. Ek hizmet numarasını seçenek numarasına eşler. Ürünün ilk dönemi ücretsiz dahil ettiği ek hizmet gönderilmese de siparişe eklenir. İlk dönem tutarı 0 olur, yenilemeler ek hizmet fiyatından faturalanır ve siparişten çıkarılamaz. (JSON)
      --addons-qty string              Adetli ek hizmetlerin sayısı. (JSON)
      --billing-profile-id int         Fatura profilinin numarası.
      --card-id int                    Çekilecek kartın numarası. Verilmezse hesabın varsayılan kartı çekilir.
      --coupons string                 Kupon kodları. Operatörün kendiliğinden uygulanan kuponları da katılır.
      --cycle string                   Ödeme döngüsü. Ürünün fiyatlı döngülerinden biri olmalıdır.
      --data string                    request body as JSON, @file or @- (stdin)
      --domain string                  Satırın alan adı ekseni. Seçim, ad, alt alan adı kökü ve lisans alanlarını taşır. (JSON)
      --dry-run                        show the price quote / request without changing anything
  -h, --help                           help for create
      --idempotency-key string         idempotency key (automatic for money operations)
      --items string                   items[]Sipariş satırları. Birden yirmiye kadar.
      --method string                  Ödeme yolu: hesap bakiyesi ya da kayıtlı kart.
      --metrics string                 Açılacak ölçümlerin numaraları.
      --notes string                   Sipariş notu. Bin karakterden uzunu kırpılır.
      --payment string                 paymentÖdeme kaynağı. Toplam sıfırın üstündeyse zorunludur. (JSON)
      --product-id int                 Sipariş edilecek ürünün numarası.
      --project-id project_not_found   Optional project for the new service(s) (0 = default project). Checked before the order: unknown id = 404 project_not_found (GyroAPI order endpoints) or 422 `checkout_blocked` with `[project_not_found]` (POST /orders).
      --quantity int                   Adet. Yalnız ürün çoklu alıma izin veriyorsa dikkate alınır.
      --requirement-files string       Dosya isteyen alanların içeriği. Kodlanmış içerik olarak gönderilir; adres kabul edilmez. (JSON)
      --requirements string            Ürünün sorduğu alanların yanıtları. Zorunlu işaretlenen her alan gönderilmelidir. (JSON)
      --set stringArray                extra body field key=value (repeatable)
      --terms                          Şartların kabulü. Doğru gönderilmelidir.
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

* [gyro order](gyro_order.md)	 - Orders (preview and create)

