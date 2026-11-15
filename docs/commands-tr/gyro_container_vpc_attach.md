## gyro container vpc attach

Ortamı bir VPC subnetine bağla

### Synopsis

Düğümlü bir Kapsayıcı ortamını (Mini, Küçük, Başlangıç, Standart, Pro) tr-ist-1'deki özel ağınızın (VPC) bir
subnetine bağlar. Ortama subnet içinde sabit bir VPC kaynak aralığı (/28) ayrılır: kapsayıcılarınızdan VPC
sunucularına giden bağlantılar buradan gelir; düğüm değişimi, ölçekleme ya da yenilemede değişmez. İlk adres VPC
uç noktasıdır: VPC sunucuları yayınladığınız servislere (--expose servis:port) uç_nokta:port ile ulaşır.
İnternetten erişilemez. Ücretsizdir. Subnet en az /26 olmalı. Sunucusuz: yakında.

```
gyro container vpc attach <ENVIRONMENT_ID> --subnet <SUBNET_ID> [--vpc <VPC_ID>] [--expose service:port ...] [flags]
```

### Examples

```
  gyro container vpc attach 17000 --vpc vpc-1a2b3c4d --subnet sn-5e6f7a8b --expose api:8080 --wait
  gyro vpc subnet list vpc-1a2b3c4d    # subnet ids
```

### Options

```
      --dry-run                 hiçbir şeyi değiştirmeden isteği göster
      --expose stringArray      servisi VPC uç noktasında yayınla: servis:port (tekrarlanabilir, en çok 20)
  -h, --help                    help for attach
      --subnet string           subnet kimliği (en az /26)
      --vpc string              VPC kimliği (gyro vpc list); subnetle eşleşmesi denetlenir
      --wait                    düğümler subnete katılana kadar bekle (state ready)
      --wait-timeout duration   --wait için en uzun süre (default 10m0s)
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

* [gyro container vpc](gyro_container_vpc.md)	 - Düğümlü ortamın özel ağı (VPC): sabit kaynak aralığı, VPC uç noktası

