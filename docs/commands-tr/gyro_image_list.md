## gyro image list

İşletim sistemi imajlarını listele

### Synopsis

İşletim sistemi imajlarını listele

API: GET /cloud/images (operationId listImages, scope cloud_read)

```
gyro image list [flags]
```

### Options

```
  -h, --help             help for list
      --plan string      region ile: bu bölgedeki bu planın imajları.
      --product-id int   Yalnız bu üründe sunulan imajlar (ürünün bölgesini de seçer).
      --region string    Yalnız bu bölgede sunulan imajlar (ör. tr-ist-1). Verilmezse küresel bulut kataloğunun imajları (değişmedi).
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

* [gyro image](gyro_image.md)	 - İşletim sistemi imajları

