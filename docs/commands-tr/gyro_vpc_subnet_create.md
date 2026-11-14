## gyro vpc subnet create

Subnet ekle

### Synopsis

Subnet ekle

VPC aralığı içinde /20 - /28, başka subnetle çakışmadan. `cidr` verin ya da sıradaki boş aralık için `prefix` (varsayılan 24). Her subnette .1 kapıdır; .0-.4 ve son adres ayrılmıştır.

API: POST /cloud/vpcs/{vpc_id}/subnets (operationId createVpcSubnet, scope cloud_write)

```
gyro vpc subnet create <VPC-ID> [flags]
```

### Options

```
      --cidr string              cidr
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              name
      --prefix int               prefix
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro vpc subnet](gyro_vpc_subnet.md)	 - Subnetleri listele

