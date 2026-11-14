## gyro container egress-ips

Kapsayıcı ortamının çıkış IP'leri

### Synopsis

Kapsayıcı ortamının çıkış IP'leri

Kapsayıcıların internete çıktığı güncel IPv4 adresleri - harici bir servisin IP beyaz listesine ekleyin. **Ayrılmış** ortamın çıkış IP'leri = kendi düğümlerinin IP'leri + güncelleme sırasında kullanılan paylaşımlı havuz düğümü (yeni kopya yaklaşık bir dakika orada başlar, sonra düğümünüze döner). **Sunucusuz** = paylaşımlı yalıtılmış havuz düğümleri (havuz büyürse yeni adres eklenebilir). Gyro Pipeline derleme makineleri ayrı listelenir (`build_ips`). Düğümler değişince (paket değişimi, düğüm yenileme, havuz ölçeklemesi) liste değişir ve ortam sahibine bildirim gider (Uyarılar kanalları, `egress` olayı). IPv6 çıkışı yok; TCP 25 kapalı.

API: GET /cloud/containers/{id}/egress-ips (operationId getContainerEgressIps, scope services_read)

```
gyro container egress-ips <ID> [flags]
```

### Options

```
  -h, --help   help for egress-ips
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

