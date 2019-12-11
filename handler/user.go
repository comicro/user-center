package handler

import (
	"context"
	database "user-center/db"
	"user-center/repository"
	"user-center/service"

	"github.com/micro/go-micro/util/log"
	"golang.org/x/crypto/bcrypt"

	user "user-center/proto/user"
)


type User struct{
	repo repository.AuthRepository
	tokenSrv service.TokenService
}

func UserHandle() *User {
	rep := repository.UserRepository{database.Orm()}
	tokenSrv := service.TokenService{}
	return &User{repo: &rep, tokenSrv: tokenSrv}
}

func authError(resp *user.AuthResponse, code int32, message string) error {
	resp.Error = &user.Error{
		Code:    code,
		Message: message,
	}
	return nil
}

func (e *User) Register(ctx context.Context, req *user.User, resp *user.AuthResponse) error {

	if len(req.Name) < 6 {
		return authError(resp, 401, "用户名不能少于6位")
	}
	if len(req.Email) < 6 {
		return authError(resp, 401, "邮箱不能少于6位")
	}
	if len(req.Password) < 6 {
		return authError(resp, 401, "密码不能少于6位")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := e.repo.Create(req); err != nil {
		return err
	}
	token, err := e.tokenSrv.Encode(req)
	if err != nil {
		return err
	}
	log.Info("Logging in with:", req.Email, req.Password)
	resp.User = req
	resp.User.Password = ""
	resp.Token = token
	return nil
}


// Call is a single request handler called via client.Call or the generated client code
func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Log("Received User.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.UserService_StreamStream) error {
	log.Logf("Received User.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&user.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) PingPong(ctx context.Context, stream user.UserService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
