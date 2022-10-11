package domain

type Ticket struct{
	Id int `json:"Id"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	Country string `json:"Country"`
	Date string `json:"Date"`
	Time string `json:"Time"`
	Price string `json:"Price"`
}	
