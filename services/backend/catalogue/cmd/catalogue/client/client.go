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

	"github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.CatalogueServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewCatalogueServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// PostCatalogue ...
func (c *Client) PostCatalogue(ctx context.Context, name, details, description, typ string, prods, tags []string) (string, error) {
	r, err := c.service.PostCatalogue(
		ctx,
		&pb.PostCatalogueRequest{
			Name:        name,
			Details:     details,
			Description: description,
			Type:        typ,
			ProductIDs:  prods,
			Tags:        tags,
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Catalogue)
	return r.Catalogue, nil
}

// PutCatalogue ...
func (c *Client) PutCatalogue(ctx context.Context, id, name, details, description, typ string, prods, tags []string) (string, error) {
	r, err := c.service.PutCatalogue(
		ctx,
		&pb.PutCatalogueRequest{
			Id:          id,
			Name:        name,
			Details:     details,
			Description: description,
			Type:        typ,
			ProductIDs:  prods,
			Tags:        tags,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Catalogue, nil
}

// GetCatalogue ...
func (c *Client) GetCatalogue(ctx context.Context, id string, userID uint64) (*domain.Catalogue, error) {
	r, err := c.service.GetCatalogue(
		ctx,
		&pb.GetCatalogueRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	var prodIDs []bson.ObjectId
	for _, m := range r.Catalogue.ProductIDs {
		prodIDs = append(prodIDs, bson.ObjectIdHex(m))
	}
	return &domain.Catalogue{
		ID:          bson.ObjectIdHex(r.Catalogue.Id),
		Name:        r.Catalogue.Name,
		Details:     r.Catalogue.Details,
		Description: r.Catalogue.Description,
		Type:        r.Catalogue.Type,
		Products:    prodIDs,
		Tags:        r.Catalogue.Tags,
	}, nil
}

// GetMultipleCatalogues is used to get the list of specified catalogues
func (c *Client) GetMultipleCatalogues(ctx context.Context, skip uint64, take uint64) ([]domain.Catalogue, error) {
	var prodIDs []bson.ObjectId

	r, err := c.service.GetMultipleCatalogues(
		ctx,
		&pb.GetMultipleCataloguesRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	catalogues := []domain.Catalogue{}
	for _, a := range r.Catalogues {

		for _, m := range a.ProductIDs {
			prodIDs = append(prodIDs, bson.ObjectIdHex(m))
		}

		catalogues = append(catalogues, domain.Catalogue{
			ID:          bson.ObjectIdHex(a.Id),
			Name:        a.Name,
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Products:    prodIDs,
			Tags:        a.Tags,
		})
	}
	return catalogues, nil
}

// DeleteCatalogue removes a catalogue with passed identifier
func (c *Client) DeleteCatalogue(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteCatalogue(
		ctx,
		&pb.DeleteCatalogueRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}

// PostProduct ...
func (c *Client) PostProduct(ctx context.Context, name, details, description, typ string, tags []string, price float64, discount float32) (string, error) {
	r, err := c.service.PostProduct(
		ctx,
		&pb.PostProductRequest{
			Name:        name,
			Details:     details,
			Description: description,
			Type:        typ,
			Price:       price,
			Discount:    discount,
			Tags:        tags,
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Product)
	return "", nil
}

// PutProduct ...
func (c *Client) PutProduct(ctx context.Context, id, name, details, description, typ string, tags []string, price float64, discount float32) (string, error) {
	r, err := c.service.PutProduct(
		ctx,
		&pb.PutProductRequest{
			Id:          id,
			Name:        name,
			Details:     details,
			Description: description,
			Type:        typ,
			Price:       price,
			Discount:    discount,
			Tags:        tags,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Product, nil
}

// GetProduct ...
func (c *Client) GetProduct(ctx context.Context, id string, userID string) (*domain.Product, error) {
	r, err := c.service.GetProduct(
		ctx,
		&pb.GetProductRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		ID:          bson.ObjectIdHex(r.Product.Id),
		Name:        r.Product.Name,
		Details:     r.Product.Details,
		Description: r.Product.Description,
		Type:        r.Product.Type,
		Price:       r.Product.Price,
		Discount:    r.Product.Discount,
		Tags:        r.Product.Tags,
	}, nil
}

// GetMultipleProducts is used to get the list of specified catalogues
func (c *Client) GetMultipleProducts(ctx context.Context, skip uint64, take uint64) ([]domain.Product, error) {
	r, err := c.service.GetMultipleProducts(
		ctx,
		&pb.GetMultipleProductsRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	products := []domain.Product{}
	for _, a := range r.Products {

		products = append(products, domain.Product{
			ID:          bson.ObjectIdHex(a.Id),
			Name:        a.Name,
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Discount:    a.Discount,
			Price:       a.Price,
			Tags:        a.Tags,
		})
	}
	return products, nil
}

// DeleteProduct removes a catalogue with passed identifier
func (c *Client) DeleteProduct(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteProduct(
		ctx,
		&pb.DeleteProductRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
