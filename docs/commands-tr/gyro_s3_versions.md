## gyro s3 versions

Nesne sürümlerini listele

```
gyro s3 versions [flags]
```

### Options

```
  -h, --help   help for versions
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
* [gyro s3 versions delete](gyro_s3_versions_delete.md)	 - Sürümü kalıcı sil
* [gyro s3 versions list](gyro_s3_versions_list.md)	 - Nesne sürümlerini listele
* [gyro s3 versions restore](gyro_s3_versions_restore.md)	 - Eski sürümü geri yükle
* [gyro s3 versions undelete](gyro_s3_versions_undelete.md)	 - Silme işaretini kaldır

