package tailscale

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestCreateNode(t *testing.T) {
	handler := &Handler{}

	newNode := Node{
		HostName:   "test-node",
		Router:     "test-router",
		AllowedIPs: []string{"192.168.1.1/32"},
	}

	body, _ := json.Marshal(newNode)
	req, err := http.NewRequest("POST", "/create_node", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.CreateNode(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Data received successfully"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAllNodes(t *testing.T) {
	nodes := []Node{
		{
			HostName:   "test-node-1",
			Router:     "router-1",
			AllowedIPs: []string{"192.168.1.1/32"},
		},
		{
			HostName:   "test-node-2",
			Router:     "router-2",
			AllowedIPs: []string{"192.168.1.2/32"},
		},
	}

	handler := &Handler{
		Nodes: nodes,
	}

	for _, node := range nodes {
		body, _ := json.Marshal(node)
		_, err := http.NewRequest("POST", "/create_node", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
	}

	req, err := http.NewRequest("GET", "/get_nodes", nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()
	handler.GetAllNodes(newRecorder, req)

	if status := newRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "[{\"Router\":\"router-1\",\"ID\":\"\",\"HostName\":\"test-node-1\",\"OS\":\"\",\"AllowedIPs\":[\"192.168.1.1/32\"],\"CurAddr\":\"\",\"Active\":false},{\"Router\":\"router-2\",\"ID\":\"\",\"HostName\":\"test-node-2\",\"OS\":\"\",\"AllowedIPs\":[\"192.168.1.2/32\"],\"CurAddr\":\"\",\"Active\":false}]"
	if slices.Equal(newRecorder.Body.Bytes(), []byte(expected)) {
		t.Errorf("handler returned unexpected body: \ngot\n %v\n expected\n %v",
			newRecorder.Body.Bytes(), []byte(expected))
	}
}

func TestFindNodeByHostName(t *testing.T) {
	node := Node{
		HostName:   "test-node-1",
		Router:     "router-1",
		AllowedIPs: []string{"192.168.1.1/32"},
	}

	handler := &Handler{
		Nodes: []Node{node},
	}

	body, _ := json.Marshal(node)
	_, err := http.NewRequest("POST", "/create_node", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("GET", "/get_nodes/"+node.HostName, nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()
	handler.GetAllNodes(newRecorder, request)

	if status := newRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "[{\"Router\":\"router-1\",\"ID\":\"\",\"HostName\":\"test-node-1\",\"OS\":\"\",\"AllowedIPs\":[\"192.168.1.1/32\"],\"CurAddr\":\"\",\"Active\":false}]"
	if slices.Equal(newRecorder.Body.Bytes(), []byte(expected)) {
		t.Errorf("handler returned unexpected body: \ngot\n %v\n expected\n %v",
			newRecorder.Body.Bytes(), []byte(expected))
	}
}
