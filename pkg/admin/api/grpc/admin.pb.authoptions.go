// Code generated by protoc-gen-authmethod. DO NOT EDIT.

package grpc

import (
	"google.golang.org/grpc"

	utils_auth "github.com/caos/citadel/utils/auth"
	utils_grpc "github.com/caos/citadel/utils/grpc"
)

/**
 * AdminService
 */

var AdminService_AuthMethods = utils_auth.AuthMethodMapping{}

func AdminService_Authorization_Interceptor(verifier utils_auth.TokenVerifier, authConf *utils_auth.AuthConfig) grpc.UnaryServerInterceptor {
	return utils_grpc.AuthorizationInterceptor(verifier, authConf, AdminService_AuthMethods)
}
