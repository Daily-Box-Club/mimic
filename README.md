# Example

```go
import (
    http "github.com/saucesteals/fhttp"
    "github.com/saucesteals/mimic"
)

...

client := &http.Client{
    Transport: mimic.Chrome101.NewTransport(),
}

req, err := http.NewRequest("GET", "https://tls.peet.ws/api/all", nil)

req.Header = http.Header{
    "cache-control":             {"max-age=0"},
    "upgrade-insecure-requests": {"1"},
    "user-agent":                {mimic.Chrome101.UserAgent},
    "accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
    "sec-gpc":                   {"1"},
    "sec-fetch-site":            {"none"},
    "sec-fetch-mode":            {"navigate"},
    "sec-fetch-user":            {"?1"},
    "sec-fetch-dest":            {"document"},
    "accept-encoding":           {"gzip, deflate, br"},
    "accept-language":           {"en-US,en;q=0.9"},
    http.PHeaderOrderKey: mimic.Chrome101.H2Options.PseudoHeaderOrder,
    http.HeaderOrderKey: {
        "cache-control", "upgrade-insecure-requests",
        "user-agent", "accept", "sec-gpc", "sec-fetch-site",
        "sec-fetch-mode", "sec-fetch-user", "sec-fetch-dest",
        "accept-encoding", "accept-language",
    },
}

res, err := client.Do(req)
```

# Result

```json
{
  "http_version": "h2",
  "path": "/api/all",
  "method": "GET",
  "tls": {
    "ciphers": [
      "TLS_GREASE (0xCACA)",
      "TLS_AES_128_GCM_SHA256",
      "TLS_AES_256_GCM_SHA384",
      "TLS_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
      "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
      "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
      "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
      "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
      "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
      "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
      "TLS_RSA_WITH_AES_128_GCM_SHA256",
      "TLS_RSA_WITH_AES_256_GCM_SHA384",
      "TLS_RSA_WITH_AES_128_CBC_SHA",
      "TLS_RSA_WITH_AES_256_CBC_SHA"
    ],
    "curves": [
      "TLS_GREASE (0x5A5A)",
      "X25519 (29)",
      "P-256 (23)",
      "P-384 (24)"
    ],
    "extensions": [
      "TLS_GREASE (0x2A2A)",
      "server_name (0)",
      "extended_master_secret (23)",
      "extensionRenegotiationInfo (boringssl) (65281)",
      "supported_groups (10)",
      "ec_point_formats (11)",
      "session_ticket (35)",
      "application_layer_protocol_negotiation (16)",
      "status_request (5)",
      "signature_algorithms (13)",
      "signed_certificate_timestamp (18)",
      "key_share (51)",
      "psk_key_exchange_modes (45)",
      "supported_versions (43)",
      "compress_certificate (27)",
      "extensionApplicationSettings (boringssl) (17513)",
      "TLS_GREASE (0xEAEA)",
      "padding (21)"
    ],
    "points": ["0"],
    "version": "771",
    "protocols": ["h2", "http/1.1"],
    "versions": ["771", "770", "769", "768"],
    "ja3": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-17513,29-23-24,0",
    "ja3_hash": "e1d8b04eeb8ef3954ec4f49267a783ef"
  },
  "http2": {
    "akamai_fingerprint": "1:65536,3:1000,4:6291456,6:262144|15663105|0|m,a,s,p",
    "akamai_fingerprint_hash": "7ad845f20fc17cc8088a0d9312b17da1",
    "sent_frames": [
      {
        "frame_type": "SETTINGS",
        "length": 24,
        "settings": [
          "HEADER_TABLE_SIZE = 65536",
          "MAX_CONCURRENT_STREAMS = 1000",
          "INITIAL_WINDOW_SIZE = 6291456",
          "MAX_HEADER_LIST_SIZE = 262144"
        ]
      },
      {
        "frame_type": "WINDOW_UPDATE",
        "length": 4,
        "increment": 15663105
      },
      {
        "frame_type": "HEADERS",
        "stream_id": 1,
        "length": 355,
        "headers": [
          ":method: GET",
          ":authority: tls.peet.ws",
          ":scheme: https",
          ":path: /api/all",
          "upgrade-insecure-requests: 1",
          "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36",
          "accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
          "sec-gpc: 1",
          "dnt: 1",
          "sec-fetch-site: none",
          "sec-fetch-mode: navigate",
          "sec-fetch-user: ?1",
          "sec-fetch-dest: document",
          "accept-encoding: gzip, deflate, br",
          "accept-language: en-US,en;q=0.9"
        ]
      }
    ]
  }
}
```
