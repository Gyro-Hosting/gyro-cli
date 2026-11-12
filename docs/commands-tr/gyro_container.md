## gyro container

Gyro Kapsayıcı (ECS), İstanbul: servisler, sürümler, loglar, alan adları

```
gyro container [flags]
```

### Options

```
  -h, --help   help for container
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
* [gyro container domain](gyro_container_domain.md)	 - Servisin alan adları
* [gyro container env](gyro_container_env.md)	 - Ortam değişkenlerini ayarla
* [gyro container events](gyro_container_events.md)	 - Son uyarılar
* [gyro container external-registry](gyro_container_external-registry.md)	 - Özel depo bilgisini sil
* [gyro container logs](gyro_container_logs.md)	 - Servis logları
* [gyro container metrics](gyro_container_metrics.md)	 - Servis işlemci / bellek
* [gyro container nodes](gyro_container_nodes.md)	 - Düğüm havuzu
* [gyro container notifications](gyro_container_notifications.md)	 - Bildirim ayarları ve süren uyarılar
* [gyro container registry](gyro_container_registry.md)	 - Ortamın Gyro imaj deposu
* [gyro container secret](gyro_container_secret.md)	 - Gizli değer sil
* [gyro container service](gyro_container_service.md)	 - Kapsayıcı servislerini listele
* [gyro container sizes](gyro_container_sizes.md)	 - Sunucusuz kapsayıcı boyutları (vCPU / bellek) ve fiyatları

