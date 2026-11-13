## gyro container pipeline

Pipeline'ları listele

```
gyro container pipeline [flags]
```

### Options

```
  -h, --help   help for pipeline
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
* [gyro container pipeline approve](gyro_container_pipeline_approve.md)	 - Dağıtımı onayla
* [gyro container pipeline cancel](gyro_container_pipeline_cancel.md)	 - Çalıştırmayı durdur
* [gyro container pipeline create](gyro_container_pipeline_create.md)	 - Pipeline oluştur
* [gyro container pipeline delete](gyro_container_pipeline_delete.md)	 - Pipeline'ı sil
* [gyro container pipeline get](gyro_container_pipeline_get.md)	 - Pipeline'ı getir
* [gyro container pipeline list](gyro_container_pipeline_list.md)	 - Pipeline'ları listele
* [gyro container pipeline logs](gyro_container_pipeline_logs.md)	 - Çalıştırma logu
* [gyro container pipeline redeploy](gyro_container_pipeline_redeploy.md)	 - Bu çalıştırmanın imajını yeniden dağıt (geri dönüş)
* [gyro container pipeline reject](gyro_container_pipeline_reject.md)	 - Dağıtımı reddet
* [gyro container pipeline rotate-secret](gyro_container_pipeline_rotate-secret.md)	 - Webhook gizli anahtarını yenile
* [gyro container pipeline run](gyro_container_pipeline_run.md)	 - Çalıştırma başlat
* [gyro container pipeline run-get](gyro_container_pipeline_run-get.md)	 - Çalıştırmayı getir
* [gyro container pipeline runs](gyro_container_pipeline_runs.md)	 - Pipeline çalıştırmaları
* [gyro container pipeline secret](gyro_container_pipeline_secret.md)	 - Derleme gizli değerini sil
* [gyro container pipeline test](gyro_container_pipeline_test.md)	 - Depo erişimini test et
* [gyro container pipeline update](gyro_container_pipeline_update.md)	 - Pipeline'ı değiştir
* [gyro container pipeline webhook](gyro_container_pipeline_webhook.md)	 - Webhook adresi ve gizli anahtar

