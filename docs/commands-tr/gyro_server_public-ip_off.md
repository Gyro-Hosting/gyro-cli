## gyro server public-ip off

Public IPv4'ü kapat

### Synopsis

Public IPv4'ü kapat

Public IPv4 kapanınca sunucuya yalnız VPC'den, kapı port yönlendirmesinden ya da web konsoldan (her zaman çalışır) erişilir. Onay için sunucunun mevcut IP'si (`confirm_ip`) yazılır. Hazır bir özel kart (409 `no_private_network`) gerekir ve ek IPv4 olmamalı (409 `extra_ips`). VPC kapısı (NAT) kapalıysa sunucu internete çıkamaz: `allow_no_egress: true` verilmezse 409 `no_egress`. Adres sunucuya ayrılmış kalır; açınca aynı IP gelir. Eski misafir ajanı 409 `agent_update_required` döner (bkz. GET /cloud/servers/{id}/agent).

API: DELETE /cloud/servers/{id}/public-ip (operationId disableServerPublicIp, scope cloud_write)

```
gyro server public-ip off <ID> [flags]
```

### Options

```
      --allow-no-egress          allow_no_egress
      --confirm-ip string        confirm_ip
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for off
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
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

* [gyro server public-ip](gyro_server_public-ip.md)	 - Public IPv4 durumu

