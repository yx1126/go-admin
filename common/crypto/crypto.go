package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 公钥
const _ = `
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAKfSt+NNYfJ0Pb54LyYaoBkU8aRopv9T
s1UASMq05ztPXxLNskUWpc4g+GoGQDEiT14uNXNXkYB91GtlPMMlnCkCAwEAAQ==
-----END PUBLIC KEY-----
`

// 私钥
const privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIBOQIBAAJBAKfSt+NNYfJ0Pb54LyYaoBkU8aRopv9Ts1UASMq05ztPXxLNskUW
pc4g+GoGQDEiT14uNXNXkYB91GtlPMMlnCkCAwEAAQJAAbgXiM01daWvLuZNmqxR
0NqgEbkTYjDZ+MLvUUmXx3RnwqbUK8FFo5gR6WWXP278cLo7kc3xYrTF0Tv3OSi7
SQIhAOFj5C0Abb4dWH3xLDLitYQXF9UGTE47CHuiFfzp24TPAiEAvp1miZ8e1xcX
OJP9JI7oZK6NfifT5bmTp1Z8Emw+/YcCIB2zSBIWCGARBeQyr5xU+45YbK+JkOyO
IuQHy9X0CxVJAiBw2OCOkyx/7ESsEzjnvs8oZAqSaGPTefVHOTVE9t6n2wIgcHL+
LKU9DB7uRpFZHp3c3A86q+VfdukJFleMqWotFIs=
-----END RSA PRIVATE KEY-----
`

func parsePrivateKey(pemStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid PEM RSA private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func Parse(value []byte) ([]byte, error) {
	priv, err := parsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, value)
	if err != nil {
		return nil, err
	}
	return data, nil
}
