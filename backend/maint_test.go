package main

import (
	"testing"
)

func TestGetValue(t *testing.T) {
	sonuc := GetValue("test")
	if sonuc != rClient().Close().Error() {
		t.Log("Test succesful.")
	} else {
		t.Log("Test fail!")
	}
}

func TestSetValue(t *testing.T) {
	sonuc := SetValue("test", "34")
	if sonuc != nil {
		t.Log("Test fail!", sonuc)
	} else {
		t.Log("Test succesful.")
	}
}

func TestCheckValue(t *testing.T) {
	sonuc := CheckValue("test")
	if sonuc != true || false {
		t.Log("Test fail!", sonuc)
	} else {
		t.Log("Test succesful.")
	}
}
