// Copyright 2019 kemar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package domain

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var (
	userContextKey contextKey = "user"
	now                       = defaultNow
)

type (
	contextKey string

	// User is a user object
	// that has datatypes for
	// the user domain (User)
	User struct {
		ID                bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name              string        `json:"name" bson:"name,omitempty"`
		Username          string        `json:"username" bson:"username,omitempty"`
		Password          string        `json:"password" bson:"password,omitempty"`
		Firstname         string        `json:"firstname" bson:"firstname,omitempty"`
		Lastname          string        `json:"lastname" bson:"lastname,omitempty"`
		Middlename        string        `json:"middlename" bson:"middlename,omitempty"`
		Status            string        `json:"status" bson:"status,omitempty"`
		Type              string        `json:"account_type" bson:"account_type,omitempty"`
		EmailAddress      string        `json:"emailaddress" bson:"email_address,omitempty"`
		Gender            string        `json:"gender" bson:"gender,omitempty"`
		BirthDate         time.Time     `json:"birthdate" bson:"birthdate,omitempty"`
		DateJoined        time.Time     `json:"datejoined" bson:"datejoined,omitempty"`
		MobilePhoneNumber Contact       `json:"mobile_phone_number" bson:"mobile_phone_number,omitempty"`
		BillingAddress    Address       `json:"billingAddress" bson:"billing_address,omitempty"`
		MailingAddress    Address       `json:"mailingAddress" bson:"mailing_address,omitempty"`
		Languages         []Language    `json:"languages" bson:"languages,omitempty"`
	}

	// Contact is the data structure
	// for storing contact information
	// such as phone numbers (Contact)
	Contact struct {
		LineNumber  string `json:"linenumber" bson:"line_number,omitempty"`
		CountryCode string `json:"countrycode" bson:"country_code,omitempty"`
		AreaCode    string `json:"areacode" bson:"areacode,omitempty"`
		Prefix      string `json:"prefix" bson:"prefix,omitempty"`
	}

	// Address is the address of a user
	Address struct {
		StreetAddressLine1 string `json:"streetAddress1" bson:"street_address_1,omitempty"`
		StreetAddressLine2 string `json:"streetAddress2" bson:"street_address_2,omitempty"`
		PostalCode         string `json:"postalCode" bson:"postal_code,omitempty"`
		Province           string `json:"province" bson:"province,omitempty"`
		Country            Country
		State              State
		City               City
	}

	// Country struct
	Country struct {
		Name string `json:"name" bson:"name,omitempty"`
		Code string `json:"code" bson:"code,omitempty"`
	}

	// City struct
	City struct {
		Name string `json:"name" bson:"name,omitempty"`
		Code string `json:"code" bson:"code,omitempty"`
	}

	// State struct
	State struct {
		Name string `json:"name" bson:"name,omitempty"`
		Code string `json:"code" bson:"code,omitempty"`
	}

	// Language is the name of a user's
	// language spoken
	Language struct {
		Name string `json:"name" bson:"name,omitempty"`
		Code string `json:"code" bson:"code,omitempty"`
	}

	// Login ...
	Login struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	// Register ...
	Register struct {
		UserName     string    `json:"username"`
		Password     string    `json:"password"`
		FirstName    string    `json:"firstname"`
		LastName     string    `json:"lastname"`
		EmailAddress string    `json:"emailAddress"`
		Gender       string    `json:"gender"`
		Status       string    `json:"status"`
		BirthDate    time.Time `json:"birthdate"`
	}
)

// SetPassword sets user's password
func (u *User) SetPassword(p string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.Password = string(hashedPassword)
}

// GetPassword ...
func (u *User) GetPassword() []byte {
	return nil
}

// IsCredentialsVerified matches given password with user's password
func (u *User) IsCredentialsVerified(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// NewContext ...
func (u *User) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, userContextKey, u)
}

// UserFromContext gets user from context
func UserFromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userContextKey).(*User)
	return u, ok
}

// UserMustFromContext gets user from context. if can't make panic
func UserMustFromContext(ctx context.Context) *User {
	u, ok := ctx.Value(userContextKey).(*User)
	if !ok {
		panic("user can't get from request's context")
	}
	return u
}

// defaultNow used to get the
// current date time (UTC)
func defaultNow() time.Time {
	return time.Now().UTC()
}

// NewUser creates a new User!
func NewUser(name string, username, password, fname, lname, email string) *User {
	return &User{
		Name:         name,
		Firstname:    fname,
		Lastname:     lname,
		Username:     username,
		Password:     password,
		EmailAddress: email,
		DateJoined:   now(),
	}
}
