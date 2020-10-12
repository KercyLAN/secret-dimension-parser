// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 11:03

package properties

import (
	"fmt"
	"testing"
)

var properties *Properties
func init() {
	var err error
	properties, err = New("./test.properties")
	if err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	properties, err := New("./test.properties")
	if err != nil {
		panic(err)
	}else {
		t.Log(properties)
	}
}

func TestProperties_Each(t *testing.T) {
	properties.Each(func(key string, value string) {
		t.Log(fmt.Sprintf("%v = %v", key, value))
	})
}

func TestProperties_GetBool(t *testing.T) {
	t.Log(properties.GetBool("properties.bool.true"))
	t.Log(properties.GetBool("properties.bool.false"))
}

func TestProperties_GetFloat32(t *testing.T) {
	t.Log(properties.GetFloat32("properties.float"))
	t.Log(properties.GetFloat32("properties.float.space"))
}

func TestProperties_GetFloat64(t *testing.T) {
	t.Log(properties.GetFloat64("properties.float"))
	t.Log(properties.GetFloat64("properties.float.space"))
}

func TestProperties_GetInt(t *testing.T) {
	t.Log(properties.GetInt("properties.int"))
	t.Log(properties.GetInt("properties.int.space"))
}

func TestProperties_GetInt8(t *testing.T) {
	t.Log(properties.GetInt8("properties.int"))
	t.Log(properties.GetInt8("properties.int.space"))
}

func TestProperties_GetInt16(t *testing.T) {
	t.Log(properties.GetInt16("properties.int"))
	t.Log(properties.GetInt16("properties.int.space"))
}

func TestProperties_GetInt32(t *testing.T) {
	t.Log(properties.GetInt32("properties.int"))
	t.Log(properties.GetInt32("properties.int.space"))
}

func TestProperties_GetInt64(t *testing.T) {
	t.Log(properties.GetInt64("properties.int"))
	t.Log(properties.GetInt64("properties.int.space"))
}

func TestProperties_GetUint(t *testing.T) {
	t.Log(properties.GetUint("properties.int"))
	t.Log(properties.GetUint("properties.int.space"))
}

func TestProperties_GetUint8(t *testing.T) {
	t.Log(properties.GetUint8("properties.int"))
	t.Log(properties.GetUint8("properties.int.space"))
}

func TestProperties_GetUint16(t *testing.T) {
	t.Log(properties.GetUint16("properties.int"))
	t.Log(properties.GetUint16("properties.int.space"))
}

func TestProperties_GetUint32(t *testing.T) {
	t.Log(properties.GetUint32("properties.int"))
	t.Log(properties.GetUint32("properties.int.space"))
}

func TestProperties_GetUint64(t *testing.T) {
	t.Log(properties.GetUint64("properties.int"))
	t.Log(properties.GetUint64("properties.int.space"))
}

func TestProperties_GetInterface(t *testing.T) {
	t.Log(properties.GetInterface("properties.string"))
}

func TestProperties_GetString(t *testing.T) {
	t.Log(properties.GetString("properties.string"))
}

func TestProperties_HasKey(t *testing.T) {
	t.Log(properties.HasKey("properties.string"))
	t.Log(properties.HasKey("properties.strings"))
}