## gyro server rdns set

Bir adresin PTR kaydını ayarla (veya sıfırla)

### Synopsis

Bir adresin PTR kaydını ayarla (veya sıfırla)

**tr-ist-1 (Gyro Bulut İstanbul):** her adres (birincil ve ek IPv4) için otomatik PTR `<a-b-c-d>.ip.gyrohosting.com` ve eşleşen A kaydı vardır. `hostname` ile kendi tam alan adınızı yazın (saniyeler içinde yayında); boş değer varsayılana döndürür. Gyro alan adları altı ve korunan marka adları reddedilir (HTTP 422, açıklayıcı mesaj); sunucu başına saatte en çok 30 değişiklik. E-posta teslimi için adınızın A kaydı bu adresi göstermeli (zorunlu değil). Sunucu silinince ya da ek IP bırakılınca PTR kaldırılır.

API: PUT /cloud/servers/{id}/rdns (operationId setRdns, scope cloud_write)

```
gyro server rdns set <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --hostname string          FQDN; empty string resets to the default
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --ip string                An address of this server (IPv6 addresses inside the server's /64 are allowed)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro server rdns](gyro_server_rdns.md)	 - Ters DNS kayıtlarını listele

