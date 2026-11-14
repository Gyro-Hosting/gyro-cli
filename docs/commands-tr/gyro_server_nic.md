## gyro server nic

Sunucunun özel ağ kartları

```
gyro server nic [flags]
```

### Options

```
  -h, --help   help for nic
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

* [gyro server](gyro_server.md)	 - Bulut sunucular: oluşturma, güç, snapshot, SSH ...
* [gyro server nic attach](gyro_server_nic_attach.md)	 - Sunucuyu subnete bağla
* [gyro server nic detach](gyro_server_nic_detach.md)	 - Özel ağ kartını çıkar
* [gyro server nic list](gyro_server_nic_list.md)	 - Sunucunun özel ağ kartları

