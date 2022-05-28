package recordhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	tagName = "hash"

	MD5    = "md5"
	SHA1   = "sha1"
	SHA256 = "sha256"
	SHA512 = "sha512"
)

type (
	Hasher interface {
		Hash() string
	}

	Md5Hasher struct {
		v interface{}
	}

	Sha1Hasher struct {
		v interface{}
	}

	Sha256Hasher struct {
		v interface{}
	}

	Sha512Hasher struct {
		v interface{}
	}
)

func New(function string, v interface{}) Hasher {
	switch function {
	case MD5:
		return &Md5Hasher{v: v}
	case SHA1:
		return &Sha1Hasher{v: v}
	case SHA256:
		return &Sha256Hasher{v: v}
	case SHA512:
		return &Sha512Hasher{v: v}
	default:
		return &Sha1Hasher{v: v}
	}
}

func Hash(v interface{}) string {
	return New("", v).Hash()
}

func (h Md5Hasher) Hash() string {
	val, _ := json.Marshal(fields(h.v))

	return fmt.Sprintf("%x", md5.Sum(val))
}

func (h Sha1Hasher) Hash() string {
	val, _ := json.Marshal(fields(h.v))

	return fmt.Sprintf("%x", sha1.Sum(val))
}

func (h Sha256Hasher) Hash() string {
	val, _ := json.Marshal(fields(h.v))

	return fmt.Sprintf("%x", sha256.Sum256(val))
}

func (h Sha512Hasher) Hash() string {
	val, _ := json.Marshal(fields(h.v))

	return fmt.Sprintf("%x", sha512.Sum512(val))
}

func fields(v interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	ref := reflect.TypeOf(v)
	if ref.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		tag := field.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		result[tag] = reflect.ValueOf(field).Interface()
	}

	return result
}
