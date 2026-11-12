## gyro template

Şablonlarınızı listele

```
gyro template [flags]
```

### Options

```
  -h, --help   help for template
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
* [gyro template create](gyro_template_create.md)	 - Sunucudan şablon oluştur
* [gyro template delete](gyro_template_delete.md)	 - Şablonu sil
* [gyro template deploy](gyro_template_deploy.md)	 - Şablondan yeni sunucu
* [gyro template get](gyro_template_get.md)	 - Şablonu getir
* [gyro template list](gyro_template_list.md)	 - Şablonlarınızı listele
* [gyro template rebuild](gyro_template_rebuild.md)	 - Sunucuyu şablonla yeniden kur
* [gyro template update](gyro_template_update.md)	 - Şablonu yeniden adlandır

