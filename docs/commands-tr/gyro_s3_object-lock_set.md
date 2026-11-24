## gyro s3 object-lock set

Kovanın varsayılan saklamasını ayarla

### Synopsis

Kovanın varsayılan saklamasını ayarla

Yalnız kilitli kovalar. `mode` GOVERNANCE ya da COMPLIANCE ve `days` 1-3650, ya da OFF (yeni nesnelere otomatik saklama yok; mevcut kilitler sürer). GOVERNANCE: yalnız Gyro personeli talep üzerine aşabilir. COMPLIANCE kilidi süresi dolana kadar kimse (Gyro dahil) kısaltamaz, GOVERNANCE'a çeviremez, kaldıramaz; veri saklanır ve iptalde bile süre sonuna kadar ücret sürer. `"confirm": true` ve `"acknowledge": "COMPLIANCE"` gönderin. Kilit bitişi hizmet bitiş (yenileme) tarihini aşamaz: bugün + `days` (ya da yeni bitiş) daha sonraysa yanıt 422 `retention_exceeds_service_end` olur (`details.service_end`, `details.max_days`). `GET /cloud/storage/{id}/object-lock` `service_end` ve `max_retention_days` gösterir; daha uzun kilit için hizmeti yenileyin ya da daha uzun döneme (`biennially`, `triennially`) geçirin.

API: PUT /cloud/storage/{id}/buckets/{bucket}/object-lock (operationId setStorageBucketObjectLock, scope services_write)

```
gyro s3 object-lock set <ID> <BUCKET> [flags]
```

### Options

```
      --acknowledge string       Yalnız COMPLIANCE: kaldırılamayacağını kabul için COMPLIANCE yazın
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --days int                 Retention days (not with OFF); at most the days left until the service end date (see max_retention_days)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --mode string              mode (GOVERNANCE|COMPLIANCE|OFF)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
```

### Options inherited from parent commands

```
      --account gyro account list   panel kullanıcısı olduğunuz başka bir hesapta işlem yap (gyro account list id'si; "self" = kendi hesabınız; ortam GYRO_ACCOUNT, profil varsayılanı: `gyro account use`)
      --api-url string              API taban adresi (ortam GYRO_API_URL)
      --debug                       HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string                 mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color                    renkleri kapat (NO_COLOR da olur)
      --no-header                   tablo başlığını yazma
  -o, --output string               çıktı biçimi: table, json, yaml, text
  -p, --profile string              kullanılacak profil (ortam GYRO_PROFILE)
      --query string                sonuca uygulanacak JMESPath sorgusu
  -q, --quiet                       yalnız kimlikleri yaz
      --timeout duration            toplam zaman aşımı, ör. 30s, 10m
      --wide                        tüm tablo sütunlarını göster
  -y, --yes                         onay sorma
```

### SEE ALSO

* [gyro s3 object-lock](gyro_s3_object-lock.md)	 - Kovanın varsayılan saklaması

