## gyro monitor check

İzleme kontrolleri

```
gyro monitor check [flags]
```

### Options

```
  -h, --help   help for check
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

* [gyro monitor](gyro_monitor.md)	 - Erişilebilirlik izleme
* [gyro monitor check create](gyro_monitor_check_create.md)	 - create (alanlar --set anahtar=değer)
* [gyro monitor check delete](gyro_monitor_check_delete.md)	 - delete (alanlar --set anahtar=değer)
* [gyro monitor check list](gyro_monitor_check_list.md)	 - Göster
* [gyro monitor check update](gyro_monitor_check_update.md)	 - update (alanlar --set anahtar=değer)

