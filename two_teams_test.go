package main

import "testing"

func TestStuff(t *testing.T) {
	var a, b, n = 10, 8, 2
	var expected = "Yes"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}

func TestStuff2(t *testing.T) {
	var a, b, n = 3, 10, 3
	var expected = "No"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}

func TestZero(t *testing.T) {
	var a, b, n = 0, 0, 10000
	var expected = "No"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}
func TestZeroaoeu(t *testing.T) {
	var a, b, n = 100, 0, 10000
	var expected = "Yes"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}
func TestZeroaoeu2(t *testing.T) {
	var a, b, n = 0, 100, 10000
	var expected = "No"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}
func TestZeroaoeu2One(t *testing.T) {
	var a, b, n = 0, 100, 1
	var expected = "No"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}
func TestZeroaoeu2Onei(t *testing.T) {
	var a, b, n = 100, 0, 1
	var expected = "Yes"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}

func Test(t *testing.T) {
	var a, b, n = 2, 10, 5
	var expected = "No"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}
func TestA(t *testing.T) {
	var a, b, n = 0, 0, 1
	var expected = "No"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}

func TestB(t *testing.T) {
	var a, b, n = 10000, 10000, 10000
	var expected = "Yes"
	var actual = fn(a, b, n)
	if expected != actual {
		t.Fatal("aoeuhtnsS")
	}
}
