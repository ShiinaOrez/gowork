package main

import (
    "log"
    "encoding/json"
    "strings"
    "context"

    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {
    es, _ := elasticsearch.NewDefaultClient()
    req := esapi.IndexRequest {
        Index:      "class",
        DocumentID: "49038782",
        Body:       strings.NewReader(`{"title" : "class_1"}`),
        Refresh:    "true",
    }
    resp, err := req.Do(context.Background(), es)
    if err != nil {
        log.Fatalf("Error getting response: %s", err.Error())
    }
    defer resp.Body.Close()

    if resp.IsError() {
        log.Printf("[%s] Error indexing document ID=%d", resp.Status(), 1)
    } else {
        // Deserialize the response into a map.
        var r map[string]interface{}
        if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
            log.Printf("Error parsing the response body: %s", err)
        } else {
            // Print the response status and indexed document version.
            log.Printf("[%s] %s; version=%d", resp.Status(), r["result"], int(r["_version"].(float64)))
        }
    }
}
