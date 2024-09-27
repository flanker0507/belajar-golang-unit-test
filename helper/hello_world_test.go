package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	Benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Yuda",
			request: "YUDA WIRA PRATAMA",
		},

		{
			name:    "Dilla",
			request: "ADE MARIAM FADILLA MASRI",
		},
	}

	for _, benchmark := range Benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				helloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Yuda", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helloWorld("Yuda")
		}
	})
	b.Run("Dilla", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helloWorld("Dilla")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld("Dilla")
	}
}

func BenchmarkHelloWorldYuda(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld("Yuda Wira Pratama")
	}
}

func TestTableHelloWOrld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Yuda",
			request:  "Yuda",
			expected: "Hello Yuda",
		},
		{
			name:     "Dilla",
			request:  "Dilla",
			expected: "Hello Dilla",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := helloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Yuda", func(t *testing.T) {
		result := helloWorld("Yuda")
		assert.Equal(t, "Hello Yuda", result, "Result must be 'Hello Yuda'")
	})

	t.Run("Dilla", func(t *testing.T) {
		result := helloWorld("Dilla")
		assert.Equal(t, "Hello Dilla", result, "Result must be 'Hello Dilla'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After unit Test")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "Darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := helloWorld("Yuda")
	assert.Equal(t, "Hello Yuda", result, "Result must be 'Hello Yuda'")
}

func TestHelloWorldYudaAssetion(t *testing.T) {
	result := helloWorld("Yuda")
	assert.Equal(t, "Hello Yuda", result, "Result must be 'Hello Yuda'")
	fmt.Println("Test Error Yuda Done")
}

func TestHelloWorldYuda(t *testing.T) {
	result := helloWorld("Yuda")

	if result != "Hello Yuda" {
		//error

		t.Error("Result Must Be 'Hello Yuda'")
	}

	fmt.Println("Test Error Yuda Done")

}

func TestHelloWorldDillaRequire(t *testing.T) {
	result := helloWorld("Dilla")
	require.Equal(t, "Hello Dilla", result, "Result Must Be 'Hello Dilla'")

	fmt.Println("Test Error Dilla Done")

}

func TestHelloWorldDilla(t *testing.T) {
	result := helloWorld("Dilla")

	if result != "Hello Dilla" {
		t.Fatal("Result Must Be 'Hello Dilla'")
	}

	fmt.Println("Test Error Dilla Done")

}
