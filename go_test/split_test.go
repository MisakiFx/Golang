package split

//自带单元测试要求文件名_test结尾
//测试的函数前加Test，参数必带*testing.T
//测试文件不会进行编译
import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abcabc", "b")
	want := []string{"a", "ca", "c"}
	//切片不能直接比较，借助反射包中的方法进行比较
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v, got:%v\n", want, got)
	}
}
func TestSplit2(t *testing.T) {
	got := Split("abc", "b")
	want := []string{"a", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
func TestSplit3(t *testing.T) {
	got := Split("123ab356ab7", "ab")
	want := []string{"123", "356", "7"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
//测试组，把多个测试放在一个函数里
type test struct {
	str string
	sep string
	want []string
}
func TestSplit4(t *testing.T) {
	tests := []test{
		{"a.b.c.d", ".", []string{"a", "b", "c", "d"}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
	}
	for i, tst := range tests {
		got := Split(tst.str, tst.sep)
		if !reflect.DeepEqual(got, tst.want) {
			t.Errorf("%d error : want : %v, got : %v", i, tst.want, got)
		}
	}
}
func TestSplit5(t *testing.T) {
	tests := map[string]test {
		"case 1" : test{"a,b,c", ",", []string{"a", "b", "c"}},
		"case 2" : test{"a:b:c", ":", []string{"a", "b", "c"}},
		"case 3" : test{"a.b.c", ".", []string{"a", "b", "c"}},
	}
	t.Run("case 1", func(t *testing.T){
		for name, tst := range tests {
			got := Split(tst.str, tst.sep)
			if !reflect.DeepEqual(got, tst.want) {
				t.Errorf("%s error : want : %v, got : %v", name, tst.want, got)
			}
		}
	})
}
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
