//go:generate protoc ./wallet.proto --go_out=plugins=grpc:./pb

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

	userProfile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
	"github.com/hackerrithm/blackfox/services/backend/wallet/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/wallet/pkg/model"
)

type grpcServer struct {
	service engine.Wallet
	profile *userProfile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Wallet, userURL string, port int) error {
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
	pb.RegisterWalletServiceServer(serv, &grpcServer{s, profileClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostWallet(ctx context.Context, r *pb.PostWalletRequest) (*pb.PostWalletResponse, error) {
	var currency = pb.Currency{}
	var money = pb.Money{}
	var token = pb.Token{}

	err := s.service.Insert(ctx, r.UserID, r.Details, r.Description, r.Type, currency.Abbreviation, currency.Name, currency.Type, token.Type, money.Amount, token.Value, token.Amount)

	if err != nil {
		log.Println("here error in wallet method")
		return nil, err
	}
	return &pb.PostWalletResponse{
		Wallet: "inserted",
	}, nil
}

func (s *grpcServer) PutWallet(ctx context.Context, r *pb.PutWalletRequest) (*pb.PutWalletResponse, error) {
	var currency = pb.Currency{}
	var money = pb.Money{}
	var token = pb.Token{}

	err := s.service.Update(ctx, r.Id, r.UserID, r.Details, r.Description, r.Type, currency.Abbreviation, currency.Name, currency.Type, token.Type, money.Amount, token.Value, token.Amount)
	if err != nil {
		return nil, err
	}
	return &pb.PutWalletResponse{
		Wallet: "updated",
	}, nil
}

func (s *grpcServer) GetWallet(ctx context.Context, r *pb.GetWalletRequest) (*pb.GetWalletResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	var currency = pb.Currency{}
	var money = pb.Money{}
	var token = pb.Token{}

	token.Amount = a.Tokens.Amount
	token.Type = a.Tokens.Type
	token.Value = a.Tokens.Value

	money.Amount = a.Balance.Amount
	money.Currency = &currency

	return &pb.GetWalletResponse{
		Wallet: &pb.Wallet{
			Id:          a.ID.Hex(),
			UserID:      a.UserID.Hex(),
			Balance:     &money,
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Tokens:      &token,
		},
	}, nil
}

func (s *grpcServer) GetMultipleWallets(ctx context.Context, r *pb.GetMultipleWalletsRequest) (*pb.GetMultipleWalletsResponse, error) {
	log.Println("got to grpc ")
	var currency = pb.Currency{}
	var money = pb.Money{}
	var token = pb.Token{}

	res, err := s.service.ListAllWallets(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	wallets := []*pb.Wallet{}
	for _, a := range *res {

		token.Amount = a.Tokens.Amount
		token.Type = a.Tokens.Type
		token.Value = a.Tokens.Value

		money.Amount = a.Balance.Amount
		money.Currency = &currency

		wallets = append(
			wallets,
			&pb.Wallet{
				Id:          a.ID.Hex(),
				UserID:      a.UserID.Hex(),
				Balance:     &money,
				Details:     a.Details,
				Description: a.Description,
				Type:        a.Type,
				Tokens:      &token,
			},
		)
	}

	return &pb.GetMultipleWalletsResponse{
		Wallets: wallets,
	}, nil
}

func (s *grpcServer) DeleteWallet(ctx context.Context, r *pb.DeleteWalletRequest) (*pb.DeleteWalletResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteWalletResponse{
		Id: a,
	}, nil
}
