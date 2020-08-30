package handler

import (
	"errors"
	"log"
	"strconv"
)

/*Sort ...
query parameter
*/
type Sort string

const (
	// UpdatedAt ... sort by created_at
	UpdatedAt = Sort("updated_at")

	// CreatedAt ... sort by updated_at
	CreatedAt = Sort("created_at")

	// Content ... sort by content
	Content = Sort("content")

	// Ascending ... ascending order
	Ascending = Order("asc")

	// Descending ... descending order
	Descending = Order("desc")
)

var sortOptions = map[Sort]Sort{
	CreatedAt: CreatedAt,
	UpdatedAt: UpdatedAt,
	Content:   Content,
}

// Valid ... check if it's valid
func (s Sort) Valid() error {
	_, ok := sortOptions[s]
	if !ok {
		log.Println("error while validating query param: sort")
		log.Printf("value: %s", string(s))
		return errors.New("invalid query param: sort")
	}
	return nil
}

// String ... get the string value
func (s Sort) String() string {
	return string(s)
}

/*Order ...
query parameter
*/
type Order string

var orderOptions = map[Order]Order{
	Ascending:  Ascending,
	Descending: Descending,
}

// Valid ... check if it's valid
func (o Order) Valid() error {
	_, ok := orderOptions[o]
	if !ok {
		log.Println("error while validating query param: order")
		log.Printf("value: %s", string(o))
		return errors.New("invalid query param: order")
	}
	return nil
}

// String ... get the string value
func (o Order) String() string {
	return string(o)
}

/*Limit ...
query parameter
*/
type Limit string

// Valid ... check if valid
func (l Limit) Valid() error {
	_, err := strconv.Atoi(string(l))
	if err != nil {
		log.Println("error while validating query param: limit")
		log.Printf("value: %s, error: %s", string(l), err.Error())
		return errors.New("invalid query param: limit")
	}
	return nil
}

// Int ... get the int32 value
func (l Limit) Int() int32 {
	v, _ := strconv.Atoi(string(l))
	return int32(v)
}

/*Offset ...
query parameter
*/
type Offset string

// Valid ... check if valid
func (o Offset) Valid() error {
	_, err := strconv.Atoi(string(o))
	if err != nil {
		log.Println("error while validating query param: offset")
		log.Printf("value: %s, error: %s", string(o), err.Error())
		return errors.New("invalid query param: offset")
	}
	return nil
}

// Int ... get the int32 value
func (o Offset) Int() int32 {
	l, _ := strconv.Atoi(string(o))
	return int32(l)
}
