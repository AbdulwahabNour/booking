package models

import "time"




type User struct{
        ID int  `json:"id" db:"id"`  
        FirstName string 
        LastName string
        Email string
        Password string
        AccessLevel int
        CreatedeAt time.Time
        UpdatedAt time.Time
}

type Room  struct{
        ID int
        RoomName string
        RoomImage string
        CreatedeAt time.Time
        UpdatedAt time.Time
}
type Reservation struct{
        ID int `json:"id" db:"id"`
        FirstName string
        LastName string
        Email string
        Phone string
        StartDate time.Time
        EndDate time.Time
        RoomId int
        CreatedeAt time.Time
        UpdatedAt time.Time
} 

type Restriction struct{
        ID int  `json:"id" db:"id"`  
        CreatedeAt time.Time
        UpdatedAt time.Time
}
type RoomRestriction struct{
        ID int `json:"id" db:"id"`
        RoomID int 
        StartDate time.Time
        EndDate time.Time
        ReservationID int
        RestrictionID int
        CreatedeAt time.Time
        UpdatedAt time.Time
}