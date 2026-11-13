## gyro server app create

Uygulama kur (diski yeniden kurar; confirm zorunlu; sunucu başına günde 3)

### Synopsis

Uygulama kur (diski yeniden kurar; confirm zorunlu; sunucu başına günde 3)

**2.19.1**: tr-ist-1 (Gyro Bulut İstanbul) sunucularında da. `POST /cloud/servers/{id}/actions/rebuild` + `app` ile aynıdır.

API: POST /cloud/servers/{id}/app (operationId installApp, scope cloud_write)

```
gyro server app create <ID> [flags]
```

### Options

```
      --app string               app
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --domain string            domain
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --email string             email
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --os string                os
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

