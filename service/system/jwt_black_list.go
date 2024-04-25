package system

import (
	"gateway/global"
	"gateway/model/system"
	"go.uber.org/zap"
)

type JwtService struct{}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList Model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JwtInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GatewayDBs["system"].Create(&jwtList).Error
	if err != nil {
		return
	}
	global.GatewayCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.GatewayDBs["system"].Get(jwt)
	return ok
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

//func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
//	redisJWT, err = global.PalRedis.Get(context.Background(), userName).Result()
//	return redisJWT, err
//}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

//func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
//	// 此处过期时间等于jwt过期时间
//	dr, err := utils.ParseDuration(global.PalConfig.JWT.ExpiresTime)
//	if err != nil {
//		return err
//	}
//	timer := dr
//	err = global.PalRedis.Set(context.Background(), userName, jwt, timer).Err()
//	return err
//}

func LoadAll() {
	var data []string
	err := global.GatewayDBs["system"].Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GatewayLog.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.GatewayCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 PalCache 中
}
