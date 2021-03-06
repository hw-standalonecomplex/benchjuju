package to

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	v := ""
	if String(&v) != v {
		t.Errorf("to: String failed to return the correct string -- expected %v, received %v",
			v, String(&v))
	}
}

func TestStringHandlesNil(t *testing.T) {
	if String(nil) != "" {
		t.Errorf("to: String failed to correctly convert nil -- expected %v, received %v",
			"", String(nil))
	}
}

func TestStringPtr(t *testing.T) {
	v := ""
	if *StringPtr(v) != v {
		t.Errorf("to: StringPtr failed to return the correct string -- expected %v, received %v",
			v, *StringPtr(v))
	}
}

func TestStringSlice(t *testing.T) {
	v := []string{}
	if out := StringSlice(&v); !reflect.DeepEqual(out, v) {
		t.Errorf("to: StringSlice failed to return the correct slice -- expected %v, received %v",
			v, out)
	}
}

func TestStringSliceHandlesNil(t *testing.T) {
	if out := StringSlice(nil); out != nil {
		t.Errorf("to: StringSlice failed to correctly convert nil -- expected %v, received %v",
			nil, out)
	}
}

func TestStringSlicePtr(t *testing.T) {
	v := []string{"a", "b"}
	if out := StringSlicePtr(v); !reflect.DeepEqual(*out, v) {
		t.Errorf("to: StringSlicePtr failed to return the correct slice -- expected %v, received %v",
			v, *out)
	}
}

func TestStringMap(t *testing.T) {
	msp := map[string]*string{"foo": StringPtr("foo"), "bar": StringPtr("bar"), "baz": StringPtr("baz")}
	for k, v := range StringMap(msp) {
		if *msp[k] != v {
			t.Errorf("to: StringMap incorrectly converted an entry -- expected [%s]%v, received[%s]%v",
				k, v, k, *msp[k])
		}
	}
}

func TestStringMapHandlesNil(t *testing.T) {
	msp := map[string]*string{"foo": StringPtr("foo"), "bar": nil, "baz": StringPtr("baz")}
	for k, v := range StringMap(msp) {
		if msp[k] == nil && v != "" {
			t.Errorf("to: StringMap incorrectly converted a nil entry -- expected [%s]%v, received[%s]%v",
				k, v, k, *msp[k])
		}
	}
}

func TestStringMapPtr(t *testing.T) {
	ms := map[string]string{"foo": "foo", "bar": "bar", "baz": "baz"}
	for k, msp := range *StringMapPtr(ms) {
		if ms[k] != *msp {
			t.Errorf("to: StringMapPtr incorrectly converted an entry -- expected [%s]%v, received[%s]%v",
				k, ms[k], k, *msp)
		}
	}
}

func TestBool(t *testing.T) {
	v := false
	if Bool(&v) != v {
		t.Errorf("to: Bool failed to return the correct string -- expected %v, received %v",
			v, Bool(&v))
	}
}

func TestBoolHandlesNil(t *testing.T) {
	if Bool(nil) != false {
		t.Errorf("to: Bool failed to correctly convert nil -- expected %v, received %v",
			false, Bool(nil))
	}
}

func TestBoolPtr(t *testing.T) {
	v := false
	if *BoolPtr(v) != v {
		t.Errorf("to: BoolPtr failed to return the correct string -- expected %v, received %v",
			v, *BoolPtr(v))
	}
}

func TestInt(t *testing.T) {
	v := 0
	if Int(&v) != v {
		t.Errorf("to: Int failed to return the correct string -- expected %v, received %v",
			v, Int(&v))
	}
}

func TestIntHandlesNil(t *testing.T) {
	if Int(nil) != 0 {
		t.Errorf("to: Int failed to correctly convert nil -- expected %v, received %v",
			0, Int(nil))
	}
}

func TestIntPtr(t *testing.T) {
	v := 0
	if *IntPtr(v) != v {
		t.Errorf("to: IntPtr failed to return the correct string -- expected %v, received %v",
			v, *IntPtr(v))
	}
}

func TestInt32(t *testing.T) {
	v := int32(0)
	if Int32(&v) != v {
		t.Errorf("to: Int32 failed to return the correct string -- expected %v, received %v",
			v, Int32(&v))
	}
}

func TestInt32HandlesNil(t *testing.T) {
	if Int32(nil) != int32(0) {
		t.Errorf("to: Int32 failed to correctly convert nil -- expected %v, received %v",
			0, Int32(nil))
	}
}

func TestInt32Ptr(t *testing.T) {
	v := int32(0)
	if *Int32Ptr(v) != v {
		t.Errorf("to: Int32Ptr failed to return the correct string -- expected %v, received %v",
			v, *Int32Ptr(v))
	}
}

func TestInt64(t *testing.T) {
	v := int64(0)
	if Int64(&v) != v {
		t.Errorf("to: Int64 failed to return the correct string -- expected %v, received %v",
			v, Int64(&v))
	}
}

func TestInt64HandlesNil(t *testing.T) {
	if Int64(nil) != int64(0) {
		t.Errorf("to: Int64 failed to correctly convert nil -- expected %v, received %v",
			0, Int64(nil))
	}
}

func TestInt64Ptr(t *testing.T) {
	v := int64(0)
	if *Int64Ptr(v) != v {
		t.Errorf("to: Int64Ptr failed to return the correct string -- expected %v, received %v",
			v, *Int64Ptr(v))
	}
}

func TestFloat32(t *testing.T) {
	v := float32(0)
	if Float32(&v) != v {
		t.Errorf("to: Float32 failed to return the correct string -- expected %v, received %v",
			v, Float32(&v))
	}
}

func TestFloat32HandlesNil(t *testing.T) {
	if Float32(nil) != float32(0) {
		t.Errorf("to: Float32 failed to correctly convert nil -- expected %v, received %v",
			0, Float32(nil))
	}
}

func TestFloat32Ptr(t *testing.T) {
	v := float32(0)
	if *Float32Ptr(v) != v {
		t.Errorf("to: Float32Ptr failed to return the correct string -- expected %v, received %v",
			v, *Float32Ptr(v))
	}
}

func TestFloat64(t *testing.T) {
	v := float64(0)
	if Float64(&v) != v {
		t.Errorf("to: Float64 failed to return the correct string -- expected %v, received %v",
			v, Float64(&v))
	}
}

func TestFloat64HandlesNil(t *testing.T) {
	if Float64(nil) != float64(0) {
		t.Errorf("to: Float64 failed to correctly convert nil -- expected %v, received %v",
			0, Float64(nil))
	}
}

func TestFloat64Ptr(t *testing.T) {
	v := float64(0)
	if *Float64Ptr(v) != v {
		t.Errorf("to: Float64Ptr failed to return the correct string -- expected %v, received %v",
			v, *Float64Ptr(v))
	}
}
