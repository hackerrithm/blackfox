// // Copyright 2019 kemar
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //     http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.

// package mongodb

// import (
// 	"log"

// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// type (
// 	// ID represents the last used integer
// 	// id for any collection
// 	ID struct {
// 		Next int64 `bson:"n"`
// 	}
// )

// var (
// 	idCollection = "id"
// )

// // simple way of using integer IDs with MongoDB
// func getNextSequence(s *mgo.Session, name string) int64 {
// 	c := s.DB("test1").C(idCollection)
// 	change := mgo.Change{
// 		Update:    bson.M{"$inc": bson.M{"n": 1}},
// 		Upsert:    true,
// 		ReturnNew: true,
// 	}
// 	id := new(ID)
// 	log.Println("this is: ", id)
// 	c.Find(bson.M{"_id": name}).Apply(change, id)
// 	return id.Next
// }

// func toBSON(ID string) bson.ObjectId {
// 	return bson.ObjectIdHex(ID)
// }
