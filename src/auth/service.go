package main

import (
	"db/auth_tbl"
	"errors"
	"os"
	"services"
	"services/proto"
	"strings"

	"golang.org/x/net/context"
)

import (
	pb "proto"
	. "types"

	log "github.com/gonet2/libs/nsq-logger"
)

const (
	SERVICE = "[AUTH]"
)

const (
	TABLE_NAME  = "auth"
	SEQS_UID    = "uid"
	SEQS_UNIQUE = "unique"

	LOGIN_TYPE_UUID   = 0
	LOGIN_TYPE_CERT   = 1
	LOGIN_TYPE_WECHAT = 2
	LOGIN_TYPE_ALIPAY = 3
)

type server struct {
	snowflake proto.SnowflakeServiceClient
}

func (s *server) init() {
	c, err := services.GetService(services.SERVICE_SNOWFLAKE)
	if err != nil {
		log.Critical("Snowflake err", err)
		os.Exit(-1)
	}
	s.snowflake, _ = c.(proto.SnowflakeServiceClient)

}

//------------------------------------------------------- user login
func (s *server) Login(ctx context.Context, in *pb.User_LoginInfo) (*pb.User_LoginResp, error) {
	uuid := strings.ToUpper(in.Uuid)
	auth_type := uint8(in.AuthType)
	if uuid == "" {
		return nil, errors.New("require uuid")
	}

	auth := &Auth{}
	switch auth_type {
	case LOGIN_TYPE_UUID:
		auth = auth_tbl.FindByBindID(uuid, auth_type, in.Gsid)
		if auth == nil {
			//insert a new user
			auth = auth_tbl.New(s.next_uid(), uuid, in.Gsid, auth_type, s.next_unique())
		}
	case LOGIN_TYPE_WECHAT:
		fallthrough
	case LOGIN_TYPE_ALIPAY:
		return nil, errors.New("not support yet")
	}

	return &pb.User_LoginResp{Uid: auth.Id, UniqueId: auth.UniqueId, Cert: auth.Cert}, nil
}

func (s *server) next_uid() int32 {
	uid, err := s.snowflake.Next(context.Background(), &proto.Snowflake_Key{Name: SEQS_UID})
	if err != nil {
		log.Critical(err)
		return 0
	}
	return int32(uid.Value)
}
func (s *server) next_unique() uint64 {
	uid, err := s.snowflake.Next(context.Background(), &proto.Snowflake_Key{Name: SEQS_UID})
	if err != nil {
		log.Critical(err)
		return 0
	}
	return uint64(uid.Value)
}