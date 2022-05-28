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
	var float float32 = 1
	var hasher Hasher = New(MD5, Simple{
		Field1: 1,
		Field2: "test",
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(MD5, SimpleWithInterface{
		Field1: 1,
		Field2: true,
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(MD5, Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(MD5, ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(MD5, ComplexWithSkipPrimitivePointer{
		Field1: &float,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))
}

func Test_Sha1_Hash(t *testing.T) {
	var float float32 = 1
	var hasher Hasher = New(SHA1, Simple{
		Field1: 1,
		Field2: "test",
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA1, SimpleWithInterface{
		Field1: 1,
		Field2: false,
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA1, Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA1, ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA1, ComplexWithSkipPrimitivePointer{
		Field1: &float,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))
}

func Test_Sha256_Hash(t *testing.T) {
	var float float32 = 1
	var hasher Hasher = New(SHA256, Simple{
		Field1: 1,
		Field2: "test",
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA256, SimpleWithInterface{
		Field1: 1,
		Field2: true,
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA256, Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA256, ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA256, ComplexWithSkipPrimitivePointer{
		Field1: &float,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))
}

func Test_Sha512_Hash(t *testing.T) {
	var float float32 = 1
	var hasher Hasher = New(SHA512, Simple{
		Field1: 1,
		Field2: "test",
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA512, SimpleWithInterface{
		Field1: 1,
		Field2: false,
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA512, Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA512, ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))

	hasher = New(SHA512, ComplexWithSkipPrimitivePointer{
		Field1: &float,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	})
	assert.True(t, hasher.Compare(hasher.Hash()))
}

func Test_Alias(t *testing.T) {
	var float float32 = 1
	v1 := Simple{
		Field1: 1,
		Field2: "test",
	}
	assert.True(t, Compare(v1, Hash(v1)))

	v2 := SimpleWithInterface{
		Field1: 1,
		Field2: true,
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	}
	assert.True(t, Compare(v2, Hash(v2)))

	v3 := Complex{
		Field1: 1,
		Field2: "test",
		Field3: Simple{
			Field1: 1,
			Field2: "test",
		},
	}
	assert.True(t, Compare(v3, Hash(v3)))

	v4 := ComplexWithSkip{
		Field1: 1,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	}
	assert.True(t, Compare(v4, Hash(v4)))

	v5 := ComplexWithSkipPrimitivePointer{
		Field1: &float,
		Field2: "test",
		Field3: &Simple{
			Field1: 1,
			Field2: "test",
		},
	}
	assert.True(t, Compare(v5, Hash(v5)))
	assert.True(t, Compare("", Hash("")))
}
