package main

import (
        "strconv"

        "github.com/graphql-go/graphql"
)

func CastToString(value interface{}) (casted *string) {
        if str, ok := value.(string); ok {
                return &str
        } else {
                return nil
        }
}

func CastToInt(value interface{}) (casted *int64) {
        if id, ok := value.(int64); ok {
                return &id
        } else {
                return nil
        }
}

type Book struct {
        ID int64 `json:"id"`
        Name string  `json:"name"`
        Description string `json:"description"`
        Image *string `json:image_path`
        ISBN *string `json:isbn`
        House *string `json:publishing_house`
        Date *string `json:published_date`
        Categories []interface{} `json:categories`
        Authors []interface{} `json:authors_id`
        RelatedBook *int64 `json:related_book`
}

type Author struct {
        ID int64 `json:"id"`
        FirstName string  `json:"first_name"`
        LastName string `json:"last_name"`
        FullName string `json:"full_name"`
        BDay *string `json:birthday_date`
        Biography *string `json:biography`
        Image *string `json:image_path`
        Books []interface{} `json:books`
}

var bookType = graphql.NewObject(
        graphql.ObjectConfig{
                Name: "Book",
                Fields: graphql.Fields{
                        "id": &graphql.Field{
                                Type: graphql.Int,
                        },
                        "name": &graphql.Field{
                                Type: graphql.String,
                        },
                        "description": &graphql.Field{
                                Type: graphql.String,
                        },
                        "imagePath": &graphql.Field{
                                Type: graphql.String,
                        },
                        "isbn": &graphql.Field{
                                Type: graphql.String,
                        },
                        "publishingHouse": &graphql.Field{
                                Type: graphql.String,
                        },
                        "publishedDate": &graphql.Field{
                                Type: graphql.String,
                        },
                        "categories": &graphql.Field{
                                Type: graphql.NewList(graphql.Int),
                        },
                        "authorsId": &graphql.Field{
                                Type: graphql.NewList(graphql.Int),
                        },
                        "relatedBook": &graphql.Field{
                                Type: graphql.Int,
                        },
                },
        },
)

var authorType = graphql.NewObject(
        graphql.ObjectConfig{
                Name: "Author",
                Fields: graphql.Fields{
                        "id": &graphql.Field{
                                Type: graphql.Int,
                        },
                        "firstName": &graphql.Field{
                                Type: graphql.String,
                        },
                        "lastName": &graphql.Field{
                                Type: graphql.String,
                        },
                        "fullName": &graphql.Field{
                                Type: graphql.String,
                        },
                        "birthdayDate": &graphql.Field{
                                Type: graphql.String,
                        },
                        "biography": &graphql.Field{
                                Type: graphql.String,
                        },
                        "imagePath": &graphql.Field{
                                Type: graphql.String,
                        },
                        "books": &graphql.Field{
                                Type: graphql.NewList(graphql.Int),
                        },
                },
        },
)

var queryType = graphql.NewObject(
        graphql.ObjectConfig{
                Name: "Query",
                Fields: graphql.Fields{
                        "authors": &graphql.Field{
                                Type: graphql.NewList(authorType),
                                Args: graphql.FieldConfigArgument{
                                        "search": &graphql.ArgumentConfig{
                                                Type: graphql.String,
                                        },
                                },
                                Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                                        search := params.Args["search"].(string)
                                        results := FullSearchAuthors(search)
                                        var authors []Author

                                        for _, hit := range results["hits"].(map[string]interface{})["hits"].([]interface{}) {
                                                source := hit.(map[string]interface{})["_source"].(map[string]interface{})

                                                id := hit.(map[string]interface{})["_id"].(string)
                                                id_int, _ := strconv.ParseInt(id, 10, 64)

                                                authors = append(authors, Author{
                                                        ID: id_int,
                                                        FirstName: source["first_name"].(string),
                                                        LastName: source["last_name"].(string),
                                                        FullName: source["full_name"].(string),
                                                        BDay: CastToString(source["birthday_date"]),
                                                        Biography: CastToString(source["biography"]),
                                                        Image: CastToString(source["image_path"]),
                                                        Books: source["books"].([]interface{}),
                                                })
                                        }

                                        return authors, nil
                                },
                        },
                        "books": &graphql.Field{
                                Type: graphql.NewList(bookType),
                                Args: graphql.FieldConfigArgument{
                                        "search": &graphql.ArgumentConfig{
                                                Type: graphql.String,
                                        },
                                },
                                Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                                        search := params.Args["search"].(string)
                                        results := FullSearchBooks(search)
                                        var books []Book

                                        for _, hit := range results["hits"].(map[string]interface{})["hits"].([]interface{}) {
                                                source := hit.(map[string]interface{})["_source"].(map[string]interface{})

                                                id := hit.(map[string]interface{})["_id"].(string)
                                                id_int, _ := strconv.ParseInt(id, 10, 64)

                                                books = append(books, Book{
                                                        ID: id_int,
                                                        Name: source["name"].(string),
                                                        Description: source["description"].(string),
                                                        Image: CastToString(source["image_path"]),
                                                        ISBN: CastToString(source["isbn"]),
                                                        House: CastToString(source["publishing_house"]),
                                                        Date: CastToString(source["publishing_date"]),
                                                        Categories: source["categories"].([]interface{}),
                                                        Authors: source["authors_id"].([]interface{}),
                                                        RelatedBook: CastToInt(source["related_book"]),
                                                })
                                        }

                                        return books, nil
                                },
                        },
                },
        },
)