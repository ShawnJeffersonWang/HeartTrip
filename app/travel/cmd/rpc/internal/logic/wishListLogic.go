package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"
	"golodge/app/travel/cmd/rpc/internal/svc"
	"golodge/app/travel/cmd/rpc/pb"
	"golodge/app/travel/model"
	"golodge/common/globalkey"

	"github.com/zeromicro/go-zero/core/logx"
)

type WishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WishListLogic {
	return &WishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WishListLogic) WishList(in *pb.WishListReq) (*pb.WishListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.UserHomestayModel.SelectBuilder().Where(squirrel.Eq{
		"user_id":   in.Id,
		"del_state": globalkey.DelStateNo,
	})
	userHomestays, _ := l.svcCtx.UserHomestayModel.FindAll(l.ctx, whereBuilder, "id desc")
	//fmt.Println("userHomestays: ", *userHomestays[0], *userHomestays[1])
	var resp []*pb.Homestay
	if len(userHomestays) > 0 { // mapreduce example
		mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, userHomestay := range userHomestays {
				source <- userHomestay.HomestayId
			}
		}, func(item interface{}, writer mr.Writer[*model.Homestay], cancel func(error)) {
			id := item.(int64)

			homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
			if err != nil && err != model.ErrNotFound {
				logx.WithContext(l.ctx).Errorf("WishListLogic WishList 获取活动数据失败 id : %d ,err : %v", id, err)
				return
			}
			writer.Write(homestay)
		}, func(pipe <-chan *model.Homestay, cancel func(error)) {

			for homestay := range pipe {
				var pbHomestay pb.Homestay
				_ = copier.Copy(&pbHomestay, homestay)
				resp = append(resp, &pbHomestay)
			}
		})
	}

	return &pb.WishListResp{
		List: resp,
	}, nil
}
