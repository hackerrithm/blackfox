// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type Address struct {
	StreetAddressLine1 string   `json:"streetAddressLine1"`
	StreetAddressLine2 string   `json:"streetAddressLine2"`
	PostalCode         string   `json:"postalCode"`
	Province           string   `json:"province"`
	Country            *Country `json:"country"`
	State              *State   `json:"state"`
	City               *City    `json:"city"`
}

type Chat struct {
	ID       string         `json:"id"`
	Messages []*UserMessage `json:"messages"`
}

type City struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Comment struct {
	ID     string  `json:"id"`
	UserID *string `json:"userID"`
	Text   *string `json:"text"`
}

type Contact struct {
	LineNumber  string `json:"lineNumber"`
	CountryCode string `json:"countryCode"`
	AreaCode    string `json:"areaCode"`
	Premfix     string `json:"premfix"`
}

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type File struct {
	Name   string  `json:"name"`
	Type   *string `json:"type"`
	Size   *int    `json:"size"`
	Width  *int    `json:"width"`
	Height *int    `json:"height"`
}

type Goal struct {
	ID           string   `json:"id"`
	Creator      string   `json:"creator"`
	Aim          string   `json:"aim"`
	Participants []string `json:"participants"`
	Likes        []string `json:"likes"`
	Watchers     []string `json:"watchers"`
	Reason       string   `json:"reason"`
	Inspiration  string   `json:"inspiration"`
	Details      string   `json:"details"`
	Type         string   `json:"type"`
	Tags         []string `json:"tags"`
	SimilarGoals []string `json:"similarGoals"`
	IsAchieved   *bool    `json:"isAchieved"`
	IsPrivate    *bool    `json:"isPrivate"`
	Journey      *Journey `json:"journey"`
}

type Group struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Details     string   `json:"details"`
	Description *string  `json:"description"`
	Type        *string  `json:"type"`
	People      []string `json:"people"`
}

type Image struct {
	Name   string  `json:"name"`
	Type   *string `json:"type"`
	Size   *int    `json:"size"`
	Width  *int    `json:"width"`
	Height *int    `json:"height"`
}

type Journey struct {
	Details      string   `json:"details"`
	IsComplete   *bool    `json:"isComplete"`
	IsInProgress *bool    `json:"isInProgress"`
	IsStarted    *bool    `json:"isStarted"`
	Type         string   `json:"type"`
	Steps        []string `json:"steps"`
	Progress     *int     `json:"progress"`
}

type JourneyInput struct {
	Details      string   `json:"details"`
	IsComplete   bool     `json:"isComplete"`
	IsInProgress bool     `json:"isInProgress"`
	IsStarted    bool     `json:"isStarted"`
	Type         string   `json:"type"`
	Steps        []string `json:"steps"`
	Progress     *int     `json:"progress"`
}

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Location struct {
	ID        string  `json:"id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type LocationInput struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MatchedUser struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Status       string `json:"status"`
	Type         string `json:"type"`
	EmailAddress string `json:"emailAddress"`
	Gender       string `json:"gender"`
}

type PaginationInput struct {
	Skip *int `json:"skip"`
	Take *int `json:"take"`
}

type PostGoalInput struct {
	Creator      string        `json:"creator"`
	Aim          string        `json:"aim"`
	Participants []string      `json:"participants"`
	Likes        []string      `json:"likes"`
	Watchers     []string      `json:"watchers"`
	Reason       string        `json:"reason"`
	Inspiration  string        `json:"inspiration"`
	Details      string        `json:"details"`
	Type         string        `json:"type"`
	Tags         []string      `json:"tags"`
	SimilarGoals []string      `json:"similarGoals"`
	IsAchieved   bool          `json:"isAchieved"`
	IsPrivate    bool          `json:"isPrivate"`
	Journey      *JourneyInput `json:"journey"`
}

type PostGroupInput struct {
	Title       string   `json:"title"`
	Details     string   `json:"details"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	People      []string `json:"people"`
}

type PostPostInput struct {
	Author           string  `json:"author"`
	Anonymous        *bool   `json:"anonymous"`
	Topic            string  `json:"topic"`
	Category         string  `json:"category"`
	ContentText      string  `json:"contentText"`
	Type             string  `json:"type"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	ContentPhotoName string  `json:"contentPhotoName"`
	ContentFileName  string  `json:"contentFileName"`
}

type PostSpaceInput struct {
	Creator     string     `json:"creator"`
	Topic       string     `json:"topic"`
	Details     string     `json:"details"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Tags        []string   `json:"tags"`
	Date        *time.Time `json:"date"`
	Managers    []string   `json:"managers"`
	Followers   []string   `json:"followers"`
}

type PostTaskInput struct {
	Text string     `json:"text"`
	Date *time.Time `json:"date"`
}

type Profile struct {
	ID              string     `json:"id"`
	UserID          string     `json:"userID"`
	Level           *string    `json:"level"`
	Rings           *int       `json:"rings"`
	About           string     `json:"about"`
	ProfileImage    *Image     `json:"profileImage"`
	BackgroundImage *Image     `json:"backgroundImage"`
	Followers       []string   `json:"followers"`
	Following       []string   `json:"following"`
	DateLastUpdated *time.Time `json:"dateLastUpdated"`
}

type RegisterInput struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	EmailAddress string `json:"emailAddress"`
	Gender       string `json:"gender"`
}

type Share struct {
	Party string `json:"party"`
}

type Space struct {
	ID          int        `json:"id"`
	Creator     *string    `json:"creator"`
	Topic       *string    `json:"topic"`
	Details     *string    `json:"details"`
	Description *string    `json:"description"`
	Type        *string    `json:"type"`
	Tags        []string   `json:"tags"`
	Date        *time.Time `json:"date"`
	Followers   []string   `json:"followers"`
	Managers    []string   `json:"managers"`
}

type State struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Task struct {
	ID   *string `json:"id"`
	Text string  `json:"text"`
}

type UpdateGoalInput struct {
	ID           string        `json:"id"`
	Creator      string        `json:"creator"`
	Aim          string        `json:"aim"`
	Participants []string      `json:"participants"`
	Likes        []string      `json:"likes"`
	Watchers     []string      `json:"watchers"`
	Reason       string        `json:"reason"`
	Inspiration  string        `json:"inspiration"`
	Details      string        `json:"details"`
	Type         string        `json:"type"`
	Tags         []string      `json:"tags"`
	SimilarGoals []string      `json:"similarGoals"`
	IsAchieved   bool          `json:"isAchieved"`
	IsPrivate    bool          `json:"isPrivate"`
	Journey      *JourneyInput `json:"journey"`
}

type UpdateGroupInput struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Details     string   `json:"details"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	People      []string `json:"people"`
}

type UpdatePostInput struct {
	ID               string  `json:"id"`
	Author           string  `json:"author"`
	Anonymous        *bool   `json:"anonymous"`
	Topic            string  `json:"topic"`
	Category         string  `json:"category"`
	ContentText      string  `json:"contentText"`
	Type             string  `json:"type"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	ContentPhotoName string  `json:"contentPhotoName"`
	ContentFileName  string  `json:"contentFileName"`
}

type UpdateSpaceInput struct {
	ID          int        `json:"id"`
	Creator     string     `json:"creator"`
	Topic       string     `json:"topic"`
	Details     string     `json:"details"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Tags        []string   `json:"tags"`
	Date        *time.Time `json:"date"`
	Managers    []string   `json:"managers"`
	Followers   []string   `json:"followers"`
}

type UpdateTaskInput struct {
	ID   string     `json:"id"`
	Text string     `json:"text"`
	Date *time.Time `json:"date"`
}

type User struct {
	ID                *string     `json:"id"`
	Name              *string     `json:"name"`
	Username          *string     `json:"username"`
	Password          *string     `json:"password"`
	Firstname         *string     `json:"firstname"`
	Lastname          *string     `json:"lastname"`
	Middlename        *string     `json:"middlename"`
	Status            *string     `json:"status"`
	Type              *string     `json:"type"`
	EmailAddress      *string     `json:"emailAddress"`
	Gender            *string     `json:"gender"`
	BirthDate         *time.Time  `json:"birthDate"`
	DateJoined        *time.Time  `json:"dateJoined"`
	MobilePhoneNumber *Contact    `json:"mobilePhoneNumber"`
	BillingAddress    *Address    `json:"billingAddress"`
	MailingAddress    *Address    `json:"mailingAddress"`
	Languages         []*Language `json:"languages"`
}

type UserMessage struct {
	ID         string    `json:"id"`
	Sender     string    `json:"sender"`
	Receiver   string    `json:"receiver"`
	Type       string    `json:"type"`
	Text       string    `json:"text"`
	Timestamp  time.Time `json:"timestamp"`
	IsSeen     bool      `json:"isSeen"`
	IsSent     bool      `json:"isSent"`
	IsReceived bool      `json:"isReceived"`
}

type UserPost struct {
	ID           string     `json:"id"`
	Author       string     `json:"author"`
	Anonymous    *bool      `json:"anonymous"`
	Topic        string     `json:"topic"`
	Category     *string    `json:"category"`
	ContentText  string     `json:"contentText"`
	Type         string     `json:"type"`
	Latitude     *float64   `json:"latitude"`
	Longitude    *float64   `json:"longitude"`
	Date         *time.Time `json:"date"`
	ContentPhoto *Image     `json:"contentPhoto"`
	ContentFile  *File      `json:"contentFile"`
	Likes        []string   `json:"likes"`
	Agreements   []string   `json:"agreements"`
	Followers    []string   `json:"followers"`
	Comments     []*Comment `json:"comments"`
	Shares       []*Share   `json:"shares"`
}