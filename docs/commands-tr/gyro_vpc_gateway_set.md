## gyro vpc gateway set

VPC kapısını aç/kapat, ayarla

### Synopsis

VPC kapısını aç/kapat, ayarla

Kapı; public IP'si kapalı sunucuların internete çıkışını (NAT), port yönlendirmeyi ve subnetler arası kuralı sağlar. Açıkken ön ödemeli bakiyeden saatlik ücretlenir (`monthly_eur` / 672, en az 1 saat, ayda en çok `monthly_eur`) ve hesapta en az bir ön ödemeli tr-ist-1 sunucusu gerekir (409 `no_billing_server`). Kapatınca public IP bırakılır; port yönlendirme varken ya da public IP'si kapalı sunucular kullanırken reddedilir (409 `gateway_in_use`, `allow_no_egress: true` ile tekrarlayın). Dışarı çıkan bağlantılar 5651 sayılı Kanun gereği kaydedilir (2 yıl). Yalnız verilen alanlar değişir.

API: PUT /cloud/vpcs/{vpc_id}/gateway (operationId setVpcGateway, scope cloud_write)

```
gyro vpc gateway set <VPC-ID> [flags]
```

### Options

```
      --allow-no-egress          allow_no_egress
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --enabled                  enabled
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --inter-subnet string      inter_subnet (allow|deny)
      --nat                      nat
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro vpc gateway](gyro_vpc_gateway.md)	 - VPC kapısını getir

