package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a/s/s/s", "/")
	want := []string{"a", "s", "s", "s"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}
