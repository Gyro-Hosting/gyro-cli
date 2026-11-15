## gyro container schedule

Zamanlanmış görevleri sonraki ve son çalışmalarıyla listele

```
gyro container schedule [flags]
```

### Options

```
  -h, --help   help for schedule
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
* [gyro container schedule create](gyro_container_schedule_create.md)	 - Zamanlanmış görev oluştur (cron, her N dakikada ya da tek sefer)
* [gyro container schedule delete](gyro_container_schedule_delete.md)	 - Zamanlanmış görevi sil
* [gyro container schedule get](gyro_container_schedule_get.md)	 - Zamanlanmış görevi getir
* [gyro container schedule list](gyro_container_schedule_list.md)	 - Zamanlanmış görevleri sonraki ve son çalışmalarıyla listele
* [gyro container schedule pause](gyro_container_schedule_pause.md)	 - Zamanlanmış görevi duraklat
* [gyro container schedule preview](gyro_container_schedule_preview.md)	 - Zamanlamanın sonraki çalışma saatleri
* [gyro container schedule resume](gyro_container_schedule_resume.md)	 - Zamanlanmış görevi devam ettir
* [gyro container schedule run-now](gyro_container_schedule_run-now.md)	 - Zamanlanmış görevi şimdi çalıştır
* [gyro container schedule update](gyro_container_schedule_update.md)	 - Zamanlanmış görevi değiştir (yalnız verilen seçenekler değişir)

