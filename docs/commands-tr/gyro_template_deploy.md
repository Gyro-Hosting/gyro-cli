## gyro template deploy

Şablondan yeni sunucu

### Synopsis

Şablondan yeni sunucu

Şablondan kurulan normal ön ödemeli tr-ist-1 siparişi (24 saat bakiye kuralı, hesap sınırları). `product_id` (ya da `plan`) GET /cloud/products?region=tr-ist-1'den; diski şablonun `min_disk_gb` değerinden küçük olamaz. `"confirm": true` ve `Idempotency-Key` başlığı gerekir (teklif için `dry_run`).

API: POST /cloud/templates/{id}/servers (operationId createServerFromTemplate, scope cloud_order)

```
gyro template deploy <ID> [flags]
```

### Options

```
      --data string                    istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                        hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                           help for deploy
      --hostname string                hostname
      --idempotency-key string         tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --plan string                    plan
      --product-id int                 product_id
      --project-id project_not_found   Optional project for the new service(s) (0 = default project). Checked before the order: unknown id = 404 project_not_found (GyroAPI order endpoints) or 422 `checkout_blocked` with `[project_not_found]` (POST /orders).
      --set stringArray                ek gövde alanı anahtar=değer (tekrarlanabilir)
      --ssh-keys int64Slice            ssh_keys (default [])
      --user-data string               user_data
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

* [gyro template](gyro_template.md)	 - Şablonlarınızı listele

