## gyro reseller service

Hizmetin bayi görünümü (durum, bitiş, açık faturalar)

```
gyro reseller service [flags]
```

### Options

```
  -h, --help   help for service
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

* [gyro reseller](gyro_reseller.md)	 - Bayi: uç-müşteriler, atamalar, anahtarlar, rapor
* [gyro reseller service get](gyro_reseller_service_get.md)	 - Hizmetin bayi görünümü (durum, bitiş, açık faturalar)
* [gyro reseller service suspend](gyro_reseller_service_suspend.md)	 - Hizmeti askıya al (bayi)
* [gyro reseller service terminate](gyro_reseller_service_terminate.md)	 - Hizmeti sonlandır (bayi)
* [gyro reseller service unsuspend](gyro_reseller_service_unsuspend.md)	 - Hizmeti askıdan çıkar (bayi)

