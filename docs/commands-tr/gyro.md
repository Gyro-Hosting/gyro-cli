## gyro

Gyro Hosting komut satırı aracı

### Synopsis

gyro, Gyro Hosting hizmetlerinizi terminalden yönetir: bulut sunucular, SSH anahtarları, faturalar,
bakiye ve Gyro Public API'nin sunduğu her şey (API sürümü 2.18.0).

Başlarken:
  gyro configure                 # API anahtarını kaydedin (Müşteri paneli → sağ üstteki hesap menüsü → API Kimlik Bilgileri)
  gyro server list
  gyro server create --plan standart-2-4 --location falkenstein --image ubuntu-24.04 --ssh-key anahtarim --wait
  gyro server ssh web-1

Para harcayan ve geri alınamayan işlemler onay sorar (betiklerde --yes, fiyat için --dry-run).

### Options

```
      --api-url string     API taban adresi (ortam GYRO_API_URL)
      --debug              HTTP isteklerini stderr'e yaz (sırlar maskeli)
  -h, --help               help for gyro
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

* [gyro activity](gyro_activity.md)	 - API anahtarlarınızın yazma kaydı
* [gyro address](gyro_address.md)	 - Fatura adresleri
* [gyro affiliate](gyro_affiliate.md)	 - Program Durumunu Okuma
* [gyro api](gyro_api.md)	 - Herhangi bir API ucunu doğrudan çağır ('gh api' gibi)
* [gyro app](gyro_app.md)	 - Tek tık uygulamalar
* [gyro balance](gyro_balance.md)	 - Hesap bakiyesi (tek başına) ve hareketler
* [gyro billing](gyro_billing.md)	 - Ön ödemeli saatlik kullanım (İstanbul bulut, kapsayıcı): bakiye, kullanım, tahmin, döküm
* [gyro bulk](gyro_bulk.md)	 - Bir eylemi 50 sunucuya kadar toplu çalıştır
* [gyro completion](gyro_completion.md)	 - Kabuk tamamlama betiği üret
* [gyro configure](gyro_configure.md)	 - API anahtarını ve tercihleri adlı bir profile kaydet
* [gyro container](gyro_container.md)	 - Gyro Kapsayıcı (ECS), İstanbul: servisler, sürümler, loglar, alan adları
* [gyro context](gyro_context.md)	 - Etkin profili, API adresini ve anahtarı göster (anahtarı doğrula)
* [gyro dedicated](gyro_dedicated.md)	 - Fiziksel sunucular
* [gyro dns](gyro_dns.md)	 - Alan adlarınızın DNS kayıtları
* [gyro domain](gyro_domain.md)	 - Alan adları: sorgu, kayıt, transfer, ad sunucuları, DNS, WHOIS
* [gyro firewall](gyro_firewall.md)	 - Ağ güvenlik duvarı politikası (kurallar, varsayılan kural, durum)
* [gyro image](gyro_image.md)	 - İşletim sistemi imajları
* [gyro invoice](gyro_invoice.md)	 - Faturalar
* [gyro key](gyro_key.md)	 - Alt hesap API anahtarları
* [gyro lb](gyro_lb.md)	 - Yük dengeleyicileri listele
* [gyro license](gyro_license.md)	 - Yazılım lisansları (cPanel, Plesk, LiteSpeed ...)
* [gyro location](gyro_location.md)	 - Lokasyonlar
* [gyro mail](gyro_mail.md)	 - Kurumsal mail: kutular, takma adlar, yönlendirme ...
* [gyro me](gyro_me.md)	 - Profili Okuma
* [gyro monitor](gyro_monitor.md)	 - Erişilebilirlik izleme
* [gyro order](gyro_order.md)	 - Siparişler (önizleme ve oluşturma)
* [gyro ping](gyro_ping.md)	 - Ping
* [gyro plan](gyro_plan.md)	 - Sunucu planları ve fiyatlar (katalog)
* [gyro product](gyro_product.md)	 - Tüm ürün kataloğu (tüm ürün hatları)
* [gyro project](gyro_project.md)	 - Projeleri listele
* [gyro reference](gyro_reference.md)	 - Ülkeler, iller ve ilçeler
* [gyro reseller](gyro_reseller.md)	 - Bayi: uç-müşteriler, atamalar, anahtarlar, rapor
* [gyro s3](gyro_s3.md)	 - S3 nesne depolama: kovalar ve erişim anahtarları
* [gyro security-group](gyro_security-group.md)	 - Güvenlik gruplarını listele
* [gyro server](gyro_server.md)	 - Bulut sunucular: oluşturma, güç, snapshot, SSH ...
* [gyro service](gyro_service.md)	 - Tüm ürün hatlarındaki hizmetler (genel köprü)
* [gyro sms](gyro_sms.md)	 - SMS mesajları
* [gyro ssh-key](gyro_ssh-key.md)	 - Hesabınızdaki SSH açık anahtarları
* [gyro template](gyro_template.md)	 - Şablonlarınızı listele
* [gyro ticket](gyro_ticket.md)	 - Destek talepleri
* [gyro user](gyro_user.md)	 - Alt kullanıcıları listele
* [gyro vds](gyro_vds.md)	 - VDS sunucular
* [gyro version](gyro_version.md)	 - Sürümü yaz (--check ile yeni sürüm denetimi)
* [gyro webhook](gyro_webhook.md)	 - Webhook'lar (imzalı olay bildirimleri)
* [gyro whoami](gyro_whoami.md)	 - Kullanılan API anahtarını ve yetkilerini göster

