package api

import (
	"bytes"
	"encoding/json"
	"github.com/sakiib/apiServer/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	testCases := []struct {
		Method             string
		ExpectedStatusCode int
	}{
		{Method: "GET", ExpectedStatusCode: http.StatusOK},
	}

	for index, test := range testCases {
		request, err := http.NewRequest(test.Method, "http://localhost:8080/api/users", nil)
		if err != nil {
			t.Fatalf("unable to create any request: %v", err)
		}

		response := httptest.NewRecorder()
		GetUsers(response, request)

		if res := response.Result(); res.StatusCode != test.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, test.ExpectedStatusCode, res.Status)
		}
	}
}

func TestGetUser(t *testing.T) {
	testCases := []struct {
		Method             string
		ID                 string
		ExpectedStatusCode int
	}{
		{Method: "GET", ID: "1", ExpectedStatusCode: http.StatusOK},
		{Method: "GET", ID: "2", ExpectedStatusCode: http.StatusOK},
		{Method: "GET", ID: "3", ExpectedStatusCode: http.StatusOK},
		{Method: "GET", ID: "4", ExpectedStatusCode: http.StatusOK},
		{Method: "GET", ID: "5", ExpectedStatusCode: http.StatusNoContent},
	}

	for index, test := range testCases {
		request, err := http.NewRequest(test.Method, "http://localhost:8080/api/user/userID?id="+test.ID, nil)
		if err != nil {
			t.Fatalf("unable to create any request: %v", err)
		}

		response := httptest.NewRecorder()
		GetUser(response, request)

		if res := response.Result(); res.StatusCode != test.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, test.ExpectedStatusCode, res.Status)
		}
	}
}

func TestAddUser(t *testing.T) {
	testCases := []struct {
		Method             string
		Item               model.User
		ExpectedStatusCode int
	}{
		{Method: "POST", Item: model.User{ID: "5", FirstName: "kamol", LastName: "hasan"}, ExpectedStatusCode: http.StatusCreated},
		{Method: "POST", Item: model.User{ID: "1", FirstName: "emruz", LastName: "hossain"}, ExpectedStatusCode: http.StatusConflict},
		{Method: "POST", Item: model.User{ID: "5", FirstName: "emruz", LastName: "hossain"}, ExpectedStatusCode: http.StatusConflict},
	}

	for index, test := range testCases {
		b, _ := json.Marshal(test.Item)
		iorData := bytes.NewReader(b)
		request, err := http.NewRequest(test.Method, "http://localhost:8080/api/user/userID?id="+test.Item.ID, iorData)
		if err != nil {
			t.Fatalf("unable to create any request: %v", err)
		}

		response := httptest.NewRecorder()
		AddUser(response, request)

		if res := response.Result(); res.StatusCode != test.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, test.ExpectedStatusCode, res.Status)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	testCases := []struct {
		Method             string
		Item               model.User
		ExpectedStatusCode int
	}{
		{Method: "PUT", Item: model.User{ID: "1", FirstName: "emruz", LastName: "hossain"}, ExpectedStatusCode: http.StatusCreated},
		{Method: "PUT", Item: model.User{ID: "6", FirstName: "some", LastName: "one"}, ExpectedStatusCode: http.StatusNoContent},
	}

	for index, test := range testCases {
		b, _ := json.Marshal(test.Item)
		iorData := bytes.NewReader(b)
		request, err := http.NewRequest(test.Method, "http://localhost:8080/api/user/userID?id="+test.Item.ID, iorData)
		if err != nil {
			t.Fatalf("unable to create any request: %v", err)
		}

		response := httptest.NewRecorder()
		UpdateUser(response, request)

		if res := response.Result(); res.StatusCode != test.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, test.ExpectedStatusCode, res.Status)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	testCases := []struct {
		Method             string
		Item               model.User
		ExpectedStatusCode int
	}{
		{Method: "DELETE", Item: model.User{ID: "1", FirstName: "emruz", LastName: "hossain"}, ExpectedStatusCode: http.StatusOK},
		{Method: "DELETE", Item: model.User{ID: "1", FirstName: "emruz", LastName: "hossain"}, ExpectedStatusCode: http.StatusNoContent},
		{Method: "DELETE", Item: model.User{ID: "10", FirstName: "user", LastName: "unavailable"}, ExpectedStatusCode: http.StatusNoContent},
	}

	for index, test := range testCases {
		b, _ := json.Marshal(test.Item)
		iorData := bytes.NewReader(b)
		request, err := http.NewRequest(test.Method, "http://localhost:8080/api/user/userID?id="+test.Item.ID, iorData)
		if err != nil {
			t.Fatalf("unable to create any request: %v", err)
		}

		response := httptest.NewRecorder()
		DeleteUser(response, request)

		if res := response.Result(); res.StatusCode != test.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, test.ExpectedStatusCode, res.Status)
		}
	}
}
