## gyro lb

Yük dengeleyicileri listele

```
gyro lb [flags]
```

### Options

```
  -h, --help   help for lb
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
* [gyro lb algorithm](gyro_lb_algorithm.md)	 - Dağıtım algoritmasını değiştir
* [gyro lb cert](gyro_lb_cert.md)	 - Sertifikaları listele
* [gyro lb create](gyro_lb_create.md)	 - Yük dengeleyici sipariş et
* [gyro lb delete](gyro_lb_delete.md)	 - Yük dengeleyiciyi sil (iptal)
* [gyro lb get](gyro_lb_get.md)	 - Yük dengeleyiciyi getir
* [gyro lb list](gyro_lb_list.md)	 - Yük dengeleyicileri listele
* [gyro lb metrics](gyro_lb_metrics.md)	 - Yük dengeleyici ölçümleri
* [gyro lb network](gyro_lb_network.md)	 - Özel ağa bağla
* [gyro lb rdns](gyro_lb_rdns.md)	 - Ters DNS ayarla
* [gyro lb service](gyro_lb_service.md)	 - Yük dengeleyicinin servislerini (portlarını) listele
* [gyro lb settings](gyro_lb_settings.md)	 - Algoritma, ters DNS, özel ağ
* [gyro lb target](gyro_lb_target.md)	 - Hedefleri sağlık durumuyla listele
* [gyro lb types](gyro_lb_types.md)	 - Yük dengeleyici boyutları ve fiyatları

