package mcp

import "testing"

var TestApiKey = ""
var TestMcpServiceUrl = ""
var TestMcpServiceToken = ""

func TestClaudeClient(t *testing.T) {
	client := NewClaudeClientWithOptions(WithAPIKey(TestApiKey), WithMcpServerUrl(TestMcpServiceUrl), WithMcpServerToken(TestMcpServiceToken))
	resp, err := client.CallWithMessages("You are a professional cryptocurrency market analyst with many years of experience.", "Please predict Bitcoin's price movement for the next week.")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
