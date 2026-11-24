## gyro s3 events update

Olay webhook'unu değiştir

### Synopsis

Olay webhook'unu değiştir

Yalnız değişecek alanları gönderin. `status` active / disabled (etkinleştirmek hata sayacını sıfırlar). Kova değişmez.

API: PATCH /cloud/storage/{id}/events/{event} (operationId updateStorageEventWebhook, scope services_write)

```
gyro s3 events update <ID> <EVENT> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --description string       description
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --events strings           default: all three. Short names created / deleted / deleted_bulk are accepted
  -h, --help                     help for update
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --prefix string            prefix
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --status string            status (active|disabled)
      --suffix strings           a comma separated text is accepted too
      --url string               url
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

* [gyro s3 events](gyro_s3_events.md)	 - Kovanın olay webhook'larını listele

