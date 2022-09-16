package main

import "fmt"

// todo display gajets
type Phone struct {
	prize  string
	brand  string
	RAM    string
	IOS    string
	store  string
	camera string
}

// Phone
func (obj *Phone) Display() {
	fmt.Printf("Phone ==> prize=%v brand =%v, Ram =%v,IOS = %v, store=%v, camera=%v \n", obj.prize, obj.brand, obj.RAM, obj.IOS, obj.store, obj.camera)
}

type Laptop struct {
	brand     string
	prize     string
	RAM       string
	videocard string
	CPU       string
	store     string
	display   string
}

// Notebook
func (obj *Laptop) Display() {
	fmt.Printf("Laptop ==>  brand =%v, prize =%v,RAM = %v, videocard=%v, CPU=%v, Strore=%v,Display=%v\n", obj.brand, obj.prize, obj.RAM, obj.videocard, obj.CPU, obj.store, obj.display)
}

type Book struct {
	Name           string
	genres         string
	IBSN           string
	published_year string
	prize          string
}

// Book
func (obj *Book) Display() {
	fmt.Printf("Book ==> Name =%v, geners =%v,IbSN = %v, published_year=%v, prize=%v\n", obj.Name, obj.genres, obj.IBSN, obj.published_year, obj.prize)
}

type Car struct {
	brand string
	power string
	color string
	year  string
	prize string
}

// Car
func (obj *Car) Display() {
	fmt.Printf("Car ==> brand =%v, power =%v,color = %v, year=%v, prize=%v\n", obj.brand, obj.power, obj.color, obj.year, obj.prize)
}

type Items interface {
	Display()
}

func main() {
	var phone Items = &Phone{"1300$", "IPhone", "16GB", "IOS", "256GB", "128Px"}
	phone.Display()

	var laptop Items = &Laptop{"Hp", "1200$", "8Gb", "getforce 1650", "intel core7.12", "512 GB SSD", "15,6 FHD"}
	laptop.Display()

	var book Items = &Book{"Habit", "Biology", "100004557899", "2005", "50$"}
	book.Display()

	var car Items = &Car{"BMV", "750h", "red", "2015", "12000$"}
	car.Display()

}
