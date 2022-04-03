package keypair

import (
    "encoding/json"
    "encoding/base64"
    "github.com/danielhavir/go-xmss"
)

func Base64UrlEncode(src []byte) []byte {
    return []byte(base64.RawURLEncoding.EncodeToString(src))
}

func Base64UrlDecode(src []byte) ([]byte, error) {
    return base64.RawURLEncoding.DecodeString(string(src))
}

var params = xmss.SHA2_10_256


func Generate() (string, error) {
    k := make(map[string]string)
    k["kty"] = "PQK"
    k["alg"] = "xmss.SHA2_10_256"
    prv, pub := xmss.GenerateXMSSKeypair(params)
    k["x"] = string(Base64UrlEncode(*pub))
    k["d"] = string(Base64UrlEncode(*prv))
	jwk, err := json.Marshal(k)
    return string(jwk), err
}

func Sign(message []byte, jwk string) (string, error) {
    res := make(map[string]string)
    k := map[string]string{}
    json.Unmarshal([]byte(jwk), &k)
    d, err := Base64UrlDecode([]byte(k["d"]))
    if err != nil {
        // fmt.Println("failed to convert private key")
    }
    var prv xmss.PrivateXMSS = d
    var sig xmss.SignatureXMSS
    sig = *prv.Sign(params, message)
    res["signature"] = string(Base64UrlEncode([]byte(sig)))
	result, err := json.Marshal(res)
    return string(result), err
}

func Verify(message []byte, signature []byte, jwk string) (string, error) {
    res := make(map[string]bool)
    k := map[string]string{}
    json.Unmarshal([]byte(jwk), &k)
    x, err := Base64UrlDecode([]byte(k["x"]))
    if err != nil {
        // fmt.Println("failed to decode public key")
    }
    var pub xmss.PublicXMSS = x
    m := make([]byte, params.SignBytes()+len(message))
    v := xmss.Verify(params, m, signature, pub)
    res["verified"] = v
	result, err := json.Marshal(res)
    return string(result), err
}