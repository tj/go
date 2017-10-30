// Package dynamo provides dynamodb utilities.
package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Item is a map of attributes.
type Item map[string]*dynamodb.AttributeValue

// Marshal returns a new item from struct.
func Marshal(value interface{}) (Item, error) {
	v, err := dynamodbattribute.MarshalMap(value)
	if err != nil {
		return nil, err
	}

	return Item(v), nil
}

// MustMarshal returns a new item from struct.
func MustMarshal(value interface{}) Item {
	v, err := dynamodbattribute.MarshalMap(value)
	if err != nil {
		panic(err)
	}

	return Item(v)
}

// Unmarshal the item.
func Unmarshal(i Item, value interface{}) error {
	return dynamodbattribute.UnmarshalMap(i, value)
}
