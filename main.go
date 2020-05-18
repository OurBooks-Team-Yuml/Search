package main

import (
        "fmt"
        "log"
        "encoding/json"
        "net/http"

        "github.com/graphql-go/graphql"
        "github.com/graphql-go/handler"
)

func FullSearchAuthors() (map[string]interface{}) {
        var (
                r  map[string]interface{}
        )

        res, err := SearchAuthors()

        if err != nil {
                log.Fatalf("%s", err)
                return nil
        }

        defer res.Body.Close()

        if res.IsError() {
                var e map[string]interface{}
                if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
                        log.Fatalf("Error parsing the response body: %s", err)
                        return nil
                } else {
                        // Print the response status and error information.
                        log.Fatalf("[%s] %s: %s",
                                res.Status(),
                                e["error"].(map[string]interface{})["type"],
                                e["error"].(map[string]interface{})["reason"],
                        )
                        return nil
                }
        }

        if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
                log.Fatalf("Error parsing the response body: %s", err)
                return nil
        }

        return r
}

func FullSearchBooks() (map[string]interface{}) {
        var (
                r  map[string]interface{}
        )

        res, err := SearchBooks()

        if err != nil {
                log.Fatalf("%s", err)
                return nil
        }

        defer res.Body.Close()

        if res.IsError() {
                var e map[string]interface{}
                if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
                        log.Fatalf("Error parsing the response body: %s", err)
                        return nil
                } else {
                        log.Fatalf("[%s] %s: %s",
                                res.Status(),
                                e["error"].(map[string]interface{})["type"],
                                e["error"].(map[string]interface{})["reason"],
                        )
                        return nil
                }
        }

        if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
                log.Fatalf("Error parsing the response body: %s", err)
                return nil
        }

        return r
}

func main() {
        schemaConfig := graphql.SchemaConfig{Query: queryType}
        schema, _ := graphql.NewSchema(schemaConfig)
        
        h := handler.New(&handler.Config{
                Schema: &schema,
                Pretty: true,
                GraphiQL: true,
        })

        fmt.Println("Server is running on port 8003")
        http.Handle("/search", h)
        http.ListenAndServe(":8003", nil)
}
