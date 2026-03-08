package main

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []*Bin
}

func constructBin(id string, name string, private bool) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("Id or name can't be empty")
	}
	bin := Bin{
		id:        id,
		name:      name,
		createdAt: time.Now(),
		private:   private,
	}

	return &bin, nil
}

func constructBinList(bins []*Bin) *BinList {
	slice := make([]*Bin, 0)
	binList := BinList{
		bins: append(slice, bins...),
	}

	return &binList
}

func main() {}
