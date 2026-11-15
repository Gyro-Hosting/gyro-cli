## gyro container vpc expose

VPC uç noktasında yayınlanan servisleri değiştir

### Synopsis

Bir Kapsayıcı ortamını tr-ist-1'deki özel ağınızın (VPC) bir subnetine bağlar. Ortama subnet içinde sabit bir
VPC kaynak aralığı (/28) ayrılır: kapsayıcılarınızdan VPC sunucularına giden bağlantılar buradan gelir; düğüm
değişimi, ölçekleme ya da yenilemede değişmez. İlk adres VPC uç noktasıdır: VPC sunucuları yayınladığınız
servislere (--expose servis:port) uç_nokta:port ile ulaşır. İnternetten erişilemez. Subnet en az /26 olmalı.
Ücret: düğümlü ortamlarda (Mini, Küçük, Başlangıç, Standart, Pro) ücretsiz. Sunucusuz ortamlarda ortama ayrılmış
bir bağlayıcı çalıştığı için bağlantı ücretlidir (2,99 €/ay, bağlıyken saatlik, ön ödemeli bakiyeden); attach
fiyatı TL ve € olarak gösterir ve onay ister (--yes onaylar).

```
gyro container vpc expose <ENVIRONMENT_ID> --expose service:port ... | --clear [flags]
```

### Examples

```
  gyro container vpc expose 17000 --expose api:8080 --expose web:80
  gyro container vpc expose 17000 --clear
```

### Options

```
      --clear                yayınlanan tüm servisleri kaldır
      --dry-run              hiçbir şeyi değiştirmeden isteği göster
      --expose stringArray   servis:port (tekrarlanabilir; liste eskisinin yerine geçer)
  -h, --help                 help for expose
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

* [gyro container vpc](gyro_container_vpc.md)	 - Kapsayıcı ortamının özel ağı (VPC): sabit kaynak aralığı, VPC uç noktası

