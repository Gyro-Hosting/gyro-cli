## gyro s3 bucket rename

Kova adını değiştir

### Synopsis

Kova adını değiştir

Kova adları hizmete özeldir (aynı adı başka müşteriler de kullanabilir). Veri taşınmaz. Eski adla S3 erişimi hemen durur; web sitesi `<yeni>-<hizmet kodu>.web.gyro-s3.com` olur, eski web adresi 40 gün 301 ile yönlendirilir; paylaşım bağlantıları çalışır. Kova başına günde en çok 5; eski ad 40 gün bu hizmette yeniden kullanılamaz. `"confirm": true` gerekir.

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

