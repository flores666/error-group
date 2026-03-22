package database

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type database1 struct {
	data []Data
}

func NewDatabase1() Database {
	return &database1{
		data: []Data{
			Data{"Value1"},
			Data{"Value2"},
		},
	}
}

type database2 struct {
	data []Data
}

func NewDatabase2() Database {
	return &database2{
		data: []Data{
			Data{"Value3"},
			Data{"Value4"},
		},
	}
}

func (d *database1) Get() ([]Data, error) {
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	fmt.Println("Get 1 is done!")

	if rand.Intn(5) > 3 {
		return []Data{}, errors.New("Error")
	}

	return d.data, nil
}

func (d *database2) Get() ([]Data, error) {
	time.Sleep(time.Duration(2+rand.Intn(2)) * time.Second)
	fmt.Println("Get 2 is done!")

	if rand.Intn(5) > 3 {
		return []Data{}, errors.New("Error")
	}

	return d.data, nil
}
