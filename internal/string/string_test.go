package string

import (
	"fmt"
	"strings"
	"testing"
)

////func TestStringTrim(t *testing.T) {
////	//s := `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`
////	//cs := `O3K5UmBvDSl3EjXyuQ5hk6s0u29pYXJP9yJic6wvijNPxFkPYc4NyMN9U1Th9P9kfy0v8LOyKGqCnS2ONvOYlkMnhSjlsWqnrPM`
////	s := "abcabcbhahahahaha"
////	cs := "abc"
////	s1 := strings.TrimPrefix(s, cs)
////	fmt.Println(s1)
////
////	s2 := myTrimLeft([]byte(s), cs)
////	fmt.Println(string(s2))
////	//
////	assert.Equal(t, s1, string(s2))
////}
//
////
////func TestMyTrimLeft(t *testing.T) {
////	s := "abcabchahahahahcdcabcabc"
////	s2 := myTrimLeft([]byte(s), "abc")
////	fmt.Println(string(s2))
////}
//
//func findLeftIndex(src []byte, cutset string) int {
//	st := 0
//	ps := 0
//
//	for st < len(src) {
//		if ps == len(cutset) {
//			ps = 0
//		}
//		if src[st] == cutset[ps] {
//			st++
//			ps++
//		} else {
//			return st
//		}
//	}
//	return -1
//}
//
//func myTrimLeft(src []byte, cutset string) []byte {
//	idx := findLeftIndex(src, cutset)
//	if idx == -1 {
//		return src
//	}
//	return src[idx:]
//}
//
//type testset struct {
//	s  string
//	cs string
//}
//type myTestset struct {
//	s  []byte
//	cs string
//}
//
//var (
//	testsets   []testset
//	myTestsets []myTestset
//)
//
//func init() {
//	//Init()
//}
//
//func Init() {
//	stime := time.Now()
//
//	N := 0
//	minLen := 100
//	maxLen := 10000
//	ss, err := randomStrings(N, minLen, maxLen)
//	if err != nil {
//		panic(err)
//	}
//
//	testsets = make([]testset, 0, N)
//	for _, s := range ss {
//		testsets = append(testsets, testset{
//			s:  strings.Clone(s),
//			cs: s[:minLen-1],
//		})
//	}
//
//	myTestsets = make([]myTestset, 0, N)
//	for _, s := range ss {
//		myTestsets = append(myTestsets, myTestset{
//			s:  []byte(s),
//			cs: s[:minLen-1],
//		})
//	}
//
//	fmt.Printf("start benchmark, prepare %v s \n", time.Since(stime).Seconds())
//
//}
//
//func BenchmarkMyTrimLeft(b *testing.B) {
//	fmt.Println(b.Name(), len(myTestsets))
//	b.ResetTimer()
//	b.N = 10000
//
//	for i := 0; i < b.N; i++ {
//		for _, tt := range myTestsets {
//			myTrimLeft(tt.s, tt.cs)
//		}
//	}
//}
//
//func BenchmarkTrimLeft(b *testing.B) {
//	fmt.Println(b.Name(), len(testsets))
//
//	b.ResetTimer()
//	b.N = 10000
//
//	for i := 0; i < b.N; i++ {
//		for _, tt := range testsets {
//			strings.TrimPrefix(tt.s, tt.cs)
//		}
//	}
//}
//
////	func TestTrimLeft(t *testing.T) {
////		for i := 0; i < len(testsets); i++ {
////			assert.Equal(t, testsets[i].s, string(myTestsets[i].s))
////			assert.Equal(t, testsets[i].cs, myTestsets[i].cs)
////
////			srcStr := strings.TrimPrefix(testsets[i].s, testsets[i].cs)
////			myStr := myTrimLeft(myTestsets[i].s, myTestsets[i].cs)
////			assert.Equal(t, srcStr, string(myStr), fmt.Sprintf("\nstr: %v \ncutset: %v \n", testsets[i].s, testsets[i].cs))
////		}
////	}
//const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//
//func randomString(length int) (string, error) {
//	b := make([]byte, length)
//	for i := range b {
//		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
//		if err != nil {
//			return "", err
//		}
//		b[i] = charset[num.Int64()]
//	}
//	return string(b), nil
//}
//
//func randomStrings(N, minLen, maxLen int) ([]string, error) {
//	result := make([]string, N)
//	for i := 0; i < N; i++ {
//		// 生成[minLen, maxLen]范围内的随机长度
//		lengthRange := maxLen - minLen + 1
//		length, err := rand.Int(rand.Reader, big.NewInt(int64(lengthRange)))
//		if err != nil {
//			return nil, err
//		}
//		length = length.Add(length, big.NewInt(int64(minLen)))
//
//		// 生成指定长度的随机字符串
//		randomStr, err := randomString(int(length.Int64()))
//		if err != nil {
//			return nil, err
//		}
//		result[i] = randomStr
//	}
//	return result, nil
//}

var (
	raw    = "abcdefg123"
	cutset = "123"
)

//func TestMyString(t *testing.T) {
//	str := New([]byte(raw))
//	cutSet := New([]byte(cutset))
//	str.Trim(cutset)
//	str.Cat(cutSet)
//
//	fmt.Println(string(str.buf))
//}

func Benchmark_MyString(b *testing.B) {
	str := New([]byte(raw))
	cutSet := New([]byte(cutset))
	f := func() {
		for i := 0; i < 10000; i++ {
			str.Trim(cutset)
			str.Cat(cutSet)
		}
	}
	fmt.Println("start benchmark")

	b.ResetTimer()
	b.N = 10000
	for i := 0; i < b.N; i++ {
		f()
	}
}

func Benchmark_RawString(b *testing.B) {
	str := strings.Clone(raw)
	f := func() {
		for i := 0; i < 10000; i++ {
			str = strings.TrimPrefix(str, cutset)
			str = cutset + str
		}
	}
	fmt.Println("start benchmark")

	b.ResetTimer()
	b.N = 10000
	for i := 0; i < b.N; i++ {
		f()
	}
}
