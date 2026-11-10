## gyro lb service

Yük dengeleyicinin servislerini (portlarını) listele

```
gyro lb service [flags]
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

* [gyro lb](gyro_lb.md)	 - Yük dengeleyicileri listele
* [gyro lb service add](gyro_lb_service_add.md)	 - Servis ekle
* [gyro lb service delete](gyro_lb_service_delete.md)	 - Servisi sil
* [gyro lb service list](gyro_lb_service_list.md)	 - Yük dengeleyicinin servislerini (portlarını) listele
* [gyro lb service update](gyro_lb_service_update.md)	 - Servisi güncelle

