## gyro server app log get

Son uygulama kurulumunun canlı ilerleme günlüğü (adımlar, sunucu günlüğü)

### Synopsis

Son uygulama kurulumunun canlı ilerleme günlüğü (adımlar, sunucu günlüğü)

`after` değerinden büyük kimlikli satırlar, eskiden yeniye (elinizdeki son kimlikle sorgulayın). `source`: step, server, system, staff. Sunucu satırları sunucudan gelen güvenilmez metindir (sırlar maskelenir); düz metin olarak gösterin.

API: GET /cloud/servers/{id}/app/log (operationId getAppInstallLog, scope cloud_read)

```
gyro server app log get <ID> [flags]
```

### Options

```
      --after int         
  -h, --help              help for get
      --param-limit int   
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

* [gyro server app log](gyro_server_app_log.md)	 - Son uygulama kurulumunun canlı ilerleme günlüğü (adımlar, sunucu günlüğü)

