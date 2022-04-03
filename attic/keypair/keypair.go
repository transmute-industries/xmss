package keypair

import (
    "fmt"
    "encoding/json"
    "encoding/base64"
    "github.com/danielhavir/go-xmss"
)

var params = xmss.SHA2_10_256

func Base64UrlEncode(src []byte) []byte {
    return []byte(base64.RawURLEncoding.EncodeToString(src))
}

func Base64UrlDecode(src []byte) ([]byte, error) {
    return base64.RawURLEncoding.DecodeString(string(src))
}

func Generate() (string, error) {
    jwk := make(map[string]string)
    prv, pub := xmss.GenerateXMSSKeypair(params)
    jwk["kty"] = "PQK"
    jwk["alg"] = "xmss.SHA2_10_256"
    jwk["d"] = string(Base64UrlEncode(*prv))
    jwk["x"] = string(Base64UrlEncode(*pub))
    sjwk, err := json.Marshal(jwk)
    return string(sjwk), err
}

func Sign(kp string, msg []byte) (string, error) {
    a := map[string]string{}
    json.Unmarshal([]byte(kp), &a)
    d, err := Base64UrlDecode([]byte(a["d"]))
    if err != nil {
        fmt.Println("failed to convert private key")
    }
    var prv xmss.PrivateXMSS = d
    var sig xmss.SignatureXMSS
    sig = *prv.Sign(params, msg)
    return string(Base64UrlEncode([]byte(sig))), nil
}

func Verify( message []byte, signature string, jwk string) (bool, error){
    
    jwkm := map[string]string{}
    json.Unmarshal([]byte(jwk), &jwkm)
    x, err := Base64UrlDecode([]byte(jwkm["x"]))
    if err != nil {
        fmt.Println("failed to decode public key")
    }
    var pub xmss.PublicXMSS = x
    m := make([]byte, params.SignBytes()+len(message))
    sig, err2 := Base64UrlDecode([]byte(signature))
    if err2 != nil {
        fmt.Println("failed to decode signature")
    }
    v := xmss.Verify(params, m, sig, pub)
    return v, nil
}

