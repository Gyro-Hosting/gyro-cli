## gyro service

Tüm ürün hatlarındaki hizmetler (genel köprü)

```
gyro service [flags]
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

* [gyro](gyro.md)	 - Gyro Hosting komut satırı aracı
* [gyro service addon](gyro_service_addon.md)	 - Hizmet Eklentilerini Okuma
* [gyro service addons](gyro_service_addons.md)	 - Ek hizmet satın al (bakiyeden ödenir)
* [gyro service cancellation](gyro_service_cancellation.md)	 - İptal Talebi Açma
* [gyro service dashboard](gyro_service_dashboard.md)	 - Panel Özetini Okuma
* [gyro service get](gyro_service_get.md)	 - Hizmeti getir (yenileme, iptal, erişim bilgileri)
* [gyro service invoice](gyro_service_invoice.md)	 - Hizmetin Faturalarını Okuma
* [gyro service label](gyro_service_label.md)	 - Hizmetin takma adını ve etiketlerini getir
* [gyro service list](gyro_service_list.md)	 - Tüm ürün hatlarındaki hizmetleri listele
* [gyro service method](gyro_service_method.md)	 - Müşteri-çağrılabilir modül metodunu çağır (örn. lisans change_ip)
* [gyro service metrics](gyro_service_metrics.md)	 - Ölçümleri Okuma
* [gyro service module-method](gyro_service_module-method.md)	 - Modül Metodu Çalıştırma
* [gyro service password](gyro_service_password.md)	 - Panel Parolasını Değiştirme
* [gyro service renew](gyro_service_renew.md)	 - Yenileme Faturası Kesme
* [gyro service resize](gyro_service_resize.md)	 - Paketi yükselt (bakiyeden ödenir)
* [gyro service resize-option](gyro_service_resize-option.md)	 - Hizmetin yükseltilebileceği ürünler ve fiyatları
* [gyro service sso](gyro_service_sso.md)	 - Panele Tek Oturum Açma
* [gyro service tool](gyro_service_tool.md)	 - Hizmetin araçlarını, eylemlerini ve güvenlik politikasını keşfet
* [gyro service update](gyro_service_update.md)	 - Hizmet Tercihlerini Değiştirme

