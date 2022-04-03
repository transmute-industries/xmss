package main

import (
	"syscall/js"
	"encoding/json"
	"encoding/base64"
	"example.com/hello/src/keypair"
)


func Base64UrlEncode(src []byte) []byte {
    return []byte(base64.RawURLEncoding.EncodeToString(src))
}

func Base64UrlDecode(src []byte) ([]byte, error) {
    return base64.RawURLEncoding.DecodeString(string(src))
}

func handleRequest(this js.Value, args []js.Value) interface{} {
	
	requestString := args[0].String()
	requestMap := map[string]string{}
    json.Unmarshal([]byte(requestString), &requestMap)

	if (requestMap["command"] == "generate"){
		res, err := keypair.Generate()
		if err != nil {
			return js.ValueOf(nil)
		}
		return js.ValueOf(res)
	}

	if (requestMap["command"] == "sign"){
		message, err := Base64UrlDecode([]byte(requestMap["message"]))
		if err != nil {
			return js.ValueOf(nil)
		}
		publicKeyJwk := requestMap["jwk"]
		res, err := keypair.Sign(message, publicKeyJwk)
		if err != nil {
			return js.ValueOf(nil)
		}
		return js.ValueOf(res)
	}

	if (requestMap["command"] == "verify"){
		message, err := Base64UrlDecode([]byte(requestMap["message"]))
		if err != nil {
			return js.ValueOf(nil)
		}
		signature, err := Base64UrlDecode([]byte(requestMap["signature"]))
		if err != nil {
			return js.ValueOf(nil)
		}
		publicKeyJwk := requestMap["jwk"]
		res, err := keypair.Verify(message, signature, publicKeyJwk)
		if err != nil {
			return js.ValueOf(nil)
		}
		return js.ValueOf(res)
	}

	return js.ValueOf(nil)
}

func init() {
	js.Global().Set("handleRequest", js.FuncOf(handleRequest))
}

func main() {
    <-make(chan bool)
}


