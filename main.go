package main

import (
	"context"

	"github.com/megotours/protos/gen/go/auth"
	"github.com/megotours/protos/gen/go/comment"
	"github.com/megotours/protos/gen/go/favorite"
	"github.com/megotours/protos/gen/go/like"
	"github.com/megotours/protos/gen/go/permission"
	"github.com/megotours/protos/gen/go/post"
	"github.com/megotours/protos/gen/go/resource"
	"github.com/megotours/protos/gen/go/role"
	"github.com/megotours/protos/gen/go/storage"
	"github.com/megotours/protos/gen/go/user"
	"google.golang.org/grpc"

	rkboot "github.com/rookie-ninja/rk-boot"
	rkgrpc "github.com/rookie-ninja/rk-grpc/boot"
)

func main() {
	boot := rkboot.NewBoot()

	grpcEntry := rkgrpc.GetGrpcEntry("protos-api")

	grpcEntry.AddRegFuncGrpc(
		registerAuthService,
		registerCommentService,
		registerFavoriteService,
		registerLikeService,
		registerPermissionService,
		registerPostService,
		registerResourceService,
		registerRoleService,
		registerStorageService,
		registerUserService,
	)
	grpcEntry.AddRegFuncGw(
		auth.RegisterAuthServiceHandlerFromEndpoint,
		comment.RegisterCommentServiceHandlerFromEndpoint,
		favorite.RegisterFavoriteServiceHandlerFromEndpoint,
		like.RegisterLikeServiceHandlerFromEndpoint,
		permission.RegisterPermissionServiceHandlerFromEndpoint,
		post.RegisterPostServiceHandlerFromEndpoint,
		resource.RegisterResourceServiceHandlerFromEndpoint,
		role.RegisterRoleServiceHandlerFromEndpoint,
		storage.RegisterStorageServiceHandlerFromEndpoint,
		user.RegisterUserServiceHandlerFromEndpoint,
	)

	boot.Bootstrap(context.Background())

	boot.WaitForShutdownSig(context.Background())
}


func registerAuthService(server *grpc.Server) {
	auth.RegisterAuthServiceServer(server, &AuthServiceServer{})
}

func registerCommentService(server *grpc.Server) {
	comment.RegisterCommentServiceServer(server, &CommentServiceServer{})
}

func registerFavoriteService(server *grpc.Server) {
	favorite.RegisterFavoriteServiceServer(server, &FavoriteServiceServer{})
}

func registerLikeService(server *grpc.Server) {
	like.RegisterLikeServiceServer(server, &LikeServiceServer{})
}

func registerPermissionService(server *grpc.Server) {
	permission.RegisterPermissionServiceServer(server, &PermissionServiceServer{})
}

func registerPostService(server *grpc.Server) {
	post.RegisterPostServiceServer(server, &PostServiceServer{})
}

func registerResourceService(server *grpc.Server) {
	resource.RegisterResourceServiceServer(server, &ResourceServiceServer{})
}

func registerRoleService(server *grpc.Server) {
	role.RegisterRoleServiceServer(server, &RoleServiceServer{})
}

func registerStorageService(server *grpc.Server) {
	storage.RegisterStorageServiceServer(server, &StorageServiceServer{})
}

func registerUserService(server *grpc.Server) {
	user.RegisterUserServiceServer(server, &UserServiceServer{})
}

type AuthServiceServer struct {
	auth.AuthServiceServer
}

type CommentServiceServer struct {
	comment.CommentServiceServer
}

type FavoriteServiceServer struct {
	favorite.FavoriteServiceServer
}

type LikeServiceServer struct {
	like.LikeServiceServer
}

type PermissionServiceServer struct {
	permission.PermissionServiceServer
}

type PostServiceServer struct {
	post.PostServiceServer
}

type ResourceServiceServer struct {
	resource.ResourceServiceServer
}

type RoleServiceServer struct {
	role.RoleServiceServer
}

type StorageServiceServer struct {
	storage.StorageServiceServer
}

type UserServiceServer struct {
	user.UserServiceServer
}
