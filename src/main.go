package main

import (
	"fmt"
	"sync"
)

// Maximum stored last viewed items
const maxLastViewedItem int = 3

// Store last viewed items
var lastViewedItems []*Item

var mux sync.Mutex

// Item is any item in the store (Eg. Vehicle, Accessory)
type Item interface {
	ViewItem()
	GetDetailsOfItem()
}

// Vehicle is one of the item in store
type Vehicle struct {
	Name  string
	Price float64
}

// Accessory is one of the item for a specific vehicle in the store
type Accessory struct {
	Name  string
	Price float64
}

func (v *Vehicle) ViewItem() {
	viewItem(v)
}

func (a *Accessory) ViewItem() {
	viewItem(a)
}

func viewItem(item Item) {
	mux.Lock()
	defer mux.Unlock()
	// Check if the item is already visited
	for i, lastViewedItem := range lastViewedItems {
		if item == *lastViewedItem {
			lastViewedItems = append(lastViewedItems[:i], lastViewedItems[i+1:]...)
		}
	}
	// remove item from last viewed if the size of array has become greater than max size
	if len(lastViewedItems) >= maxLastViewedItem {
		lastViewedItems = lastViewedItems[1:]
	}
	lastViewedItems = append(lastViewedItems, &item)
}

func (v *Vehicle) GetDetailsOfItem() {
	fmt.Println("Name:", v.Name, "Price:", v.Price)
}

func (a *Accessory) GetDetailsOfItem() {
	fmt.Println("Name:", a.Name, "Price:", a.Price)
}

// GetLastViewedItems will print the last viewed items
func GetLastViewedItems() {
	mux.Lock()
	defer mux.Unlock()

	// Print items in reverse order
	for i := len(lastViewedItems) - 1; i >= 0; i-- {
		item := *lastViewedItems[i]
		item.GetDetailsOfItem()
	}
}

func main() {
	var item1 Item = &Vehicle{"Bajaj Pulsar 150", float64(100000)}
	var item2 Item = &Vehicle{"BMW 600d ", float64(10000000000)}

	var item3 Item = &Accessory{"Nitro Boost for Pulsar 150", float64(12000)}
	var item4 Item = &Accessory{"Headlights for BMW 600d", float64(80000)}

	item1.ViewItem()
	item2.ViewItem()
	item3.ViewItem()
	// Here the first item will get poped from lastViewedItems
	item4.ViewItem()
	// Viewing the same item multiple items
	item4.ViewItem()
	item4.ViewItem()
	item4.ViewItem()
	// Here the 2nd item will come on the top
	item2.ViewItem()
	item3.ViewItem()

	GetLastViewedItems()
}
