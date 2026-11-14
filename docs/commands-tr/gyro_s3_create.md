## gyro s3 create

Nesne depolama (S3) sipariş et

### Synopsis

Nesne depolama (S3) sipariş et

Bir paketten 1..N depolama hizmetini TEK siparişte açar (`quantity`, varsayılan 1, en çok 25) ve hesap bakiyesinden öder (kayıtlı kart asla). `cycle` dönemine göre ücretlendirilir (monthly | quarterly | semiannually | annually; sabit EUR fiyat hesap para biriminde) ve diğer hizmetler gibi bakiyeden yenilenir. `extra_traffic` (`1tb` | `5tb`) bu ay için tek seferlik trafik paketi ekler; `object_lock_buckets` (korumalı paketler, 1-50) nesne kilidi hakkı alır (her dönem ücretlenir). `name` (`{n}` ile), `tags` ve `project_id` hizmetleri adlandırır. Teklif için `"dry_run": true`; gerçek sipariş `"confirm": true` ve `Idempotency-Key` başlığı ister. İlk kova ve ilk erişim anahtarı yaklaşık bir dakikada kendiliğinden oluşur (durum `inprocess` -> `active`); sonra `/cloud/storage/{id}/buckets`, `/keys`, `/connect`.

API: POST /cloud/storage (operationId createStorageServices, scope cloud_order)

```
gyro s3 create [flags]
```

### Options

```
      --cycle string              fatura dönemi (monthly|quarterly|semiannually|annually)
      --data string               istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                   hiçbir şeyi değiştirmeden teklifi / isteği göster
      --end-customer-id int       reseller main key: assign to this end-customer
      --extra-traffic string      bu takvim ayı için tek seferlik trafik paketi (1tb|5tb)
  -h, --help                      help for create
      --idempotency-key string    tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string               ad (takma ad); toplu siparişte {n} = sıra no
      --object-lock-buckets int   korumalı paketler: nesne kilitli kova hakkı (her dönem)
      --plan string               GET /cloud/storage/plans paketi
      --product-id int            plan yerine ürün kimliği
      --project-id int            proje (0 = Genel)
      --quantity int              tek siparişte açılacak hizmet (--count), en çok 25
      --set stringArray           ek gövde alanı anahtar=değer (tekrarlanabilir)
      --tags strings              tags
```

### Options inherited from parent commands

```
      --api-url string     API taban adresi (ortam GYRO_API_URL)
      --debug              HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string        mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color           renkleri kapat (NO_COLOR da olur)
      --no-header          tablo başlığını yazma
  -o, --output string      çıktı biçimi: table, json, yaml, text
  -p, --profile string     kullanılacak profil (ortam GYRO_PROFILE)
      --query string       sonuca uygulanacak JMESPath sorgusu
  -q, --quiet              yalnız kimlikleri yaz
      --timeout duration   toplam zaman aşımı, ör. 30s, 10m
      --wide               tüm tablo sütunlarını göster
  -y, --yes                onay sorma
```

### SEE ALSO

* [gyro s3](gyro_s3.md)	 - S3 nesne depolama: kovalar ve erişim anahtarları

