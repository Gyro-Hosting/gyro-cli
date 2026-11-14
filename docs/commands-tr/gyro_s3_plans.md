## gyro s3 plans

Nesne depolama paketlerini listele

### Synopsis

Nesne depolama paketlerini listele

Sipariş edilebilir Nesne Depolama (`standard-*`) ve Korumalı Depolama (`protected-*`) paketleri: sınırlar, dönem başına fiyat (sabit EUR, KDV hariç) ve ek hizmetler: ek trafik paketi (`1tb`, `5tb`, tek seferlik, bu takvim ayı) ve nesne kilitli kova hakkı (korumalı paketler). `available: false` = şu an kapasite yok.

API: GET /cloud/storage/plans (operationId listStoragePlans, scope cloud_read)

```
gyro s3 plans [flags]
```

### Options

```
  -h, --help   help for plans
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

