package main

import( 
	"testing"
)


func TestN_1(t *testing.T){
	s1 := "111"
	s2 := "222"
	res,err := sum(&s1,&s2)
	expected := "333"

	if err !=nil{
		t.Errorf("FAILED %s", err)
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}
func TestN_2(t *testing.T){
	s1 := ""
	s2 := "222"
	res,err := sum(&s1,&s2)
	expected := "222"

	if err !=nil{
		t.Errorf("FAILED %s", err)
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}
func TestN_3(t *testing.T){
	s1 := "111"
	s2 := ""
	res,err := sum(&s1,&s2)
	expected := "111"

	if err !=nil{
		t.Errorf("FAILED %s", err)
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}
func TestN_4(t *testing.T){
	s2 := "222"
	res,err := sum(nil,&s2)
	expected := "222"

	if err !=nil{
		t.Errorf("FAILED %s", err)
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}
func TestN_5(t *testing.T){
	s1 := "111"
	res,err := sum(&s1,nil)
	expected := "111"

	if err !=nil{
		t.Errorf("FAILED %s", err)
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}
func TestN_6(t *testing.T){
	s1 := "111"
	s2 := "2a"
	res,err := sum(&s1,&s2)
	expected := "error"

	if err !=nil{
		test := ExNumberFormationException{"ExNumberFormationException in string 2 at index:",1}
		if err == test {                                                                        
			t.Logf("PASS")
		}
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}
func TestN_7(t *testing.T){
	s1 := "111a"
	s2 := "2"
	res,err := sum(&s1,&s2)
	expected := "error"

	if err !=nil{
		test := ExNumberFormationException{"ExNumberFormationException in string 1 at index:",3}
		if err == test {                                                                        
			t.Logf("PASS")
		}
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}
func TestN_8(t *testing.T){
	s1 := "111"
	s2 := "-2"
	res,err := sum(&s1,&s2)
	expected := "error"

	if err !=nil{
		test := ExNumberFormationException{"ExNumberFormationException in string 2 at index:",0}
		if err == test {                                                                        
			t.Logf("PASS")
		}
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}
func TestN_9(t *testing.T){
	s1 := "111"
	s2 := "*21212"
	res,err := sum(&s1,&s2)
	expected := "error"

	if err !=nil{
		test := ExNumberFormationException{"ExNumberFormationException in string 2 at index:",0}
		if err == test {                                                                        
			t.Logf("PASS")
		}
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}

func TestN_10(t *testing.T){
	s1 := "111112"
	s2 := "999999"
	res,err := sum(&s1,&s2)
	expected := "1111111"

	if err !=nil{
		t.Errorf("FAILED %s", err)
	}else{
		if res != expected{
			t.Errorf("FAILED. Expected: %s, Got %s\n",expected,res)
		}else{
			t.Logf("PASS")
		}
	}
}