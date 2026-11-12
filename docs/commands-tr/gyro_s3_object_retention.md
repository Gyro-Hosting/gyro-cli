## gyro s3 object retention

Nesne sürümünün saklaması

```
gyro s3 object retention [flags]
```

### Options

```
  -h, --help   help for retention
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

* [gyro s3 object](gyro_s3_object.md)	 - Kova klasöründeki nesneleri listele
* [gyro s3 object retention get](gyro_s3_object_retention_get.md)	 - Nesne sürümünün saklaması
* [gyro s3 object retention set](gyro_s3_object_retention_set.md)	 - Nesne sürümünü kilitle

