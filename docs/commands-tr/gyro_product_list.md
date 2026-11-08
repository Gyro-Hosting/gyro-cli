## gyro product list

Ürünleri Listeleme

### Synopsis

Ürünleri Listeleme

Sipariş edilebilecek ürünleri vitrin sırasıyla döndürür. — cüzdan para birimi

API: GET /products (operationId coreGetProducts, scope Products/GetProducts)

```
gyro product list [flags]
```

### Options

```
      --all             tüm sayfaları getir
      --category int    Kategori numarası. Ana kategoriyle de ek atamalarla da eşleşir.
  -h, --help            help for list
      --limit int       sayfa başına kayıt (en çok 100)
      --page int        sayfa numarası (default 1)
      --search string   Ürün başlığında arar.
      --type string     Tek bir ürün türü. Bilinmeyen tür reddedilir ve geçerli liste yanıtın ayrıntısında gelir.
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

* [gyro product](gyro_product.md)	 - Tüm ürün kataloğu (tüm ürün hatları)

