## gyro server power

Sunucuyu aç, kapat, ACPI kapat, yeniden başlat ya da resetle

```
gyro server power [flags]
```

### Options

```
  -h, --help   help for power
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

* [gyro server](gyro_server.md)	 - Bulut sunucular: oluşturma, güç, snapshot, SSH ...
* [gyro server power off](gyro_server_power_off.md)	 - Elektriği kes
* [gyro server power on](gyro_server_power_on.md)	 - Aç
* [gyro server power reboot](gyro_server_power_reboot.md)	 - Yeniden başlat (düzgün)
* [gyro server power reset](gyro_server_power_reset.md)	 - Donanımsal reset
* [gyro server power shutdown](gyro_server_power_shutdown.md)	 - Düzgün kapat (ACPI)

