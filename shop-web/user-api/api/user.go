package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"shop-web/user-api/global/reponse"
	"shop-web/user-api/proto"
	"time"
)

func GetUserList(ctx *gin.Context) {
	ip := "localhost"
	port := 9000
	userConnection, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]",
			"msg", err.Error())
	}
	// 生成grpc的client并调用接口
	userServiceClient := proto.NewUserClient(userConnection)
	rsp, err := userServiceClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    0,
		PSize: 0,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户 列表] 失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	// 返回数据
	res := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := reponse.UserResponse{
			Id:       value.Id,
			Nickname: value.Nickname,
			Birthday: reponse.JsonTime(time.Unix(int64(value.Birthday), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		res = append(res, user)
	}
	ctx.JSON(http.StatusOK, res)

	zap.S().Debug("获取用户列表页")
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换成http的状态码
	if err != nil {
		if se, ok := status.FromError(err); ok {
			switch se.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": se.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可错误",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": se.Code(),
				})
			}
			return
		}
	}
}
