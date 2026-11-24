## gyro container schedule preview

Zamanlamanın sonraki çalışma saatleri

### Synopsis

Zamanlamanın sonraki çalışma saatleri

`cron` / `rate_minutes` / `at` + `timezone`'u denetler, o saat diliminde sonraki `count` (1-20, varsayılan 5) çalışma saatini döner. Hiçbir şey kaydedilmez.

API: POST /cloud/containers/{id}/schedules/preview (operationId previewContainerSchedule, scope services_write)

```
gyro container schedule preview <ID> [flags]
```

### Options

```
      --at timezone              RFC 3339, or local date-time in timezone
      --count int                count
      --cron string              cron
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for preview
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --rate-minutes int         rate_minutes
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --timezone string          timezone
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

* [gyro container schedule](gyro_container_schedule.md)	 - Zamanlanmış görevleri sonraki ve son çalışmalarıyla listele

