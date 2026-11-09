## gyro s3 object

Kova klasöründeki nesneleri listele

```
gyro s3 object [flags]
```

### Options

```
  -h, --help   help for object
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

* [gyro s3](gyro_s3.md)	 - S3 nesne depolama: kovalar ve erişim anahtarları
* [gyro s3 object delete](gyro_s3_object_delete.md)	 - Nesneleri sil
* [gyro s3 object download-url](gyro_s3_object_download-url.md)	 - İmzalı indirme adresi
* [gyro s3 object list](gyro_s3_object_list.md)	 - Kova klasöründeki nesneleri listele
* [gyro s3 object upload-url](gyro_s3_object_upload-url.md)	 - İmzalı yükleme adresleri

