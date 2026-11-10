## gyro server app retry

Hata veren ya da yanıt vermeyen kurulumu yeniden dene (diski yeniden kurar; confirm zorunlu; sunucu başına saatte 3)

### Synopsis

Hata veren ya da yanıt vermeyen kurulumu yeniden dene (diski yeniden kurar; confirm zorunlu; sunucu başına saatte 3)

Son kurulum `hata` durumundaysa ya da `stuck` = true ise (uzun süre haber yok) kullanılabilir. Bulut sunucuda sunucu aynı uygulama, işletim sistemi ve alan adıyla yeniden kurulur; diskteki TÜM veriler silinir.

API: POST /cloud/servers/{id}/app/retry (operationId retryAppInstall, scope cloud_write)

```
gyro server app retry <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for retry
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
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

* [gyro server app](gyro_server_app.md)	 - Uygulama kurulum durumu (hazır olunca kimlik bir kez)

