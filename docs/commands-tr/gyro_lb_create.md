## gyro lb create

Yük dengeleyici sipariş et

### Synopsis

Yük dengeleyici sipariş et

Bakiyeden sipariş + ödeme (sunucularla aynı para kuralları): `"confirm": true`, `Idempotency-Key` başlığı, teklif için `dry_run`. Yaklaşık bir dakikada hazır; sonra servis ve hedef ekleyin.

API: POST /cloud/load-balancers (operationId createLoadBalancer, scope cloud_order)

```
gyro lb create [flags]
```

### Options

```
      --cycle string                   cycle
      --data string                    istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                        hiçbir şeyi değiştirmeden teklifi / isteği göster
      --end-customer-id int            end_customer_id
  -h, --help                           help for create
      --idempotency-key string         tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --product-id int                 product_id
      --project-id project_not_found   Optional project for the new service(s) (0 = default project). Checked before the order: unknown id = 404 project_not_found (GyroAPI order endpoints) or 422 `checkout_blocked` with `[project_not_found]` (POST /orders).
      --region string                  region (eu-de-1|eu-de-2|eu-fi-1)
      --set stringArray                ek gövde alanı anahtar=değer (tekrarlanabilir)
      --size string                    size (small|medium|large)
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

* [gyro lb](gyro_lb.md)	 - Yük dengeleyicileri listele

