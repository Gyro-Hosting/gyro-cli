## gyro order preview

Preview order

### Synopsis

Preview order

Preview order. Orders of the calling account. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/musteri-olarak-siparis).

API: POST /orders/preview (operationId corePreviewOrder, scope Orders/PreviewOrder)

```
gyro order preview [flags]
```

### Options

```
      --addons string              Seçilen ek hizmetler. Ek hizmet numarasını seçenek numarasına eşler. Ürünün ilk dönemi ücretsiz dahil ettiği ek hizmet gönderilmese de siparişe eklenir. İlk dönem tutarı 0 olur, yenilemeler ek hizmet fiyatından faturalanır ve siparişten çıkarılamaz. (JSON)
      --addons-qty string          Adetli ek hizmetlerin sayısı. (JSON)
      --billing-profile-id int     Fatura profilinin numarası. Vergi bağlamını belirler; verilmezse varsayılan profil, hesapta kayıtlı profil yoksa hesabın ülkesi kullanılır.
      --card-id string             card_id
      --coupons string             Denenecek kupon kodları.
      --cycle string               Ödeme döngüsü. Ürünün fiyatlı döngülerinden biri olmalıdır.
      --data string                request body as JSON, @file or @- (stdin)
      --domain string              Satırın alan adı ekseni. Seçim, ad, alt alan adı kökü ve lisans alanlarını taşır. (JSON)
      --dry-run                    show the price quote / request without changing anything
  -h, --help                       help for preview
      --idempotency-key string     idempotency key (automatic for money operations)
      --items string               items[]Sipariş satırları. Birden yirmiye kadar; şekli sipariş ucuyla aynıdır.
      --method string              method
      --metrics string             Açılacak ölçümlerin numaraları.
      --payment string             Ödeme bloğu. Verilirse yöntemin komisyonu toplama eklenir; bakiye yeterliliği burada denetlenmez. (JSON)
      --product-id int             Sipariş edilecek ürünün numarası.
      --quantity int               Adet. Yalnız ürün çoklu alıma izin veriyorsa dikkate alınır.
      --requirement-files string   Dosya isteyen alanların içeriği. Kodlanmış içerik olarak gönderilir; adres kabul edilmez. (JSON)
      --requirements string        Ürünün sorduğu alanların yanıtları. Zorunlu işaretlenen her alan gönderilmelidir. (JSON)
      --set stringArray            extra body field key=value (repeatable)
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

