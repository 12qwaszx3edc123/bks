package server

import (
	"context"
	"errors"

	"bks/common/config"
	"bks/common/model"
	"bks/pkg"
	pbUsers "bks/proto/users"
)

type UsersService struct {
	pbUsers.UnimplementedUsersServer
}

func (s *UsersService) UsersAdd(ctx context.Context, req *pbUsers.UsersAddReq) (resp *pbUsers.UsersAddResp, err error) {
	m := model.Users{
		Mobile:   req.Mobile,
		Password: req.Password,
		Source:   req.Source,
	}
	if err := m.UsersCreate(config.DB); err != nil {
		return nil, errors.New("添加失败")
	}
	return &pbUsers.UsersAddResp{Id: int64(m.ID), Mobile: m.Mobile}, nil
}

func (s *UsersService) UsersLogin(ctx context.Context, req *pbUsers.UsersLoginReq) (resp *pbUsers.UsersLoginResp, err error) {
	var m model.Users
	if err := m.UsersFindByMobile(config.DB, req.Mobile); err != nil {
		return nil, errors.New("用户不存在")
	}
	if m.Password != req.Password {
		return nil, errors.New("密码错误")
	}
	token, expire := pkg.GenerateToken(m.ID)
	return &pbUsers.UsersLoginResp{Token: token, Expire: expire}, nil
}

func (s *UsersService) UsersInfo(ctx context.Context, req *pbUsers.UsersInfoReq) (resp *pbUsers.UsersInfoResp, err error) {
	var m model.Users
	if err := m.UsersFindById(config.DB, uint(req.Id)); err != nil {
		return nil, errors.New("用户不存在")
	}
	return &pbUsers.UsersInfoResp{
		Id:        int64(m.ID),
		Mobile:    m.Mobile,
		Username:  m.Username,
		Avatar:    m.Avatar,
		Status:    int32(m.Status),
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *UsersService) UsersList(ctx context.Context, req *pbUsers.UsersListReq) (resp *pbUsers.UsersListResp, err error) {
	var list []model.Users
	var total int64
	query := config.DB.Model(&model.Users{})
	if req.Keyword != "" {
		query = query.Where("mobile LIKE ? OR username LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.Status != 0 {
		query = query.Where("status = ?", req.Status)
	}
	query.Count(&total)
	query.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize)).Find(&list)
	var respList []*pbUsers.UsersInfoResp
	for _, v := range list {
		respList = append(respList, &pbUsers.UsersInfoResp{
			Id:        int64(v.ID),
			Mobile:    v.Mobile,
			Username:  v.Username,
			Avatar:    v.Avatar,
			Status:    int32(v.Status),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &pbUsers.UsersListResp{List: respList, Total: int32(total)}, nil
}

func (s *UsersService) UsersUpdate(ctx context.Context, req *pbUsers.UsersUpdateReq) (resp *pbUsers.UsersUpdateResp, err error) {
	var m model.Users
	if err := m.UsersFindById(config.DB, uint(req.Id)); err != nil {
		return nil, errors.New("用户不存在")
	}
	if req.Username != "" {
		m.Username = req.Username
	}
	if req.Avatar != "" {
		m.Avatar = req.Avatar
	}
	if req.Status != 0 {
		m.Status = int(req.Status)
	}
	if err := m.UsersUpdate(config.DB); err != nil {
		return nil, errors.New("更新失败")
	}
	return &pbUsers.UsersUpdateResp{Id: int64(m.ID)}, nil
}

func (s *UsersService) UsersDelete(ctx context.Context, req *pbUsers.UsersDeleteReq) (resp *pbUsers.UsersDeleteResp, err error) {
	var m model.Users
	if err := m.UsersDelete(config.DB, uint(req.Id)); err != nil {
		return nil, errors.New("删除失败")
	}
	return &pbUsers.UsersDeleteResp{Success: true}, nil
}
