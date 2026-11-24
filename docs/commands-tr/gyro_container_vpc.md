## gyro container vpc

Kapsayıcı ortamının özel ağı (VPC): sabit kaynak aralığı, VPC uç noktası

### Synopsis

Bir Kapsayıcı ortamını tr-ist-1'deki özel ağınızın (VPC) bir subnetine bağlar. Ortama subnet içinde sabit bir
VPC kaynak aralığı (/28) ayrılır: kapsayıcılarınızdan VPC sunucularına giden bağlantılar buradan gelir; düğüm
değişimi, ölçekleme ya da yenilemede değişmez. İlk adres VPC uç noktasıdır: VPC sunucuları yayınladığınız
servislere (--expose servis:port) uç_nokta:port ile ulaşır. İnternetten erişilemez. Subnet en az /26 olmalı.
Ücret: düğümlü ortamlarda (Mini, Küçük, Başlangıç, Standart, Pro) ücretsiz. Sunucusuz ortamlarda ortama ayrılmış
bir bağlayıcı çalıştığı için bağlantı ücretlidir (2,99 €/ay, bağlıyken saatlik, ön ödemeli bakiyeden); attach
fiyatı TL ve € olarak gösterir ve onay ister (--yes onaylar).

```
gyro container vpc [flags]
```

### Options

```
  -h, --help   help for vpc
```

### Options inherited from parent commands

```
      --account gyro account list   panel kullanıcısı olduğunuz başka bir hesapta işlem yap (gyro account list id'si; "self" = kendi hesabınız; ortam GYRO_ACCOUNT, profil varsayılanı: `gyro account use`)
      --api-url string              API taban adresi (ortam GYRO_API_URL)
      --debug                       HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string                 mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color                    renkleri kapat (NO_COLOR da olur)
      --no-header                   tablo başlığını yazma
  -o, --output string               çıktı biçimi: table, json, yaml, text
  -p, --profile string              kullanılacak profil (ortam GYRO_PROFILE)
      --query string                sonuca uygulanacak JMESPath sorgusu
  -q, --quiet                       yalnız kimlikleri yaz
      --timeout duration            toplam zaman aşımı, ör. 30s, 10m
      --wide                        tüm tablo sütunlarını göster
  -y, --yes                         onay sorma
```

### SEE ALSO

* [gyro container](gyro_container.md)	 - Gyro Kapsayıcı (ECS), İstanbul: servisler, sürümler, loglar, alan adları
* [gyro container vpc attach](gyro_container_vpc_attach.md)	 - Ortamı bir VPC subnetine bağla
* [gyro container vpc detach](gyro_container_vpc_detach.md)	 - Ortamın VPC bağını kaldır (/28 kaynak aralığı serbest kalır)
* [gyro container vpc expose](gyro_container_vpc_expose.md)	 - VPC uç noktasında yayınlanan servisleri değiştir
* [gyro container vpc get](gyro_container_vpc_get.md)	 - Ortamın VPC bağını göster (kaynak aralığı, uç nokta, yayınlanan servisler)

