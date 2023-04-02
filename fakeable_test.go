package gofakeit

import (
	"fmt"
	"reflect"
	"testing"
)

type strTyp string

func (t strTyp) Fake(faker *Faker) interface{} {
	return faker.FirstName()
}

type strTypPtr string

func (t *strTypPtr) Fake(faker *Faker) interface{} {
	return strTypPtr("hello test ptr")
}

type testStruct1 struct {
	B string `fake:"{firstname}"`
}

type testStruct2 struct {
	B strTyp
}

func TestIsFakeable(t *testing.T) {
	var t1 testStruct2
	var t2 *testStruct2
	var t3 strTyp
	var t4 *strTyp
	var t5 strTypPtr
	var t6 *strTypPtr

	if isFakeable(reflect.ValueOf(t1).Type()) {
		t.Errorf("expected testStruct2 not to be fakeable")
	}

	if isFakeable(reflect.ValueOf(t2).Type()) {
		t.Errorf("expected *testStruct2 not to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t3).Type()) {
		t.Errorf("expected strTyp to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t4).Type()) {
		t.Errorf("expected *strTyp to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t5).Type()) {
		t.Errorf("expected strTypPtr to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t6).Type()) {
		t.Errorf("expected *strTypPtr to be fakeable")
	}
}

func ExampleFakeable() {
	var t1 testStruct1
	var t2 testStruct1
	var t3 testStruct2
	var t4 testStruct2
	New(314).Struct(&t1)
	New(314).Struct(&t2)
	New(314).Struct(&t3)
	New(314).Struct(&t4)

	fmt.Printf("%#v\n", t1)
	fmt.Printf("%#v\n", t2)
	fmt.Printf("%#v\n", t3)
	fmt.Printf("%#v\n", t4)
	// Expected Output:
	// gofakeit.testStruct1{B:"Margarette"}
	// gofakeit.testStruct1{B:"Margarette"}
	// gofakeit.testStruct2{B:"Margarette"}
	// gofakeit.testStruct2{B:"Margarette"}
}
