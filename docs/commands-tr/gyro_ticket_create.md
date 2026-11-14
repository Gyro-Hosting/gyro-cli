## gyro ticket create

Talep Açma

### Synopsis

Talep Açma

Yeni bir destek talebi açar. — ek kabul eder **Sınır artırma**: hesap sınırları (GET /cloud/account/limits) talep ile artırılır: department_id 3, konu `Limit increase: <limit_key>`, mesajda istenen sayı ve ne yaptığınız. CLI: `gyro ticket create --department-id 3 --subject ... --message ...`.

API: POST /tickets (operationId coreCreateTicket, scope Tickets/CreateTicket)

```
gyro ticket create [flags]
```

### Options

```
      --attachments strings      Ekler. Kodlanmış içerik olarak gönderilir; adres reddedilir.
      --credentials string       Erişim bilgileri. Her satır bir grup numarası ve alan haritası taşır; mesaja şifreli işlenir.
      --custom-fields string     Departman alanlarının yanıtları. Alan numarasını değere eşler; kutu tipli alanlar dizi alır. (JSON)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --department-id int        Departmanın numarası.
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --encrypt-message          Mesaj gövdesi şifreli saklansın mı.
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --message string           Açılış mesajı. En az beş karakter; düz metindir ve işaretleme ayıklanır.
      --priority int             Öncelik. Verilmezse orta alınır.
      --service-id int           Talebi bir hizmete bağlar.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --subject string           Konu.
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

* [gyro ticket](gyro_ticket.md)	 - Destek talepleri

