package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestApplyStatusToResponse(t *testing.T) {
	resp := new(Response)

	if ApplyStatusToResponse(resp, StatusOK); resp.StatusCode != StatusOK {
		t.Errorf("expected %d, got %d", StatusOK, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusOK)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}
}

var req Request = Request{
	URL: &url.URL{
		Scheme: "http",
		Host:   "www.techcrunch.com",
		Path:   "/",
	},
	Proto: "HTTP/1.1",
	Header: Header{
		"Accept":           {"text/html"},
		"Host":             {"localhost.com"},
		"User-Agent":       {"Fake"},
	},
	Body: "abcdef\n",
}

func TestHeadRequestWithGetMethodEndpoint(t *testing.T) {
	req, resp := &req, new(Response)
	req.Method = "GET"
	readTransfer(resp, *req)
	HeadRequest(resp, req)

	if resp.StatusCode != StatusMethodNotAllowed {
		t.Errorf("expected %d, got %d", StatusMethodNotAllowed, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusMethodNotAllowed)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}

	expectedAllowHeader := []string{"HEAD", "OPTIONS"}
	if actualAllowHeader := resp.Header.Values("ALLOW"); !reflect.DeepEqual(actualAllowHeader, expectedAllowHeader) {
		t.Errorf("expected %s, got %s", expectedAllowHeader, actualAllowHeader)
	}

	resp.Header.Del("ALLOW")
	return
}

func TestHeadRequestWithoutGetMethodEndpoint(t *testing.T) {
	req, resp := &req, new(Response)
	req.Method = "HEAD"
	readTransfer(resp, *req)
	HeadRequest(resp, req)

	if resp.StatusCode != StatusOK {
		t.Errorf("expected %d, got %d", StatusOK, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusOK)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}
}

func TestNotFoundEndpoint(t *testing.T) {
	req, resp := &req, new(Response)
	readTransfer(resp, *req)
	NotFound(resp, req)

	if resp.StatusCode != StatusNotFound {
		t.Errorf("expected %d, got %d", StatusNotFound, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusNotFound)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}
}

func TestSimpleGetEndpoint(t *testing.T) {
	req, resp := &req, new(Response)
	readTransfer(resp, *req)
	SimpleGet(resp, req)

	if resp.StatusCode != StatusOK {
		t.Errorf("expected %d, got %d", StatusOK, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusOK)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}
}

func TestSimpleGetWithBodyEndpoint(t *testing.T) {
	req, resp := &req, new(Response)
	readTransfer(resp, *req)
	SimpleGetWithBody(resp, req)

	if resp.StatusCode != StatusOK {
		t.Errorf("expected %d, got %d", StatusOK, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusOK)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}

	expectedBody := "Hello world"
	if actualBody := resp.Body; strings.Compare(actualBody, expectedBody) != 0 {
		t.Errorf("expected %s, got %s", expectedBody, actualBody)
	}
}

func TestSimpleHeadEndpoint(t *testing.T) {
	req, resp := &req, new(Response)
	readTransfer(resp, *req)
	SimpleHead(resp, req)

	if resp.StatusCode != StatusOK {
		t.Errorf("expected %d, got %d", StatusOK, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusOK)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}
}

func TestMethodOptionsEndpoint(t *testing.T) {
	req, resp := &req, new(Response)
	req.Method = "GET"
	readTransfer(resp, *req)
	MethodOptions(resp, req)

	if resp.StatusCode != StatusOK {
		t.Errorf("expected %d, got %d", StatusOK, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusOK)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}

	expectedAllowHeader := []string{"GET", "HEAD", "OPTIONS"}
	if actualAllowHeader := resp.Header.Values("ALLOW"); !reflect.DeepEqual(actualAllowHeader, expectedAllowHeader) {
		t.Errorf("expected %s, got %s", expectedAllowHeader, actualAllowHeader)
	}

	resp.Header.Del("ALLOW")
	return
}

func TestMethodOptions2Endpoint(t *testing.T) {
	req, resp := &req, new(Response)
	req.Method = "GET"
	readTransfer(resp, *req)
	MethodOptions2(resp, req)

	if resp.StatusCode != StatusOK {
		t.Errorf("expected %d, got %d", StatusOK, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusOK)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}

	expectedAllowHeader := []string{"GET", "HEAD", "OPTIONS", "PUT", "POST"}
	if actualAllowHeader := resp.Header.Values("ALLOW"); !reflect.DeepEqual(actualAllowHeader, expectedAllowHeader) {
		t.Errorf("expected %s, got %s", expectedAllowHeader, actualAllowHeader)
	}

	resp.Header.Del("ALLOW")
	return
}

func TestRedirectEndpoint(t *testing.T) {
	req, resp := &req, new(Response)
	req.Method = "GET"
	readTransfer(resp, *req)
	Redirect(resp, req)

	if resp.StatusCode != StatusMovedPermanently {
		t.Errorf("expected %d, got %d", StatusMovedPermanently, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusMovedPermanently)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}

	host := viperEnvVariable("REDIRECT_HOST")
	port := viperEnvVariable("server.port")
	endpoint := viperEnvVariable("REDIRECT_ENDPOINT")
	address := host + ":" + port + endpoint

	expectedLocationHeader := address
	if actualLocationHeader := resp.Header.Get("Location"); strings.Compare(actualLocationHeader, expectedLocationHeader) != 0 {
		t.Errorf("expected %s, got %s", expectedLocationHeader, actualLocationHeader)
	}

	resp.Header.Del("Location")
	return
}

func TestEchoBodyEndpoint(t *testing.T) {
	req, resp := &req, new(Response)

	expectedBody := "Test message"
	req.Body = expectedBody

	req.Method = "POST"
	readTransfer(resp, *req)
	EchoBody(resp, req)

	if resp.StatusCode != StatusOK {
		t.Errorf("expected %d, got %d", StatusOK, resp.StatusCode)
	}

	expectedStatus := StatusText(StatusOK)
	if resp.Status != expectedStatus {
		t.Errorf("expected %s, got %s", expectedStatus, resp.Status)
	}

	if actualBody := resp.Body; strings.Compare(actualBody, expectedBody) != 0 {
		t.Errorf("expected %s, got %s", expectedBody, actualBody)
	}
}
