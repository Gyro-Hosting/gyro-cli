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
* [gyro container create](gyro_container_create.md)	 - Kapsayıcı ortamı oluştur (toplu)
* [gyro container delete](gyro_container_delete.md)	 - Kapsayıcı ortamını sil (ön ödemeli)
* [gyro container domain](gyro_container_domain.md)	 - Servisin alan adları
* [gyro container egress-approval](gyro_container_egress-approval.md)	 - Çıkış IP değişikliği onay adımı
* [gyro container egress-change](gyro_container_egress-change.md)	 - Bekleyen çıkış IP değişikliğini onayla
* [gyro container egress-ips](gyro_container_egress-ips.md)	 - Kapsayıcı ortamının çıkış IP'leri
* [gyro container env](gyro_container_env.md)	 - Ortam değişkenlerini ayarla
* [gyro container events](gyro_container_events.md)	 - Son uyarılar
* [gyro container external-registry](gyro_container_external-registry.md)	 - Özel depo bilgisini sil
* [gyro container get](gyro_container_get.md)	 - Kapsayıcı ortamı ayrıntısı (namespace, iç adresleme, çıkış IP'leri)
* [gyro container job](gyro_container_job.md)	 - Tek seferlik görevleri listele
* [gyro container jobs](gyro_container_jobs.md)	 - Tek seferlik görevleri listele (son 50)
* [gyro container log-plans](gyro_container_log-plans.md)	 - Log paketleri
* [gyro container log-query](gyro_container_log-query.md)	 - Loglarda ara (JSON gövde)
* [gyro container log-retention](gyro_container_log-retention.md)	 - Log paketini değiştir
* [gyro container log-rules](gyro_container_log-rules.md)	 - Logdan alarm kuralları
* [gyro container log-stats](gyro_container_log-stats.md)	 - Log özeti
* [gyro container log-usage](gyro_container_log-usage.md)	 - Log paketi ve kullanım
* [gyro container logs](gyro_container_logs.md)	 - Ortamın ya da bir servisin loglarında ara, canlı izle
* [gyro container metrics](gyro_container_metrics.md)	 - Servis işlemci / bellek
* [gyro container nodes](gyro_container_nodes.md)	 - Düğüm havuzu
* [gyro container notifications](gyro_container_notifications.md)	 - Bildirim ayarları ve süren uyarılar
* [gyro container pipeline](gyro_container_pipeline.md)	 - Pipeline'ları listele
* [gyro container plans](gyro_container_plans.md)	 - Kapsayıcı ortam paketlerini listele
* [gyro container registry](gyro_container_registry.md)	 - Ortamın Gyro imaj deposu
* [gyro container run](gyro_container_run.md)	 - Servisin imajı, değişkenleri ve gizli değerleriyle tek seferlik görev çalıştır (migration, seed)
* [gyro container secret](gyro_container_secret.md)	 - Gizli değer sil
* [gyro container service](gyro_container_service.md)	 - Kapsayıcı servislerini listele
* [gyro container sizes](gyro_container_sizes.md)	 - Sunucusuz kapsayıcı boyutları (vCPU / bellek) ve fiyatları

