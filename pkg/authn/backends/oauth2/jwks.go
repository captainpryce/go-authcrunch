// Copyright 2022 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oauth2

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"github.com/greenpau/go-authcrunch/pkg/errors"
	"math/big"
	"strings"
)

// JwksKey is a JSON object that represents a cryptographic key.
// See https://tools.ietf.org/html/rfc7517#section-4,
// https://tools.ietf.org/html/rfc7518#section-6.3
type JwksKey struct {
	Algorithm    string `json:"alg,omitempty" xml:"alg,omitempty" yaml:"alg,omitempty"`
	Exponent     string `json:"e,omitempty" xml:"e,omitempty" yaml:"e,omitempty"`
	KeyID        string `json:"kid,omitempty" xml:"kid,omitempty" yaml:"kid,omitempty"`
	KeyType      string `json:"kty,omitempty" xml:"kty,omitempty" yaml:"kty,omitempty"`
	Modulus      string `json:"n,omitempty" xml:"n,omitempty" yaml:"n,omitempty"`
	PublicKeyUse string `json:"use,omitempty" xml:"use,omitempty" yaml:"use,omitempty"`
	NotBefore    string `json:"nbf,omitempty" xml:"nbf,omitempty" yaml:"nbf,omitempty"`

	Curve  string `json:"crv,omitempty" xml:"crv,omitempty" yaml:"crv,omitempty"`
	CoordX string `json:"x,omitempty" xml:"x,omitempty" yaml:"x,omitempty"`
	CoordY string `json:"y,omitempty" xml:"y,omitempty" yaml:"y,omitempty"`

	publicKey *rsa.PublicKey
}

// Validate returns error if JwksKey does not contain relevant information.
func (k *JwksKey) Validate() error {
	if k.KeyID == "" {
		return errors.ErrJwksKeyIDEmpty
	}

	switch k.KeyType {
	case "RSA":
		switch k.Algorithm {
		case "RS256", "RS384", "RS512", "":
		default:
			return errors.ErrJwksKeyAlgoUnsupported.WithArgs(k.Algorithm, k.KeyID)
		}
	case "EC":
		switch k.Curve {
		case "P-256", "P-521":
		case "":
			return errors.ErrJwksKeyCurveEmpty.WithArgs(k.KeyID)
		default:
			return errors.ErrJwksKeyCurveUnsupported.WithArgs(k.Curve, k.KeyID)
		}
	case "":
		return errors.ErrJwksKeyTypeEmpty.WithArgs(k.KeyID)
	default:
		return errors.ErrJwksKeyTypeUnsupported.WithArgs(k.KeyType, k.KeyID)
	}

	switch k.PublicKeyUse {
	case "sig", "":
	default:
		return errors.ErrJwksKeyUsageUnsupported.WithArgs(k.PublicKeyUse, k.KeyID)
	}

	switch k.KeyType {
	case "RSA":
		if k.Exponent == "" {
			return errors.ErrJwksKeyExponentEmpty.WithArgs(k.KeyID)
		}

		if k.Modulus == "" {
			return errors.ErrJwksKeyModulusEmpty.WithArgs(k.KeyID)
		}

		// Add padding
		if i := len(k.Modulus) % 4; i != 0 {
			k.Modulus += strings.Repeat("=", 4-i)
		}

		var mod []byte
		var err error
		if strings.ContainsAny(k.Modulus, "/+") {
			// This decoding works with + and / signs. (legacy)
			mod, err = base64.StdEncoding.DecodeString(k.Modulus)
		} else {
			// This decoding works with - and _ signs.
			mod, err = base64.URLEncoding.DecodeString(k.Modulus)
		}

		if err != nil {
			return errors.ErrJwksKeyDecodeModulus.WithArgs(k.KeyID, k.Modulus, err)
		}
		n := big.NewInt(0)
		n.SetBytes(mod)

		exp, err := base64.StdEncoding.DecodeString(k.Exponent)
		if err != nil {
			return errors.ErrJwksKeyDecodeExponent.WithArgs(k.KeyID, err)
		}
		// The "e" (exponent) parameter contains the exponent value for the RSA
		// public key.  It is represented as a Base64urlUInt-encoded value.
		//
		// For instance, when representing the value 65537, the octet sequence
		// to be base64url-encoded MUST consist of the three octets [1, 0, 1];
		// the resulting representation for this value is "AQAB".
		var eb []byte
		if len(exp) < 8 {
			eb = make([]byte, 8-len(exp), 8)
			eb = append(eb, exp...)
		} else {
			eb = exp
		}
		er := bytes.NewReader(eb)
		var e uint64
		if err := binary.Read(er, binary.BigEndian, &e); err != nil {
			return errors.ErrJwksKeyConvExponent.WithArgs(k.KeyID, err)
		}
		k.publicKey = &rsa.PublicKey{N: n, E: int(e)}

	default:
		return errors.ErrJwksKeyTypeNotImplemented.WithArgs(k.KeyID, k.KeyType)
	}

	return nil
}

// GetPublicKey returns pointer to rsa.PublicKey.
func (k *JwksKey) GetPublicKey() *rsa.PublicKey {
	return k.publicKey
}
