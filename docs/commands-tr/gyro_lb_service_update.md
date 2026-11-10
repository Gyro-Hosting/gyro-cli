## gyro lb service update

Servisi güncelle

### Synopsis

Servisi güncelle

{port} servisinin seçeneklerini değiştirir. Protokol değişmez (silip ekleyin).

API: PATCH /cloud/load-balancers/{id}/services/{port} (operationId typed_patch_load_balancers__id__services__port, scope services_write)

```
gyro lb service update <ID> <PORT> [flags]
```

### Options

```
      --certificates string      HTTPS: comma separated certificate ids of this load balancer
      --cookie-lifetime int      cookie_lifetime
      --cookie-name string       cookie_name
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --destination-port int     Port on the targets
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --hc-domain string         HTTP check Host header
      --hc-interval int          3-60 s
      --hc-path string           HTTP check path (starts with /)
      --hc-port int              hc_port
      --hc-protocol string       Health check type (tcp|http)
      --hc-retries int           0-5
      --hc-status-codes string   Comma separated, e.g. 2??,3??
      --hc-timeout int           1-30 s, below interval
      --hc-tls                   HTTP check over TLS
  -h, --help                     help for update
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --idle-timeout int         Idle timeout (s)
      --proxyprotocol            Send PROXY protocol header to targets
      --redirect-http            HTTPS: redirect port 80 to HTTPS (port 80 must be free)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --sticky                   Sticky sessions (cookie)
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

* [gyro lb service](gyro_lb_service.md)	 - Yük dengeleyicinin servislerini (portlarını) listele

