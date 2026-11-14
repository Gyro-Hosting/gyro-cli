## gyro vpc port-forward add

Kapının dış portunu yönlendir

### Synopsis

Kapının dış portunu yönlendir

Kapının public_ip:public_port adresi -> VPC içinde target_ip:target_port. Kapı açık olmalı (409 `gateway_disabled`); protokol + dış port başına bir yönlendirme; en çok 20.

API: POST /cloud/vpcs/{vpc_id}/port-forwards (operationId createVpcPortForward, scope cloud_write)

```
gyro vpc port-forward add <VPC-ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for add
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --protocol string          protocol (tcp|udp)
      --public-port int          public_port
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --target-ip string         target_ip
      --target-port int          target_port
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

