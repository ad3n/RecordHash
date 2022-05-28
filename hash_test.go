package recordhash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	Simple struct {
		Field1 int    `hash:"field1"`
		Field2 string `hash:"field2"`
	}

	SimpleWithInterface struct {
		Field1 int         `hash:"field1"`
		Field2 bool        `hash:"field2"`
		Field3 interface{} `hash:"field3"`
	}

	Complex struct {
		Field1 uint   `hash:"field1"`
		Field2 string `hash:"field2"`
		Field3 Simple `hash:"field3"`
	}

	ComplexWithSkip struct {
		Field1 float32   `hash:"field1"`
		Field2 string    `hash:"field2"`
		Field3 *Simple   `hash:"field3"`
		Field4 string    `hash:"-"`
		Field5 complex64 `hash:"field5"`
	}

	ComplexWithSkipPrimitivePointer struct {
		Field1 *float32  `hash:"field1"`
		Field2 string    `hash:"field2"`
		Field3 *Simple   `hash:"field3"`
		Field4 string    `hash:"-"`
		Field5 complex64 `hash:"field5"`
	}
)

func Test_Md5_Hash(t *testing.T) {
	assert.NotEqual(t, New(MD5, "").Hash(), New(MD5, Simple{
		Field1: 1,
		Field2: "test",
	}).Hash())

	assert.NotEqual(t, New(MD5, "").Hash(), New(MD5, SimpleWithInterface{
		Field1: 1,
		Field2: true,
		Field3: map[string]string{
			"x": "y",
		},
	}).Hash())

	assert.NotEqual(t, New(MD5, "").Hash(), New(MD5, Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())

	assert.NotEqual(t, New(MD5, "").Hash(), New(MD5, ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())

	var float float32 = 1

	assert.NotEqual(t, New(MD5, "").Hash(), New(MD5, ComplexWithSkipPrimitivePointer{
		Field1: &float,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())
}

func Test_Sha1_Hash(t *testing.T) {
	assert.NotEqual(t, New(SHA1, "").Hash(), New(SHA1, Simple{
		Field1: 1,
		Field2: "test",
	}).Hash())

	assert.NotEqual(t, New(SHA1, "").Hash(), New(SHA1, SimpleWithInterface{
		Field1: 1,
		Field2: true,
		Field3: map[string]string{
			"x": "y",
		},
	}).Hash())

	assert.NotEqual(t, New(SHA1, "").Hash(), New(SHA1, Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())

	assert.NotEqual(t, New(SHA1, "").Hash(), New(SHA1, ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())
}

func Test_Sha256_Hash(t *testing.T) {
	assert.NotEqual(t, New(SHA256, "").Hash(), New(SHA256, Simple{
		Field1: 1,
		Field2: "test",
	}).Hash())

	assert.NotEqual(t, New(SHA256, "").Hash(), New(SHA256, SimpleWithInterface{
		Field1: 1,
		Field2: true,
		Field3: map[string]string{
			"x": "y",
		},
	}).Hash())

	assert.NotEqual(t, New(SHA256, "").Hash(), New(SHA256, Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())

	assert.NotEqual(t, New(SHA256, "").Hash(), New(SHA256, ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())
}

func Test_Sha512_Hash(t *testing.T) {
	assert.NotEqual(t, New(SHA512, "").Hash(), New(SHA512, Simple{
		Field1: 1,
		Field2: "test",
	}).Hash())

	assert.NotEqual(t, New(SHA512, "").Hash(), New(SHA512, SimpleWithInterface{
		Field1: 1,
		Field2: false,
		Field3: map[string]string{
			"x": "y",
		},
	}).Hash())

	assert.NotEqual(t, New(SHA512, "").Hash(), New(SHA512, Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())

	assert.NotEqual(t, New(SHA512, "").Hash(), New(SHA512, ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	}).Hash())
}

func Test_Alias(t *testing.T) {
	assert.NotEqual(t, Hash(""), Hash(Simple{
		Field1: 1,
		Field2: "test",
	}))

	assert.NotEqual(t, Hash(""), Hash(SimpleWithInterface{
		Field1: 1,
		Field2: true,
		Field3: map[string]string{
			"x": "y",
		},
	}))

	assert.NotEqual(t, Hash(""), Hash(Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	}))

	assert.NotEqual(t, Hash(""), Hash(ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	}))
}
