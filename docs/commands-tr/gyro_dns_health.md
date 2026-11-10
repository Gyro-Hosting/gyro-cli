## gyro dns health

Sağlık kontrollü kayıtları listele

```
gyro dns health [flags]
```

### Options

```
  -h, --help   help for health
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
* [gyro dns health create](gyro_dns_health_create.md)	 - Sağlık kontrollü kayıt oluştur
* [gyro dns health delete](gyro_dns_health_delete.md)	 - Sağlık kontrolünü sil
* [gyro dns health list](gyro_dns_health_list.md)	 - Sağlık kontrollü kayıtları listele

