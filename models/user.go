package models

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	RoomNumber  string `json:"roomNumber"`
	PhoneNumber string `json:"phoneNumber"`
}

var Users = []User{
	{ID: "1", Name: "Vivek", PhoneNumber: "9100000000", RoomNumber: "101"},
	{ID: "2", Name: "Sadiq", PhoneNumber: "9100000001", RoomNumber: "102"},
}
