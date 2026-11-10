## gyro server app

Uygulama kurulum durumu (hazır olunca kimlik bir kez)

```
gyro server app [flags]
```

### Options

```
  -h, --help   help for app
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

* [gyro server](gyro_server.md)	 - Bulut sunucular: oluşturma, güç, snapshot, SSH ...
* [gyro server app create](gyro_server_app_create.md)	 - Uygulama kur (diski yeniden kurar; confirm zorunlu; sunucu başına günde 3)
* [gyro server app get](gyro_server_app_get.md)	 - Uygulama kurulum durumu (hazır olunca kimlik bir kez)
* [gyro server app log](gyro_server_app_log.md)	 - Son uygulama kurulumunun canlı ilerleme günlüğü (adımlar, sunucu günlüğü)
* [gyro server app retry](gyro_server_app_retry.md)	 - Hata veren ya da yanıt vermeyen kurulumu yeniden dene (diski yeniden kurar; confirm zorunlu; sunucu başına saatte 3)

