package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Chitanda")
	assert.Equal(t, "Hello Chitanda", result, "Result must be 'Hello Chitanda'") // ini jika makai testify (yang assert)
	fmt.Println("TestHelloWorld with Assert Done")                               // assert ini manggil fail() maka ini akan tetap tereksekusi
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Chitanda")
	require.Equal(t, "Hello Chitanda", result, "Result must be 'Hello Chitanda'") // jika gagal maka failnow()
	fmt.Println("TestHelloWorld with Require Done")                               // require = failnow, maka ini tidak akan dieksekusi jika gagal
}

func TestHelloWorldEru(t *testing.T) {
	result := HelloWorld("Eru")

	if result != "Hello Eru" {
		t.Fail() // jika pakai ini maka walaupun fail akan dilanjutkan ke kode program berikutnya
	}

	fmt.Println("TestHelloWorldEru Done")
}

func TestHelloWorldChitanda(t *testing.T) {
	result := HelloWorld("Chitanda")

	if result != "Hello Chitanda" {
		t.FailNow() // maka apabila fail tidak akan dilanjutkan ke kode program berikutnya
	}

	fmt.Println("TestHelloWorldChitanda Done") // ini tidak akan terpanggil apabila failnow()
}

func TestHelloWorldTendou(t *testing.T) {
	result := "Hello Tendou"

	if result != "Hello Tendou" {
		t.Error("Harusnya Hello Tendou") // memberi tahu pesan errornya, selanjutnya seperti t.Fail() dan mengeksekusi kode selanjutnya
	}

	fmt.Println("Dieksekusi")
}

func TestHelloWorldArisu(t *testing.T) {
	result := "Hello Arisu"

	if result != "Hello Arisu" {
		t.Fatal("Harusnya Hello Arisu") // memberi tahu pesan errornya, selanjutnya seperti t.FailNow() tidak mengeksekusi program selanjutnya
	}

	fmt.Println("Tidak Dieksekusi") // jika gagal tidak dieksekusi karena fatal() = failnow()
}
