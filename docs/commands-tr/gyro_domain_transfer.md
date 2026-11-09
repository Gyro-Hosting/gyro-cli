## gyro domain transfer

Alan Adı Transferi

### Synopsis

Alan Adı Transferi

**Para harcar / fatura oluşturur.** Para kuralları (AK-40): gövdede `"confirm": true`, `Idempotency-Key` başlığı, yalnız bakiye (kayıtlı kart yok), hesap başına günlük ücretli işlem sınırı (varsayılan 50, hesap/bayi bazında değiştirilebilir, 00:00 TR'de sıfırlanır). Başka bir sağlayıcıdaki alan adını buraya taşır. — kod gerekebilir

API: POST /domains/transfer (operationId coreTransferDomain, scope Domains/TransferDomain)

```
gyro domain transfer [flags]
```

### Options

```
      --accept-premium string    accept_premium
      --addons string            Birlikte alınacak ek hizmetler. Alan adı yönetimi, kayıt gizliliği ve yönlendirme; yalnız uzantı sunuyorsa.
      --auth-code string         Mevcut sağlayıcıdan alınan transfer kodu. Uzantı işaretliyse zorunludur ve her durumda hizmete yazılır.
      --billing-profile-id int   Fatura profilinin numarası.
      --card-id string           card_id
      --contacts string          Rol bazında kişiler. Her rol ya kayıtlı bir profil numarası ya da satır içi bir kişi taşır ve o rol için profili ezer. (JSON)
      --coupons string           Kupon kodları.
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --domain string            Tam alan adı.
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for transfer
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --method string            method
      --nameservers string       Ad sunucuları. İkiden dörde; verilmezse hesabın varsayılanları uygulanır.
      --notes string             Sipariş notu.
      --payment string           Ödeme kaynağı. Hesap bakiyesi ya da kayıtlı kart. (JSON)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --terms                    Şartların kabulü. Doğru gönderilmelidir.
      --whois-profile-id int     Kayıtlı kayıt sahibi profili. Dört rolü birden doldurur.
      --years string             years
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

* [gyro domain](gyro_domain.md)	 - Alan adları: sorgu, kayıt, transfer, ad sunucuları, DNS, WHOIS

