## gyro container service

Kapsayıcı servislerini listele

```
gyro container service [flags]
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

* [gyro container](gyro_container.md)	 - Gyro Kapsayıcı (ECS), İstanbul: servisler, sürümler, loglar, alan adları
* [gyro container service autoscale](gyro_container_service_autoscale.md)	 - Servisin otomatik ölçekleme / uyarı eşiği
* [gyro container service create](gyro_container_service_create.md)	 - Kapsayıcı servisi oluştur
* [gyro container service delete](gyro_container_service_delete.md)	 - Kapsayıcı servisini sil
* [gyro container service get](gyro_container_service_get.md)	 - Kapsayıcı servisini getir
* [gyro container service list](gyro_container_service_list.md)	 - Kapsayıcı servislerini listele
* [gyro container service restart](gyro_container_service_restart.md)	 - Servisi yeniden başlat
* [gyro container service revisions](gyro_container_service_revisions.md)	 - Sürüm geçmişi
* [gyro container service rollback](gyro_container_service_rollback.md)	 - Önceki sürüme dön
* [gyro container service scale](gyro_container_service_scale.md)	 - Servisi ölçekle
* [gyro container service update](gyro_container_service_update.md)	 - Yeni sürüm dağıt / ayar değiştir

