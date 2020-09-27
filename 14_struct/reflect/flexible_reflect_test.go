package reflect

import (
	"errors"
	"reflect"
	"testing"
)

/*
	Employee， Customer, UpdateAge() 的定义在 reflect_test.go文件内
*/

func fillBySettings(st interface{}, settings map[string]interface{}) error {

	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.

	// 检查传递进来的st是不是一个指针
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type")
	}

	// 检查是不是一个struct对象
	// Elem() 获取指针指向的值，
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type")
	}

	// 检查setting是否为空
	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok { // 检测 setting中的key和struct中的field是否一致
			continue
		}
		if field.Type == reflect.TypeOf(v) { // 如果setting中的值的类型和struct的field的类型一致
			vstr := reflect.ValueOf(st)                 // 获取st的值（指针）
			vstr = vstr.Elem()                          // 获取st指针指向的值
			vstr.FieldByName(k).Set(reflect.ValueOf(v)) // 设置field的新值
		}

	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
