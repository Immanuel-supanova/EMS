package models

import "time"

type Employee struct {
	Employee_Id uint `gorm:"primaryKey"`
	Full_Name   string
	Password    string `gorm:"unique"`
	Job_Code    uint
	Contact_Id  uint
	Profile_Id  uint
	Wages       float64
	Is_Active   *bool
	Enquiries   []Enquiry
	Raises      []Raise
	Leaves      []Leave
	Advances    []Advance
}

type Job struct {
	Job_Code  uint `gorm:"primaryKey"`
	Job_Name  string
	Employees []Employee `gorm:"foreignKey:Job_Code"`
}

type Contact struct {
	Contact_Id     uint   `gorm:"primaryKey"`
	Contact_Number string `gorm:"unique"`
	Email          string `gorm:"unique"`
	Country        string
	Address        string `gorm:"unique"`
	City           string
	Postal_code    string
	Employee       Employee
}

type Profile struct {
	Profile_Id uint `gorm:"primaryKey"`
	Image      *string
	Bio        string
	Employee   Employee
}

type Enquiry struct {
	Enquiry_Id  uint `gorm:"primaryKey"`
	Employee_Id uint
	Question_Id uint
	Answer      string
}

type Question struct {
	Question_Id uint `gorm:"primaryKey"`
	Question    string
	Enquiry     Enquiry
}

type Raise struct {
	Raise_Id    uint `gorm:"primaryKey"`
	Employee_Id uint
	Reason      string
	Amount      float64
	Approval    *bool
}

type Leave struct {
	Leave_Id    uint `gorm:"primaryKey"`
	Employee_Id uint
	Reason      string
	Start_Date  time.Time
	End_Date    time.Time
	Approval    *bool
}

type Advance struct {
	Advanced_Id    uint `gorm:"primaryKey"`
	Employee_Id    uint
	Reason         string
	Amount         float64
	Payment_Period string
	Approval       *bool
}
