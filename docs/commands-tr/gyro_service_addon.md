## gyro service addon

Hizmetin ek hizmetleri

```
gyro service addon [flags]
```

### Options

```
  -h, --help   help for addon
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
* [gyro service addon available](gyro_service_addon_available.md)	 - Bu hizmete alınabilecek ek hizmetler
* [gyro service addon create](gyro_service_addon_create.md)	 - Ek hizmet satın al (bakiyeden ödenir)
* [gyro service addon delete](gyro_service_addon_delete.md)	 - Ön ödemeli ek hizmeti hemen kaldır
* [gyro service addon list](gyro_service_addon_list.md)	 - Hizmetin ek hizmetleri

