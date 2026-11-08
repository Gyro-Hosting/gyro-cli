## gyro s3 key

Erişim anahtarları (gizli anahtar bir kez gösterilir)

```
gyro s3 key [flags]
```

### Options

```
  -h, --help   help for key
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
* [gyro s3 key create](gyro_s3_key_create.md)	 - create (alanlar --set anahtar=değer)
* [gyro s3 key delete](gyro_s3_key_delete.md)	 - delete (alanlar --set anahtar=değer)
* [gyro s3 key list](gyro_s3_key_list.md)	 - Göster
* [gyro s3 key update](gyro_s3_key_update.md)	 - update (alanlar --set anahtar=değer)

