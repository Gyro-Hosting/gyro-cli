## gyro s3 object retention set

Nesne sürümünü kilitle

### Synopsis

Nesne sürümünü kilitle

Bir sürümün saklamasını şimdiden itibaren `days` (1-3650) gün olarak koyar ya da uzatır (`version_id` boş = güncel). Saklama yalnız uzatılabilir; COMPLIANCE GOVERNANCE'a çevrilemez. COMPLIANCE kilidi süresi dolana kadar kimse (Gyro dahil) kısaltamaz, GOVERNANCE'a çeviremez, kaldıramaz; veri saklanır ve iptalde bile süre sonuna kadar ücret sürer. `"confirm": true` ve `"acknowledge": "COMPLIANCE"` gönderin.

API: PUT /cloud/storage/{id}/buckets/{bucket}/objects/retention (operationId setStorageObjectRetention, scope services_write)

```
gyro s3 object retention set <ID> <BUCKET> [flags]
```

### Options

```
      --acknowledge string       Yalnız COMPLIANCE: COMPLIANCE yazın
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --days int                 days
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --key string               key
      --mode string              mode (GOVERNANCE|COMPLIANCE)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --version-id string        Empty = current version
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

* [gyro s3 object retention](gyro_s3_object_retention.md)	 - Nesne sürümünün saklaması

