// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 用户分组
type Group struct {
	Id        int64  `json:"id"`    // 分组id，由微信分配
	Name      string `json:"name"`  // 分组名字，UTF8编码
	UserCount int    `json:"count"` // 分组内用户数量
}

var ErrNotSubscribe = errors.New("没有订阅公众号")

type UserInfo struct {
	OpenId   string `json:"openid"`   // 用户的标识，对当前公众号唯一
	Nickname string `json:"nickname"` // 用户的昵称
	Sex      int    `json:"sex"`      // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language string `json:"language"` // 用户的语言，zh_CN，zh_TW，en
	City     string `json:"city"`     // 用户所在城市
	Province string `json:"province"` // 用户所在省份
	Country  string `json:"country"`  // 用户所在国家

	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	// 用户没有头像时该项为空
	HeadImageURL string `json:"headimgurl,omitempty"`

	// 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	SubscribeTime int64 `json:"subscribe_time"`

	// 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	UnionId string `json:"unionid,omitempty"`

	// 备注名; 这个字段文档中没有，但是实际返回的数据里有这个字段
	Remark string `json:"remark,omitempty"`
}

// 获取用户图像的大小
//  @headImageURL: 用户头像URL，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像）
//  NOTE: 请确保 headImageURL 不为空
func HeadImageSize(headImageURL string) (size int, err error) {
	index := strings.LastIndex(headImageURL, "/")
	if index == -1 {
		err = fmt.Errorf("invalid headImageURL: %s", headImageURL)
		return
	}
	if index+1 == len(headImageURL) { // "/" 在最后面
		err = fmt.Errorf("invalid headImageURL: %s", headImageURL)
		return
	}

	sizeStr := headImageURL[index+1:]

	size64, err := strconv.ParseUint(sizeStr, 10, 8)
	if err != nil {
		err = fmt.Errorf("invalid headImageURL: %s", headImageURL)
		return
	}

	if size64 == 0 {
		size64 = 640
	}

	size = int(size64)
	return
}

// 获取关注者返回的数据结构
type UserGetData struct {
	TotalCount int `json:"total"` // 关注该公众账号的总用户数
	GetCount   int `json:"count"` // 拉取的OPENID个数，最大值为10000

	Data struct {
		OpenId []string `json:"openid"`
	} `json:"data"` // 列表数据，OPENID的列表

	// 拉取列表的后一个用户的OPENID, 如果 next_openid == "" 则表示没有了用户数据
	NextOpenId string `json:"next_openid"`
}

// 关注者列表的遍历器
//
//  iter, err := Client.UserIterator("beginOpenId")
//  if err != nil {
//      // TODO: 增加你的代码
//  }
//
//  for iter.HasNext() {
//      openids, err := iter.NextPage()
//      if err != nil {
//          // TODO: 增加你的代码
//      }
//      // TODO: 增加你的代码
//  }
type UserIterator interface {
	Total() int // 用户总的个数
	HasNext() bool
	NextPage() (openids []string, err error)
}
