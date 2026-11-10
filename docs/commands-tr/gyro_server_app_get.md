## gyro server app get

Uygulama kurulum durumu (hazır olunca kimlik bir kez)

### Synopsis

Uygulama kurulum durumu (hazır olunca kimlik bir kez)

2.7.0 ile ayrıca `stuck` (açık ama uzun süredir haber yok), `stuck_reason` (`no_callback` = sunucu hiç haber vermedi), `last_update_at`, `attempt`, `can_retry`, `error_code`, `step_text`. Canlı günlük: GET /cloud/servers/{id}/app/log.

API: GET /cloud/servers/{id}/app (operationId getAppInstall, scope cloud_read)

```
gyro server app get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

