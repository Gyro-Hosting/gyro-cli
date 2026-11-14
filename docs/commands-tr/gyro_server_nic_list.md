## gyro server nic list

Sunucunun özel ağ kartları

### Synopsis

Sunucunun özel ağ kartları

tr-ist-1 sunucusunun özel ağ kartları. Takma sıcak yapılır (yeniden başlatma yok): açık sunucuda yaklaşık 10-60 sn'de hazır (`state` attaching -> ready); kapalı sunucuda ayarlar açılışta uygulanır. Bir sunucu aynı anda tek VPC'ye bağlanır (409 `other_vpc`), en çok 3 kart. `ip` boş = otomatik.

API: GET /cloud/servers/{id}/nics (operationId listServerNics, scope cloud_read)

```
gyro server nic list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro server nic](gyro_server_nic.md)	 - Sunucunun özel ağ kartları

