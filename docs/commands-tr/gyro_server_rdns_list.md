## gyro server rdns list

Ters DNS kayıtlarını listele

### Synopsis

Ters DNS kayıtlarını listele

**tr-ist-1 (Gyro Bulut İstanbul):** her adres (birincil ve ek IPv4) için otomatik PTR `<a-b-c-d>.ip.gyrohosting.com` ve eşleşen A kaydı vardır. `hostname` ile kendi tam alan adınızı yazın (saniyeler içinde yayında); boş değer varsayılana döndürür. Gyro alan adları altı ve korunan marka adları reddedilir (`hostname_reserved`, `hostname_blocked`); sunucu başına saatte en çok 30 değişiklik (`rate_limited`). E-posta teslimi için adınızın A kaydı bu adresi göstermeli (zorunlu değil). Sunucu silinince ya da ek IP bırakılınca PTR kaldırılır.

API: GET /cloud/servers/{id}/rdns (operationId listRdns, scope cloud_read)

```
gyro server rdns list <ID> [flags]
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

* [gyro server rdns](gyro_server_rdns.md)	 - Ters DNS kayıtlarını listele

