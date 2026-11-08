## gyro plan list

Sipariş verilebilen planları fiyatlarıyla listele

```
gyro plan list [flags]
```

### Examples

```
  gyro plan list --location falkenstein
  gyro plan list --cycle annually -o json
```

### Options

```
      --all                  stokta olmayanları da göster
      --cycle string         monthly|quarterly|semiannually|annually (default "monthly")
  -h, --help                 help for list
      --location string      lokasyon kodu ya da şehir
      --server-type string   sunucu tipi, ör. cx23
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

* [gyro plan](gyro_plan.md)	 - Sunucu planları ve fiyatlar (katalog)

