package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gandarez/video-game-api/internal/model"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)


type ConsoleCreateParams struct {
	Name         string `json:"name" jsonschema:"the name of the console"`
	Manufacturer string `json:"manufacturer" jsonschema:"the manufacturer of the console"`
	ReleaseDate  string `json:"release_date" jsonschema:"the release date of the console in YYYY-MM-DD format"`
}


func CreateConsole(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[ConsoleCreateParams]) (*mcp.CallToolResultFor[any], error) {
	consoleInsert := model.ConsoleInsert{
		Name:         params.Arguments.Name,
		Manufacturer: params.Arguments.Manufacturer,
		ReleaseDate:  params.Arguments.ReleaseDate,
	}

	if err := consoleInsert.Validate(); err != nil {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Validation error: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	
	url := "http://localhost:17020/consoles"

	data, err := json.Marshal(consoleInsert)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Failed to marshal console data: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	reqBody := strings.NewReader(string(data))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reqBody)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Failed to create request: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Failed to send request: %v", err),
				},
			},
			IsError: true,
		}, nil
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Failed to read response: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	if resp.StatusCode != http.StatusCreated {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Console creation failed with status %d: %s", resp.StatusCode, string(respBody)),
				},
			},
			IsError: true,
		}, nil
	}

	var console model.Console
	if err := json.Unmarshal(respBody, &console); err != nil {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Failed to parse response: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Console created successfully!\nID: %s\nName: %s\nManufacturer: %s\nRelease Date: %s",
					console.ID, console.Name, console.Manufacturer, console.ReleaseDate),
			},
		},
	}, nil
}

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "video-game-api",
			Version: "v1.0.0",
		},
		nil,
	)
	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "console_create",
			Description: "Create a new video game console",
		},
		CreateConsole,
	)
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}