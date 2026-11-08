## gyro product category list

Kategorileri Listeleme

### Synopsis

Kategorileri Listeleme

Vitrinde görünen ürün kategorilerini düz bir liste olarak döndürür. — sayfalama yok

API: GET /products/categories (operationId coreGetProductCategories, scope Products/GetProductCategories)

```
gyro product category list [flags]
```

### Options

```
  -h, --help          help for list
      --type string   Yalnız bu türün kategorileri. Bilinmeyen değer boş liste verir, hata değil.
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

* [gyro product category](gyro_product_category.md)	 - Kategorileri Listeleme

