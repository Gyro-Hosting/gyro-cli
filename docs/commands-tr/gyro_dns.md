## gyro dns

Alan adlarınızın DNS kayıtları

```
gyro dns [flags]
```

### Options

```
  -h, --help   help for dns
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
* [gyro dns acme-challenge](gyro_dns_acme-challenge.md)	 - ACME doğrulama TXT değeri ekle
* [gyro dns acme-key](gyro_dns_acme-key.md)	 - ACME anahtarlarını listele
* [gyro dns dnssec](gyro_dns_dnssec.md)	 - DNSSEC kapat
* [gyro dns health](gyro_dns_health.md)	 - Sağlık kontrollü kayıtları listele
* [gyro dns record](gyro_dns_record.md)	 - DNS kayıtlarını listele, ekle, değiştir, sil
* [gyro dns rrset](gyro_dns_rrset.md)	 - Kayıt kümelerini listele
* [gyro dns zone](gyro_dns_zone.md)	 - DNS zone'larını listele

