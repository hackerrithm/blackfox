//go:generate protoc ./order.proto --go_out=plugins=grpc:./pb
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

package adapters

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	catalogue "github.com/hackerrithm/blackfox/services/backend/catalogue/cmd/catalogue/client"
	"github.com/hackerrithm/blackfox/services/backend/order/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/order/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/order/pkg/model"
	user "github.com/hackerrithm/blackfox/services/backend/user/cmd/user/client"
)

type grpcServer struct {
	service   engine.Order
	user      *user.Client
	catalogue *catalogue.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Order, userURL, catalogURL string, port int) error {

	userClient, err := user.NewClient(userURL)
	if err != nil {
		userClient.Close()
		return err
	}

	catalogClient, err := catalogue.NewClient(catalogURL)
	if err != nil {
		userClient.Close()
		catalogClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		userClient.Close()
		catalogClient.Close()
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &grpcServer{s, userClient, catalogClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostOrder(ctx context.Context, r *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	// Check if account exists
	_, err := s.user.GetUser(ctx, r.UserID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Get ordered products
	productIDs := []string{}
	for _, p := range r.Products {
		productIDs = append(productIDs, p.ProductId)
	}
	orderedProducts, err := s.catalogue.GetMultipleProducts(ctx, 0, 0)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Construct products
	products := []domain.OrderedProduct{}
	for _, p := range orderedProducts {
		product := domain.OrderedProduct{
			ID:          p.ID,
			Quantity:    0,
			Price:       p.Price,
			Name:        p.Name,
			Description: p.Description,
		}
		for _, rp := range r.Products {
			if rp.ProductId == p.ID.Hex() {
				product.Quantity = rp.Quantity
				break
			}
		}

		if product.Quantity != 0 {
			products = append(products, product)
		}
	}

	order, err := s.service.Insert(ctx, r.UserID, "", "", "", 0.0, products)
	if err != nil {
		log.Println("here error in order method")
		return nil, err
	}

	// Make response order
	orderedProds := &pb.Order{
		Id:              order.ID.Hex(),
		UserID:          order.UserID.Hex(),
		TotalPrice:      order.TotalPrice,
		OrderedProducts: []*pb.OrderedProduct{},
	}
	// orderProto.CreatedAt, _ = order.CreatedAt.MarshalBinary()
	for _, p := range order.Products {
		orderedProds.OrderedProducts = append(orderedProds.OrderedProducts, &pb.OrderedProduct{
			Id:          p.ID.Hex(),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
		})
	}
	return &pb.PostOrderResponse{
		Order: orderedProds,
	}, nil
}

func (s *grpcServer) GetOrder(ctx context.Context, r *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// a, err := s.service.FindOne(ctx, r.Id)
	// if err != nil {
	// 	return nil, err
	// }

	// var mngs, fllwrs []string
	// for _, m := range a.Managers {
	// 	mngs = append(mngs, m.Hex())
	// }
	// for _, f := range a.Followers {
	// 	fllwrs = append(fllwrs, f.Hex())
	// }

	return &pb.GetOrderResponse{
		Order: &pb.Order{},
	}, nil
}

func (s *grpcServer) GetMultipleOrders(ctx context.Context, r *pb.GetMultipleOrdersRequest) (*pb.GetMultipleOrdersResponse, error) {
	log.Println("got to grpc ")

	userAccountOrders, err := s.service.ListAllOrders(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}

	// Get all ordered products
	productIDMap := map[string]bool{}
	for _, o := range *userAccountOrders {
		for _, p := range o.Products {
			productIDMap[p.ID.Hex()] = true
		}
	}
	productIDs := []string{}
	for id := range productIDMap {
		productIDs = append(productIDs, id)
	}

	products, err := s.catalogue.GetMultipleProducts(ctx, 0, 0)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Construct orders
	orders := []*pb.Order{}
	for _, o := range *userAccountOrders {
		// Encode order
		op := &pb.Order{
			UserID:          o.UserID.Hex(),
			Id:              o.ID.Hex(),
			TotalPrice:      o.TotalPrice,
			OrderedProducts: []*pb.OrderedProduct{},
		}
		// op.CreatedAt, _ = o.CreatedAt.MarshalBinary()

		// Decorate orders with products
		for _, product := range o.Products {
			// Populate product fields
			for _, p := range products {
				if p.ID == product.ID {
					product.Name = p.Name
					product.Description = p.Description
					product.Price = p.Price
					break
				}
			}

			op.OrderedProducts = append(op.OrderedProducts, &pb.OrderedProduct{
				Id:          product.ID.Hex(),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Quantity:    product.Quantity,
			})
		}

		orders = append(orders, op)
	}

	return &pb.GetMultipleOrdersResponse{
		Orders: orders,
	}, nil
}

func (s *grpcServer) DeleteOrder(ctx context.Context, r *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteOrderResponse{
		Id: a,
	}, nil
}
