## gyro server app log

Son uygulama kurulumunun canlı ilerleme günlüğü (adımlar, sunucu günlüğü)

```
gyro server app log [flags]
```

### Options

```
  -h, --help   help for log
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
* [gyro server app log get](gyro_server_app_log_get.md)	 - Son uygulama kurulumunun canlı ilerleme günlüğü (adımlar, sunucu günlüğü)

