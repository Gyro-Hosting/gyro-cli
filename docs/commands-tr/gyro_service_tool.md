## gyro service tool

Hizmetin araçlarını, eylemlerini ve güvenlik politikasını keşfet

```
gyro service tool [flags]
```

### Options

```
  -h, --help   help for tool
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

* [gyro service](gyro_service.md)	 - Tüm ürün hatlarındaki hizmetler (genel köprü)
* [gyro service tool get](gyro_service_tool_get.md)	 - Bir aracı oku (posta kutuları, snapshot'lar, DNS kayıtları, ...)
* [gyro service tool list](gyro_service_tool_list.md)	 - Hizmetin araçlarını, eylemlerini ve güvenlik politikasını keşfet
* [gyro service tool run](gyro_service_tool_run.md)	 - Bir araç eylemi çalıştır (posta kutusu oluştur, snapshot'tan dön, ...)

