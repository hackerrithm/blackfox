//go:generate protoc ./catalogue.proto --go_out=plugins=grpc:./pb
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

	"github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/model"
	userProfile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
)

type grpcServer struct {
	catalogueService engine.Catalogue
	productService   engine.Product
	profile          *userProfile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Catalogue, userURL string, port int) error {
	profileClient, err := userProfile.NewClient(userURL)
	if err != nil {
		profileClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		profileClient.Close()
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterCatalogueServiceServer(serv, &grpcServer{s, profileClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostCatalogue(ctx context.Context, r *pb.PostCatalogueRequest) (*pb.PostCatalogueResponse, error) {
	err := s.catalogueService.Insert(ctx, r.Name, r.Details, r.Description, r.Type, r.ProductIDs, r.Tags)
	if err != nil {
		log.Println("here error in catalogue method")
		return nil, err
	}
	return &pb.PostCatalogueResponse{
		Catalogue: "inserted",
	}, nil
}

func (s *grpcServer) PutCatalogue(ctx context.Context, r *pb.PutCatalogueRequest) (*pb.PutCatalogueResponse, error) {
	err := s.catalogueService.Update(ctx, r.Id, r.Name, r.Details, r.Description, r.Type, r.ProductIDs, r.Tags)
	if err != nil {
		return nil, err
	}
	return &pb.PutCatalogueResponse{
		Catalogue: "updated",
	}, nil
}

func (s *grpcServer) GetCatalogue(ctx context.Context, r *pb.GetCatalogueRequest) (*pb.GetCatalogueResponse, error) {
	a, err := s.catalogueService.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	var prodIDs []string
	for _, m := range a.Products {
		prodIDs = append(prodIDs, m.Hex())
	}

	return &pb.GetCatalogueResponse{
		Catalogue: &pb.Catalogue{
			Id:          a.ID.Hex(),
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			ProductIDs:  prodIDs,
			Tags:        a.Tags,
		},
	}, nil
}

func (s *grpcServer) GetMultipleCatalogues(ctx context.Context, r *pb.GetMultipleCataloguesRequest) (*pb.GetMultipleCataloguesResponse, error) {
	log.Println("got to grpc ")
	var prodIDs []string

	res, err := s.catalogueService.ListAllCatalogues(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	catalogues := []*pb.Catalogue{}
	for _, p := range *res {
		for _, m := range p.Products {
			prodIDs = append(prodIDs, m.Hex())
		}
		catalogues = append(
			catalogues,
			&pb.Catalogue{
				Id:          p.ID.Hex(),
				Details:     p.Details,
				Description: p.Description,
				Type:        p.Type,
				ProductIDs:  prodIDs,
				Tags:        p.Tags,
			},
		)
	}

	return &pb.GetMultipleCataloguesResponse{
		Catalogues: catalogues,
	}, nil
}

func (s *grpcServer) DeleteCatalogue(ctx context.Context, r *pb.DeleteCatalogueRequest) (*pb.DeleteCatalogueResponse, error) {
	a, err := s.catalogueService.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCatalogueResponse{
		Id: a,
	}, nil
}

func (s *grpcServer) PostProduct(ctx context.Context, r *pb.PostProductRequest) (*pb.PostProductResponse, error) {
	err := s.productService.InsertProduct(ctx, r.Name, r.Details, r.Description, r.Type, r.Tags, r.Price, r.Discount)
	if err != nil {
		log.Println("here error in catalogue method")
		return nil, err
	}
	return &pb.PostProductResponse{
		Product: "inserted",
	}, nil
}

func (s *grpcServer) PutProduct(ctx context.Context, r *pb.PutProductRequest) (*pb.PutProductResponse, error) {
	err := s.productService.UpdateProduct(ctx, r.Id, r.Name, r.Details, r.Description, r.Type, r.Tags, r.Price, r.Discount)
	if err != nil {
		return nil, err
	}
	return &pb.PutProductResponse{
		Product: "updated",
	}, nil
}

func (s *grpcServer) GetProduct(ctx context.Context, r *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	a, err := s.productService.FindOneProduct(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:          a.ID.Hex(),
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Price:       a.Price,
			Discount:    r.Discount,
			Tags:        a.Tags,
		},
	}, nil
}

func (s *grpcServer) GetMultipleProducts(ctx context.Context, r *pb.GetMultipleCataloguesRequest) (*pb.GetMultipleCataloguesResponse, error) {
	log.Println("got to grpc ")

	res, err := s.productService.ListAllProducts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	products := []*pb.Product{}
	for _, p := range *res {
		products = append(
			product,
			&pb.Product{
				Id:          p.ID.Hex(),
				Name:        p.Name,
				Details:     p.Details,
				Description: p.Description,
				Type:        p.Type,
				Price:       p.Price,
				Discount:    p.Discount,
				Tags:        p.Tags,
			},
		)
	}

	return &pb.GetMultipleProductsResponse{
		Products: products,
	}, nil
}

func (s *grpcServer) DeleteProduct(ctx context.Context, r *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	a, err := s.productService.RemoveProduct(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProductResponse{
		Id: a,
	}, nil
}
