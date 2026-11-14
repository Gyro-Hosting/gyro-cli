## gyro vpc list

Özel ağları listele

### Synopsis

Özel ağları listele

Yalnız tr-ist-1. Sınırlar (GET /cloud/vpcs `limits`, destek talebiyle artırılır): hesap başına 3 VPC, VPC başına 4 subnet, sunucu başına 3 özel kart, VPC başına 20 port yönlendirme; bir sunucu aynı anda tek VPC'ye (o VPC'nin birden çok subnetine) bağlanır. Başka hesabın nesnesi 404 döner. `limits` ve `pricing` de döner.

API: GET /cloud/vpcs (operationId listVpcs, scope cloud_read)

```
gyro vpc list [flags]
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

* [gyro vpc](gyro_vpc.md)	 - Özel ağları listele

