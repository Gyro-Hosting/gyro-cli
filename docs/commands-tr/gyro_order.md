## gyro order

Siparişler (önizleme ve oluşturma)

```
gyro order [flags]
```

### Options

```
  -h, --help   help for order
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

* [gyro](gyro.md)	 - Gyro Hosting komut satırı aracı
* [gyro order create](gyro_order_create.md)	 - Sipariş Verme
* [gyro order get](gyro_order_get.md)	 - Sipariş Detayını Okuma
* [gyro order list](gyro_order_list.md)	 - Siparişleri Listeleme
* [gyro order preview](gyro_order_preview.md)	 - Siparişi Fiyatlandırma

