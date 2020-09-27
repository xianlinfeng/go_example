package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/* 每个对象都有type和value这两个属性，因此我们可以用反射来获取这两个属性 */

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}
func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
	CheckType(&f)
}

func CheckType(v interface{}) { //空接口表示任何类型都可以传入
	t := reflect.TypeOf(v)
	switch t.Kind() { // 用t.Kind()来检测t的类型。
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

/*
	struct对象的reflect
*/
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按名字获取成员

	// Reflect.Type和Reflect.Value都有FieldByName方法，注意他们的区别
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	// 访问struct的method
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}
