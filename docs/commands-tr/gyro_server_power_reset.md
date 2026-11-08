## gyro server power reset

Donanımsal reset

```
gyro server power reset <SERVER> [flags]
```

### Options

```
  -h, --help                    help for reset
      --wait                    sunucu running olana kadar bekle
      --wait-timeout duration   --wait için en uzun süre (default 10m0s)
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

* [gyro server power](gyro_server_power.md)	 - Sunucuyu aç, kapat, ACPI kapat, yeniden başlat ya da resetle

