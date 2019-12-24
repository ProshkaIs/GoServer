package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*Tests for GetAll()*/

func TestGetAll(t *testing.T) {
	var jsonStr1 = []byte("{\"offset\":0}")
	var jsonStr2 = []byte("{\"offset\":1}")
	var jsonStr3 = []byte("{\"offset\":0,\"price_sort\":\"desc\"}")
	var jsonStr4 = []byte("{\"offset\":0,\"price_sort\":\"asc\"}")
	var jsonStr5 = []byte("{\"offset\":0,\"date_sort\":\"desc\"}")
	var jsonStr6 = []byte("{\"offset\":0,\"date_sort\":\"asc\"}")
	var mas [6] *http.Request
	mas[0], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr1))
	mas[1], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr2))
	mas[2], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr3))
	mas[3], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr4))
	mas[4], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr5))
	mas[5], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr6))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAll)

	for i:=0;i<len(mas);i++{
		handler.ServeHTTP(rr, mas[i])
		not_expected := `400`
		if rr.Body.String() == not_expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), not_expected)
		}
	}
}

func TestGetAll400(t *testing.T) {
	var jsonStr1 = []byte("{\"off_set\":0}")
	var jsonStr2 = []byte("{\"offset\":}")
	var jsonStr3 = []byte("{\"offset\":0,\"price_sort\":\"true\"}")
	var jsonStr4 = []byte("{\"offset\":0,\"date_sort\":\"true\"}")
	var mas [4] *http.Request
	mas[0], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr1))
	mas[1], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr2))
	mas[2], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr3))
	mas[3], _ = http.NewRequest("POST", "/ad/getall", bytes.NewBuffer(jsonStr4))



	for i:=0;i<len(mas);i++{
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetAll)
		handler.ServeHTTP(rr, mas[i])
		expected := `400`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}

/*Tests fo GetOne()*/

func TestGetOne(t *testing.T) {
	var jsonStr1 = []byte("{\"id\":1}")
	var jsonStr2 = []byte("{\"id\":2}")
	var jsonStr3 = []byte("{\"id\":1,\"fields\":\"true\"}")
	var mas [3] *http.Request
	mas[0], _ = http.NewRequest("POST", "/ad/getone", bytes.NewBuffer(jsonStr1))
	mas[1], _ = http.NewRequest("POST", "/ad/getone", bytes.NewBuffer(jsonStr2))
	mas[2], _ = http.NewRequest("POST", "/ad/getone", bytes.NewBuffer(jsonStr3))



	for i:=0;i<len(mas);i++{
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetOne)
		handler.ServeHTTP(rr, mas[i])
		not_expected := "400"
		if rr.Body.String() == not_expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), not_expected)
		}
	}
}


func TestGetOne400(t *testing.T) {
	var jsonStr1 = []byte("{\"i_d\":1}")
	var jsonStr2 = []byte("{\"id\":0}")
	var jsonStr3 = []byte("{\"id\":}")
	var mas [3] *http.Request
	mas[0], _ = http.NewRequest("POST", "/ad/getone", bytes.NewBuffer(jsonStr1))
	mas[1], _ = http.NewRequest("POST", "/ad/getone", bytes.NewBuffer(jsonStr2))
	mas[2], _ = http.NewRequest("POST", "/ad/getone", bytes.NewBuffer(jsonStr3))



	for i:=0;i<len(mas);i++{
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetOne)
		handler.ServeHTTP(rr, mas[i])
		expected := "400"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}

/*Tests for SetOne*/

func TestSetOne(t *testing.T) {
	var jsonStr1 = []byte("{\"name\":\"test2\",\"link\":\"test_link\",\"price\":111222,\"description\":\"testtesttest\"}")
	req, _ := http.NewRequest("POST", "/ad/setone", bytes.NewBuffer(jsonStr1))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SetOne)

		handler.ServeHTTP(rr, req)
		not_expected := `400`
		if rr.Body.String() == not_expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), not_expected)
		}

}

func TestSetOne400(t *testing.T) {
	var jsonStr1 = []byte("{\"na_me\":\"test2\",\"link\":\"test_link\",\"price\":111222,\"description\":\"testtesttest\"}")
	var jsonStr2 = []byte("{\"name\":\"test2\",\"li_nk\":\"test_link\",\"price\":111222,\"description\":\"testtesttest\"}")
	var jsonStr3 = []byte("{\"name\":\"test2\",\"link\":\"test_link\",\"pri_ce\":111222,\"description\":\"testtesttest\"}")
	var jsonStr4 = []byte("{\"name\":\"test2\",\"link\":\"test_link\",\"price\":111222,\"descr_iption\":\"testtesttest\"}")
	var mas [4] *http.Request
	mas[0], _ = http.NewRequest("POST", "/ad/setone", bytes.NewBuffer(jsonStr1))
	mas[1], _ = http.NewRequest("POST", "/ad/setone", bytes.NewBuffer(jsonStr2))
	mas[2], _ = http.NewRequest("POST", "/ad/setone", bytes.NewBuffer(jsonStr3))
	mas[3], _ = http.NewRequest("POST", "/ad/setone", bytes.NewBuffer(jsonStr4))



	for i:=0;i<len(mas);i++{
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SetOne)
		handler.ServeHTTP(rr, mas[i])
		expected := "400"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}



