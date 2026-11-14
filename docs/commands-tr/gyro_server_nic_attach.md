## gyro server nic attach

Sunucuyu subnete bağla

### Synopsis

Sunucuyu subnete bağla

tr-ist-1 sunucusunun özel ağ kartları. Takma sıcak yapılır (yeniden başlatma yok): açık sunucuda yaklaşık 10-60 sn'de hazır (`state` attaching -> ready); kapalı sunucuda ayarlar açılışta uygulanır. Bir sunucu aynı anda tek VPC'ye bağlanır (409 `other_vpc`), en çok 3 kart. `ip` boş = otomatik.

API: POST /cloud/servers/{id}/nics (operationId attachServerNic, scope cloud_write)

```
gyro server nic attach <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for attach
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --ip string                ip
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --subnet-id string         subnet_id
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
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

* [gyro server nic](gyro_server_nic.md)	 - Sunucunun özel ağ kartları

