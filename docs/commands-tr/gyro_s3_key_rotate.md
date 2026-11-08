## gyro s3 key rotate

access-keys:rotate (GyroDepolama)

### Synopsis

access-keys:rotate (GyroDepolama)

Genel köprünün kısayolu: araç `access-keys`, eylem `rotate`. Gövde alanları aracın kendi alanlarıdır. `"confirm": true` gerekir.

API: POST /cloud/storage/{id}/keys/{key}/rotate (operationId typed_post_storage__id__keys__key__rotate, scope services_write)

```
gyro s3 key rotate <ID> <KEY> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for rotate
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

* [gyro s3 key](gyro_s3_key.md)	 - Erişim anahtarları (gizli anahtar bir kez gösterilir)

