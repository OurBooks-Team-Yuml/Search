package main

import (
        "bytes"
        "context"
        "encoding/json"

        "github.com/elastic/go-elasticsearch/v7"
        "github.com/elastic/go-elasticsearch/v7/esapi"
)

func SearchAuthors() (*esapi.Response, error) {
        es, err := GetClient()

        if err != nil {
                return nil, err
        }

        buf, err := GetAuthorQuery()

        if err != nil {
                return nil, err
        }

        res, err := Search(es, "authors", buf)

        if err != nil {
                return nil, err
        }

        return res, nil
}

func SearchBooks() (*esapi.Response, error) {
        es, err := GetClient()

        if err != nil {
                return nil, err
        }

        buf, err := GetBookQuery()

        if err != nil {
                return nil, err
        }

        res, err := Search(es, "books", buf)

        if err != nil {
                return nil, err
        }

        return res, nil
}

func GetAuthorQuery() (bytes.Buffer, error) {
        var buf bytes.Buffer
        query := map[string]interface{}{
                "query": map[string]interface{}{
                        "match": map[string]interface{}{
                                "first_name": "ESTest",
                        },
                },
        }

        if err := json.NewEncoder(&buf).Encode(query); err != nil {
                return buf, err
        }

        return buf, nil
}

func GetBookQuery() (bytes.Buffer, error) {
        var buf bytes.Buffer
        query := map[string]interface{}{
                "query": map[string]interface{}{
                        "match": map[string]interface{}{
                                "name": "ESTest",
                        },
                },
        }

        if err := json.NewEncoder(&buf).Encode(query); err != nil {
                return buf, err
        }

        return buf, nil
}

func GetClient() (*elasticsearch.Client, error) {
        es, err := elasticsearch.NewDefaultClient()

        if err != nil {
                return nil, err
        }

        _, err = es.Info()

        if err != nil {
                return nil, err
        }

        return es, nil
}

func Search(es *elasticsearch.Client, index string, query bytes.Buffer) (*esapi.Response, error) {
        return es.Search(
                es.Search.WithContext(context.Background()),
                es.Search.WithIndex(index),
                es.Search.WithBody(&query),
                es.Search.WithTrackTotalHits(true),
                es.Search.WithPretty(),
        )
}