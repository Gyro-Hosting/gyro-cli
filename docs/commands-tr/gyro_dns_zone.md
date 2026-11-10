## gyro dns zone

DNS zonlarını listele

```
gyro dns zone [flags]
```

### Options

```
  -h, --help   help for zone
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

* [gyro dns](gyro_dns.md)	 - Alan adlarınızın DNS kayıtları
* [gyro dns zone create](gyro_dns_zone_create.md)	 - Zon ekle
* [gyro dns zone delete](gyro_dns_zone_delete.md)	 - Zonu sil
* [gyro dns zone export](gyro_dns_zone_export.md)	 - Zon dosyasını dışa aktar
* [gyro dns zone get](gyro_dns_zone_get.md)	 - Zonu kayıtlarıyla getir
* [gyro dns zone import](gyro_dns_zone_import.md)	 - BIND zon dosyası içe aktar
* [gyro dns zone list](gyro_dns_zone_list.md)	 - DNS zonlarını listele
* [gyro dns zone verify](gyro_dns_zone_verify.md)	 - Sahipliği şimdi doğrula

