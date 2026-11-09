## gyro s3 bucket rename

Kova adını değiştir

### Synopsis

Kova adını değiştir

Veri taşınmaz; eski ad hemen çalışmayı durdurur (S3, web sitesi, paylaşım bağlantıları). Kova başına günde en çok 5; bırakılan ad 40 gün rezerve kalır. `"confirm": true` gerekir.

API: POST /cloud/storage/{id}/buckets/{bucket}/rename (operationId typed_post_storage__id__buckets__bucket__rename, scope services_write)

```
gyro s3 bucket rename <ID> <BUCKET> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for rename
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --new-name string          new_name
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro s3 bucket](gyro_s3_bucket.md)	 - Kovalar

