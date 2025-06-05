package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
)

// 公钥
const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAimgQx3B2n6jQecToVqdm
AUpaLnOT+0DoRdiIRncOkSUrDWR5oz3VrCE8kmhw4VaBzQs3nH9fXt3cMLGQq7MN
lCYEx+pEbx2+etrRzPV7oW1hZdkd8SodS1LaFWGElTbPlhpqjZwFvFXa5rukLuqE
velxirLYurRAcbMSRsr/HJkuSviY7Nez9sJj8eHnlarJB3N2/vHSyK8upg3Frsgg
8hXuPUX3FFhxTvntqiCvF1mlAvtjKqf99dmbwYPgaFSUqBmT1SVuEATHDavQKwTQ
RLlQ8etYVgdYIdm9sXU6rYLuaHOi8yubnPfboUCDjF0bM/dD2OZVM5DTElJ+tb0E
JQIDAQAB
-----END PUBLIC KEY-----`

// 私钥
const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAimgQx3B2n6jQecToVqdmAUpaLnOT+0DoRdiIRncOkSUrDWR5
oz3VrCE8kmhw4VaBzQs3nH9fXt3cMLGQq7MNlCYEx+pEbx2+etrRzPV7oW1hZdkd
8SodS1LaFWGElTbPlhpqjZwFvFXa5rukLuqEvelxirLYurRAcbMSRsr/HJkuSviY
7Nez9sJj8eHnlarJB3N2/vHSyK8upg3Frsgg8hXuPUX3FFhxTvntqiCvF1mlAvtj
Kqf99dmbwYPgaFSUqBmT1SVuEATHDavQKwTQRLlQ8etYVgdYIdm9sXU6rYLuaHOi
8yubnPfboUCDjF0bM/dD2OZVM5DTElJ+tb0EJQIDAQABAoIBAANHE56leymsxcMc
dqQWssO1Dw2qjwaUMuv0hMWmbCHjz3Exv++ttzGsV9LfxkyPz+MbMOXz3j5Q4rYE
fniSdO/dNp/FXHmFKs2eAPnpllPsUzImQh9D8A2omKGI/arK8X+MgJl56GjdX6G3
L9wskZgtvxKZyagiPKR0qSokBIaY0NwLaCjCViAQMg4LUXlQDL7kyihGq6sfDjzm
OAqlsQVw2QJlBWRfv5BLQWhIVrQjanpWL/M5J0NLXGVxpqeke06jmZVZWS9c6Xrk
DTI62q0Mi1kCRF+VGfN7onAmyqGCnmNqJKfKSo6Y9mN+uFEalsON21bwZugNJAP7
LQ6xMMECgYEAkidTPpiDkMpsp8ZH2sREv5x6OORJVZcPO0VmqNClPK3K3BOuwROu
sY+uFqKbkLvCVmE58PRVLLx9pxWyUYNvZQxWA2+kXTDP7YtcaATIl/M2Dn5Cy8y2
5nBCDgtrrnmBJFrhhXQUWq6NVia/kuS/DpygHgzEGWHvLayCFJdGccsCgYEA8m4o
nkLDKrHj5q/+axB8ykB+WPrZ+5QlAT2bvbbsLvdy6go+1XFMunZBAKwuMWnBwyBT
iw07+WXK2iWK8aZbdln2/cQH59jsIh5tvjS+DMm5TBeuQlWbaFw3svubkBy1RRlc
ACWIWdtcgMjWFaHcF3RE2xwC6z/PKSvit6qg488CgYAInx1AAgswzWAffYjblNxh
WT+0f3L8A484dd2ac8Rrkez3komqhhi2DL2+RBVdMlZCdjXOyvsDSpM2+gas3E4D
7fhc94cAoFKDtsbas/2JSSWVMUmLcCYR7elXnsDZZ/2AQYQwNSrGNqnLvTlr1uaW
XBLt8t7WiHptfVGnGU1khQKBgQDGqI6QYZUcv3lJ6J5Ljc80zXS75ZQLHHLWeW3O
IL+aHRwFvS+QKK//fDf4HSAdTdSB2VOcdtPWoxt/d5RMQU5hxYconCL+UjvuGZFE
dGPohsVZHW6m3Q/kjOz3tkbb6SqB3S4O0LLTL3a18d29DE4WyihChMm/jSJtUpEO
gLwm8QKBgBzZNry0scU9W26YpPTIg+vo1iTFANUfwC6/4ZY55g1l/cvcMmc6eJrK
kIcTUs9djNgq6LjNLTNvDxJF7MqAAuD7EmNlvThtuEXXNpB6+N/79MkDBdw6DtCU
dE4GJygTioJRNt2OeeELyimcF4FywLFGLKxp407iV+s6WfcjF7st
-----END RSA PRIVATE KEY-----`

func parsePrivateKey(pemStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid PEM RSA private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func Parse(value []byte) (string, error) {
	priv, err := parsePrivateKey(privateKey)
	if err != nil {
		return "", err
	}
	cipherData, err := base64.StdEncoding.DecodeString(string(value))
	if err != nil {
		return "", err
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherData)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Unmarshal(body []byte, value any) error {
	str, err := Parse(body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(str), &value); err != nil {
		return err
	}
	return nil
}
