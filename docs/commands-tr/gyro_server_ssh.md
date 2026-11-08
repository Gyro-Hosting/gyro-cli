## gyro server ssh

Sunucuya SSH bağlantısı aç

### Synopsis

Sunucunun IP adresini bulur ve yerel ssh istemcinizi çalıştırır. "--" sonrasındaki her şey ssh'a geçer,
uzaktan komut çalıştırabilirsiniz. Varsayılanlar profilden gelir (ssh_user, ssh_identity).

```
gyro server ssh <SERVER> [-- SSH-ARGS/COMMAND...] [flags]
```

### Examples

```
  gyro server ssh web-1
  gyro server ssh 47466 -i ~/.ssh/id_ed25519 -- uptime
  gyro server ssh web-1 --ipv6 -- -L 8080:localhost:80
```

### Options

```
  -h, --help              help for ssh
  -i, --identity string   özel anahtar dosyası
  -6, --ipv6              IPv6 ile bağlan
      --port int          SSH portu (default 22)
  -u, --user string       uzak kullanıcı (varsayılan root)
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

