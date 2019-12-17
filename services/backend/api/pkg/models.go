package pkg

import (
	"time"
)

// File type, represents the response of uploading a file.
// type File struct {
// 	ID      int    `json:"id"`
// 	Name    string `json:"name"`
// 	Content string `json:"content"`
// }

// UploadFile type, represents the request for uploading a file with certain payload.
// type UploadFile struct {
// 	ID   int                 `json:"id"`
// 	File graphqlllvar.Upload `json:"file"`
// }

// Goal ...
type Goal struct {
	ID           string    `json:"id"`
	Creator      string    `json:"author"`
	Participants []string  `json:"managers"`
	Likes        []string  `json:"likes"`
	Watchers     []string  `json:"watchers"`
	Aim          string    `json:"aim"`
	Reason       string    `json:"reason"`
	Details      string    `json:"details"`
	Inspiration  string    `json:"inspiration"`
	Type         string    `json:"type"`
	Tags         []string  `json:"tags"`
	SimilarGoals []string  `json:"similarGoals"`
	Journey      Journey   `json:"journey"`
	IsAchieved   bool      `json:"isAchieved"`
	IsPrivate    bool      `json:"isPrivate"`
	Date         time.Time `json:"timestamp"`
}

// Journey ...
type Journey struct {
	ID           string    `json:"id"`
	Details      string    `json:"details"`
	IsComplete   bool      `json:"isComplete"`
	IsInProgress bool      `json:"isInProgress"`
	IsStarted    bool      `json:"isStarted"`
	Type         string    `json:"type"`
	Steps        []string  `json:"steps"`
	Progress     int32     `json:"progress"`
	StartDate    time.Time `json:"startdate"`
	DueDate      time.Time `json:"duedate"`
}

// Group ...
type Group struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	People      []string  `json:"people"`
	Details     string    `json:"details"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Date        time.Time `json:"timestamp"`
}

// Match ...
type Match struct {
	ID           string    `json:"id"`
	PersonID     string    `json:"person_id"`
	Details      string    `json:"details"`
	Description  string    `json:"description"`
	Type         string    `json:"type"`
	Similarities []string  `json:"similarities"`
	Date         time.Time `json:"timestamp"`
}

// MatchedUser ...
type MatchedUser struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Status       string `json:"status"`
	Type         string `json:"account_type"`
	EmailAddress string `json:"emailaddress"`
	Gender       string `json:"gender"`
}

// // WorldEvent ...
// type WorldEvent struct {
// 	ID           string    `json:"id"`
// 	Name         string    `json:"name"`
// 	Type         string    `json:"type"`
// 	Size         int64     `json:"size"`
// 	Width        int       `json:"width"`
// 	Height       int       `json:"height"`
// 	Details      string    `json:"details"`
// 	IsComplete   bool      `json:"isComplete"`
// 	IsInProgress bool      `json:"isInProgress"`
// 	IsStarted    bool      `json:"isStarted"`
// 	Steps        []string  `json:"steps"`
// 	Progress     int       `json:"progress"`
// 	StartDate    time.Time `json:"startdate"`
// 	DueDate      time.Time `json:"duedate"`
// }

// UserMessage ...
type UserMessage struct {
	ID         string    `json:"_id"`
	Sender     string    `json:"sender"`
	Receiver   string    `json:"receiver"`
	Type       string    `json:"type"`
	Text       string    `json:"text"`
	IsSeen     bool      `json:"is_seen"`
	IsSent     bool      `json:"is_sent"`
	IsReceived bool      `json:"is_received"`
	Timestamp  time.Time `json:"timestamp"`
}

// Chat ...is the same as chatroom
type Chat struct {
	ID        string                       `json:"_id"`
	Messages  []UserMessage                `json:"messages"`
	Observers map[string]chan *UserMessage `json:"observers"`
	StartDate time.Time                    `json:"startdate"`
	EndDate   time.Time                    `json:"enddate"`
}

// Location ...
type Location struct {
	ID        string  `json:"id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Profile ...
type Profile struct {
	ID              string            `json:"id"`
	UserID          string            `json:"userID"`
	Level           string            `json:"level"`
	Rings           int32             `json:"rings"`
	About           string            `json:"about"`
	ProfileImage    Image             `json:"profileImage"`
	BackgroundImage Image             `json:"backgroundImage"`
	Followers       []string          `json:"followers"`
	Following       []string          `json:"following"`
	SocialProfiles  map[string]string `json:"socialprofiles"`
	DateLastUpdated time.Time         `json:"dateLastUpdated"`
}

// Space ...
type Space struct {
	ID          string    `json:"id"`
	Creator     string    `json:"author"`
	Managers    []string  `json:"managers"`
	Topic       string    `json:"topic"`
	Details     string    `json:"details"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	Date        time.Time `json:"timestamp"`
	Followers   []string  `json:"followers"`
}

// User ...
type User struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	Username          string     `json:"username"`
	Password          string     `json:"paswword"`
	Firstname         string     `json:"firstname"`
	Lastname          string     `json:"lastname"`
	Middlename        string     `json:"middlename"`
	Status            string     `json:"status"`
	Type              string     `json:"account_type"`
	EmailAddress      string     `json:"emailaddress"`
	Gender            string     `json:"gender"`
	BirthDate         time.Time  `json:"birthdate"`
	DateJoined        time.Time  `json:"datejoined"`
	MobilePhoneNumber Contact    `json:"mobile_phone_number"`
	BillingAddress    Address    `json:"billingAddress"`
	MailingAddress    Address    `json:"mailingAddress"`
	Languages         []Language `json:"languages"`
}

// Contact is the data structure
// for storing contact information
// such as phone numbers (Contact)
type Contact struct {
	LineNumber  string `json:"linenumber"`
	CountryCode string `json:"countrycode"`
	AreaCode    string `json:"areacode"`
	Premfix     string `json:"prefix"`
}

//Address is the address of a user
type Address struct {
	StreetAddressLine1 string `json:"streetAddress1"`
	StreetAddressLine2 string `json:"streetAddress2"`
	PostalCode         string `json:"postalCode"`
	Province           string `json:"province"`
	Country            Country
	State              State
	City               City
}

// Country struct
type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// City struct
type City struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// State struct
type State struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// Language is the name of a user's
// language spoken
type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// Login ...
type Login struct {
	UserName string `json:"username"`
	Password string `json:"paswword"`
}

// Register ...
type Register struct {
	UserName     string `json:"username"`
	Password     string `json:"paswword"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	EmailAddress string `json:"emailAddress"`
	Gender       string `json:"gender"`
}

// UserPost ...
type UserPost struct {
	ID           string    `json:"id" `
	Author       string    `json:"author"`
	Anonymous    bool      `json:"anonymous"`
	Topic        string    `json:"topic"`
	Category     string    `json:"category"`
	ContentText  string    `json:"contentText"`
	Type         string    `json:"type"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Date         time.Time `json:"timestamp"`
	ContentPhoto Image     `json:"contentPhoto"`
	ContentFile  File      `json:"contentFile"`
	Likes        []string  `json:"likes"`
	Agreements   []string  `json:"agreements"`
	Followers    []string  `json:"followers"`
	Comments     []Comment `json:"comments"`
	Shares       []Share   `json:"shares"`
	// File         UploadFile `json:"file"`
}

// Image ...
type Image struct {
	// ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Size   int64  `json:"size"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// File ...
type File struct {
	// ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Size   int64  `json:"size"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Comment ...
type Comment struct {
	ID     string `json:"id"`
	UserID string `json:"userid"`
	Text   string `json:"text"`
}

// Share ...
type Share struct {
	Party string `json:"party"`
}

// Task ...
type Task struct {
	ID   uint32 `json:"id"`
	Text string `json:"text"`
}
