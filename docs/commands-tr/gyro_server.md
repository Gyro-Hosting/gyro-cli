## gyro server

Bulut sunucular: oluşturma, güç, snapshot, SSH ...

```
gyro server [flags]
```

### Options

```
  -h, --help   help for server
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
* [gyro server activity](gyro_server_activity.md)	 - Sunucudaki son platform işlemleri
* [gyro server app](gyro_server_app.md)	 - Uygulama kurulum durumu (hazır olunca kimlik bir kez)
* [gyro server backup](gyro_server_backup.md)	 - Otomatik yedekleri listele
* [gyro server console](gyro_server_console.md)	 - Tek seferlik VNC konsol bileti al
* [gyro server create](gyro_server_create.md)	 - Yeni bulut sunucu siparişi (bakiyeden ödenir)
* [gyro server delete](gyro_server_delete.md)	 - Sunucuyu sil (iptal) - veri bekleme süresi sonunda silinir
* [gyro server firewall](gyro_server_firewall.md)	 - Gelen trafik güvenlik duvarı kuralları
* [gyro server get](gyro_server_get.md)	 - Sunucuyu canlı durumuyla göster
* [gyro server ip](gyro_server_ip.md)	 - Sunucunun genel IPv4 adresini (ya da --ipv6) yaz
* [gyro server list](gyro_server_list.md)	 - Sunucularınızı listeleyin
* [gyro server metrics](gyro_server_metrics.md)	 - CPU, disk ve ağ ölçümleri (özet)
* [gyro server power](gyro_server_power.md)	 - Sunucuyu aç, kapat, ACPI kapat, yeniden başlat ya da resetle
* [gyro server rdns](gyro_server_rdns.md)	 - Ters DNS kayıtlarını listele
* [gyro server rebuild](gyro_server_rebuild.md)	 - İşletim sistemini yeniden kur - DİSK SİLİNİR
* [gyro server rename](gyro_server_rename.md)	 - Sunucu adını değiştir
* [gyro server rescue](gyro_server_rescue.md)	 - Kurtarma sistemi (kurtarma Linux'u ile açılış)
* [gyro server reset-password](gyro_server_reset-password.md)	 - Yeni rastgele root parolası (bir kez gösterilir)
* [gyro server snapshot](gyro_server_snapshot.md)	 - Snapshot'ları ve kotayı listele
* [gyro server ssh](gyro_server_ssh.md)	 - Sunucuya SSH bağlantısı aç
* [gyro server storage-mount](gyro_server_storage-mount.md)	 - Sunucuya bağlı kovalar
* [gyro server update](gyro_server_update.md)	 - Ad ya da otomatik yenilemeyi değiştir
* [gyro server wait](gyro_server_wait.md)	 - Sunucu kurulup boşa çıkana kadar bekle

