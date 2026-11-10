## gyro dns zone get

Zonu kayıtlarıyla getir

### Synopsis

Zonu kayıtlarıyla getir

Zon durumu, seri, DNSSEC (+DS) ve tüm kayıt kümeleri. Bekleyen zonlarda doğrulama bilgisi döner.

API: GET /cloud/dns/{id}/zones/{zone} (operationId typed_get_dns__id__zones__zone, scope services_read)

```
gyro dns zone get <ID> <ZONE> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro dns zone](gyro_dns_zone.md)	 - DNS zonlarını listele

