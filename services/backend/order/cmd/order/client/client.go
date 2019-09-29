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

package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/order/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/order/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.OrderServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewOrderServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, userID string, products []domain.OrderedProduct) (domain.Order, error) {
	protoProducts := []*pb.PostOrderRequest_OrderProduct{}
	for _, p := range products {
		protoProducts = append(protoProducts, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ID.Hex(),
			Quantity:  p.Quantity,
		})
	}

	r, err := c.service.PostOrder(
		ctx,
		&pb.PostOrderRequest{},
	)
	if err != nil {
		return domain.Order{}, err
	}

	// Create response order
	newOrder := r.Order
	// newOrderCreatedAt := time.Time{}
	// newOrderCreatedAt.UnmarshalBinary(newOrder.CreatedOn.)

	log.Println(r.Order)
	return domain.Order{
		Description: newOrder.Description,
		Details:     newOrder.Details,
		Products:    products,
		UserID:      bson.ObjectIdHex(newOrder.UserID),
		TotalPrice:  newOrder.TotalPrice,
		Type:        newOrder.Type,
	}, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, creator, topic, details, description, typ string, managers, followers, tags []string) (string, error) {
	// r, err := c.service.PutOrder(
	// 	ctx,
	// 	&pb.PutOrderRequest{
	// 		Id:          id,
	// 		Creator:     creator,
	// 		Topic:       topic,
	// 		Details:     details,
	// 		Description: description,
	// 		Type:        typ,
	// 		Followers:   followers,
	// 		Managers:    managers,
	// 		Tags:        tags,
	// 	},
	// )
	// if err != nil {
	// 	return "", err
	// }
	return "", nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Order, error) {
	// r, err := c.service.GetOrder(
	// 	ctx,
	// 	&pb.GetOrderRequest{Id: id, UserID: userID},
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// var mngs, fllwrs []bson.ObjectId
	// for _, m := range r.Order.Managers {
	// 	mngs = append(mngs, bson.ObjectIdHex(m))
	// }
	// for _, f := range r.Order.Followers {
	// 	fllwrs = append(fllwrs, bson.ObjectIdHex(f))
	// }

	return &domain.Order{
		// ID:          bson.ObjectIdHex(r.Order.Id),
		// Creator:     bson.ObjectIdHex(r.Order.Creator),
		// Topic:       r.Order.Topic,
		// Details:     r.Order.Details,
		// Description: r.Order.Description,
		// Type:        r.Order.Type,
		// Followers:   fllwrs,
		// Managers:    mngs,
		// Tags:        r.Order.Tags,
	}, nil
}

// GetMultiple is used to get the list of specified orders
func (c *Client) GetMultiple(ctx context.Context, userID string, skip uint64, take uint64) ([]domain.Order, error) {
	r, err := c.service.GetMultipleOrders(
		ctx,
		&pb.GetMultipleOrdersRequest{
			Skip:   skip,
			Take:   take,
			UserID: userID,
		},
	)
	if err != nil {
		return nil, err
	}

	// Create response orders
	orders := []domain.Order{}
	for _, orderProto := range r.Orders {
		newOrder := domain.Order{
			ID:         bson.ObjectIdHex(orderProto.Id),
			TotalPrice: orderProto.TotalPrice,
			UserID:     bson.ObjectIdHex(orderProto.UserID),
		}
		// newOrder.CreatedAt = time.Time{}
		// newOrder.CreatedAt.UnmarshalBinary(orderProto.CreatedAt)

		products := []domain.OrderedProduct{}
		for _, p := range orderProto.OrderedProducts {
			products = append(products, domain.OrderedProduct{
				ID:          bson.ObjectIdHex(p.Id),
				Quantity:    p.Quantity,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			})
		}
		newOrder.Products = products

		orders = append(orders, newOrder)
	}

	// orders := []domain.Order{}
	// for _, a := range r.Orders {
	// 	for _, m := range a.Managers {
	// 		mngs = append(mngs, bson.ObjectIdHex(m))
	// 	}
	// 	for _, f := range a.Followers {
	// 		fllwrs = append(fllwrs, bson.ObjectIdHex(f))
	// 	}

	// 	orders = append(orders, domain.Order{
	// 		ID:          bson.ObjectIdHex(a.Id),
	// 		Creator:     bson.ObjectIdHex(a.Creator),
	// 		Topic:       a.Topic,
	// 		Details:     a.Details,
	// 		Description: a.Description,
	// 		Type:        a.Type,
	// 		Followers:   fllwrs,
	// 		Managers:    mngs,
	// 		Tags:        a.Tags,
	// 	})
	// }
	return orders, nil
}

// Delete removes a order with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteOrder(
		ctx,
		&pb.DeleteOrderRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
