## gyro container registry

Ortamın Gyro imaj deposu

```
gyro container registry [flags]
```

### Options

```
  -h, --help   help for registry
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

* [gyro container](gyro_container.md)	 - Gyro Kapsayıcı (ECS), İstanbul: servisler, sürümler, loglar, alan adları
* [gyro container registry credentials](gyro_container_registry_credentials.md)	 - Depo parolası
* [gyro container registry get](gyro_container_registry_get.md)	 - Ortamın Gyro imaj deposu
* [gyro container registry image-delete](gyro_container_registry_image-delete.md)	 - İmaj etiketini sil
* [gyro container registry images](gyro_container_registry_images.md)	 - Depodaki imajlar
* [gyro container registry import](gyro_container_registry_import.md)	 - Başka bir depodan imaj içe aktar
* [gyro container registry imports](gyro_container_registry_imports.md)	 - İmaj içe aktarmaları (ilerleme ve geçmiş)
* [gyro container registry policy](gyro_container_registry_policy.md)	 - İmaj temizlik ayarı
* [gyro container registry rotate](gyro_container_registry_rotate.md)	 - Depo parolasını yenile

