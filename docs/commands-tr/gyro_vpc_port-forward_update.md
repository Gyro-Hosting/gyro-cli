## gyro vpc port-forward update

Port yönlendirmeyi kimin kullanabileceğini değiştir (kaynak izin listesi)

### Synopsis

Yönlendirmeyi yalnız verilen public IPv4 ağları kullanabilir (tek IP /32 olur; ağ adresi yazılır, ör.
198.51.100.0/24; en çok 20). Liste eskisinin yerine geçer; --any herkese açar.

```
gyro vpc port-forward update <VPC_ID> <PORT_FORWARD_ID> --source <IP|CIDR> ... | --any [flags]
```

### Examples

```
  gyro vpc port-forward update vpc-1a2b3c4d pf-9c8b7a6d --source 203.0.113.10 --source 198.51.100.0/24
  gyro vpc port-forward update vpc-1a2b3c4d pf-9c8b7a6d --any
```

### Options

```
      --any                  herkese izin ver (boş liste)
      --dry-run              hiçbir şeyi değiştirmeden isteği göster
  -h, --help                 help for update
      --source stringArray   izin verilen kaynak IPv4 ya da CIDR (tekrarlanabilir ya da virgülle)
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

* [gyro vpc port-forward](gyro_vpc_port-forward.md)	 - Port yönlendirmeleri listele

