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
		Compare(hash string) bool
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

func Compare(v interface{}, hash string) bool {
	return New("", v).Hash() == hash
}

func (h Md5Hasher) Hash() string {
	val, _ := json.Marshal(fields(h.v))

	return fmt.Sprintf("%x", md5.Sum(val))
}

func (h Md5Hasher) Compare(hash string) bool {
	val, _ := json.Marshal(fields(h.v))

	return hash == fmt.Sprintf("%x", md5.Sum(val))
}

func (h Sha1Hasher) Hash() string {
	val, _ := json.Marshal(fields(h.v))

	return fmt.Sprintf("%x", sha1.Sum(val))
}

func (h Sha1Hasher) Compare(hash string) bool {
	val, _ := json.Marshal(fields(h.v))

	return hash == fmt.Sprintf("%x", sha1.Sum(val))
}

func (h Sha256Hasher) Hash() string {
	val, _ := json.Marshal(fields(h.v))

	return fmt.Sprintf("%x", sha256.Sum256(val))
}

func (h Sha256Hasher) Compare(hash string) bool {
	val, _ := json.Marshal(fields(h.v))

	return hash == fmt.Sprintf("%x", sha256.Sum256(val))
}

func (h Sha512Hasher) Hash() string {
	val, _ := json.Marshal(fields(h.v))

	return fmt.Sprintf("%x", sha512.Sum512(val))
}

func (h Sha512Hasher) Compare(hash string) bool {
	val, _ := json.Marshal(fields(h.v))

	return hash == fmt.Sprintf("%x", sha512.Sum512(val))
}

func fields(v interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	reflectTye := reflect.TypeOf(v)
	reflectValue := reflect.ValueOf(v)
	if reflectTye.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < reflectTye.NumField(); i++ {
		field := reflectTye.Field(i)
		tag := field.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		if field.Type.Kind() == reflect.Pointer {
			real := reflectValue.Field(i).Elem()
			if real.Kind() == reflect.Struct {
				result[tag] = fields(real.Interface())

				continue
			}
		}

		if field.Type.Kind() == reflect.Struct {
			result[tag] = fields(reflectValue.Field(i).Interface())

			continue
		}

		result[tag] = value(reflectValue.Field(i))
	}

	return result
}

func value(val reflect.Value) string {
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	switch val.Type().Kind() {
	case reflect.Bool:
		if val.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%d", val.Uint())
	case reflect.String:
		return val.String()
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", val.Float())
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%f", val.Complex())
	default:
		return ""
	}
}
