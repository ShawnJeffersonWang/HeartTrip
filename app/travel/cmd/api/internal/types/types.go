// Code generated by goctl. DO NOT EDIT.
package types

type AddCommentReq struct {
	HomestayId     int64  `json:"homestayId"`
	CommentTime    string `json:"commentTime"`
	Content        string `json:"content"`
	Star           string `json:"star"`
	TidyRating     string `json:"tidyRating"`
	TrafficRating  string `json:"trafficRating"`
	SecurityRating string `json:"securityRating"`
	FoodRating     string `json:"foodRating"`
	CostRating     string `json:"costRating"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	ImageUrls      string `json:"imageUrls"`
}

type AddCommentResp struct {
	Success bool `json:"success"`
}

type AddGuessReq struct {
	HomestayId int64 `json:"homestayId"`
}

type AddGuessResp struct {
	Success bool `json:"success"`
}

type AddHomestayReq struct {
	Title       string `json:"title"`
	TitleTags   string `json:"titleTags"`
	BannerUrls  string `json:"bannerUrls"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Location    string `json:"location"`
	Facilities  string `json:"facilities"`
	Area        string `json:"area"`
	RoomConfig  string `json:"roomConfig"`
	CleanVideo  string `json:"cleanVideo"`
	PriceBefore int64  `json:"priceBefore"` //民宿价格
	PriceAfter  int64  `json:"priceAfter"`
}

type AddHomestayResp struct {
	Success bool `json:"success"`
}

type AdminDeleteHomestayReq struct {
	HomestayId int64 `json:"homestayId"`
}

type AdminDeleteHomestayResp struct {
	Success bool `json:"success"`
}

type BusinessListReq struct {
	LastId             int64 `json:"lastId"`
	PageSize           int64 `json:"pageSize"`
	HomestayBusinessId int64 `json:"homestayBusinessId"`
}

type BusinessListResp struct {
	List []Homestay `json:"list"`
}

type CommentListReq struct {
	HomestayId int64 `path:"homestayId"`
	Page       int64 `form:"page"`
	PageSize   int64 `form:"pageSize"`
}

type CommentListResp struct {
	HomestayId     int64             `json:"homestayId"`
	Star           string            `json:"star"`
	TidyRating     string            `json:"tidyRating"`
	TrafficRating  string            `json:"trafficRating"`
	SecurityRating string            `json:"securityRating"`
	FoodRating     string            `json:"foodRating"`
	CostRating     string            `json:"costRating"`
	CommentCount   int64             `json:"commentCount"`
	List           []HomestayComment `json:"list"`
}

type DeleteHomestayReq struct {
	HomestayId int64 `json:"homestayId"`
}

type DeleteHomestayResp struct {
	Success bool `json:"success"`
}

type GoodBossReq struct {
}

type GoodBossResp struct {
	List []HomestayBusinessBoss `json:"list"`
}

type Guess struct {
	Id          int64  `json:"id"`
	HomestayId  int64  `json:"homestayId"`
	IsCollected bool   `json:"isCollected"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Intro       string `json:"intro"`
	Location    string `json:"location"`
	PriceBefore int64  `json:"priceBefore"` //民宿价格
	PriceAfter  int64  `json:"priceAfter"`
}

type GuessListReq struct {
}

type GuessListResp struct {
	List []Guess `json:"list"`
}

type Homestay struct {
	Id                 int64   `json:"id"`
	Title              string  `json:"title"`
	RatingStars        float64 `json:"ratingStars"`
	CommentCount       int64   `json:"commentCount"`
	TitleTags          string  `json:"titleTags"`
	BannerUrls         string  `json:"bannerUrls"`
	Latitude           string  `json:"latitude"`
	Longitude          string  `json:"longitude"`
	Location           string  `json:"location"`
	Facilities         string  `json:"facilities"`
	Cover              string  `json:"cover"`
	Area               string  `json:"area"`
	RoomConfig         string  `json:"roomConfig"`
	CleanVideo         string  `json:"cleanVideo"`
	HomestayBusinessId int64   `json:"homestayBusinessId"` //店铺id
	HostId             int64   `json:"hostId"`             //房东id
	HostAvatar         string  `json:"hostAvatar"`         // 房东头像
	HostNickname       string  `json:"hostNickname"`       // 房东昵称
	RowState           int64   `json:"rowState"`           //0:下架 1:上架
	PriceBefore        int64   `json:"priceBefore"`        //民宿价格
	PriceAfter         int64   `json:"priceAfter"`
}

type HomestayBusiness struct {
	Id        int64   `json:"id"`
	Title     string  `json:"title"` //店铺名称
	Info      string  `json:"info"`  //店铺介绍
	Tags      string  `json:"tags"`  //标签，多个用“,”分割
	Cover     string  `json:"cover"` //
	Star      float64 `json:"star"`
	IsFav     int64   `json:"isFav"`
	HeaderImg string  `json:"headerImg"` //店招门头图片
}

type HomestayBusinessBoss struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"userId"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Info     string `json:"info"` //房东介绍
	Rank     int64  `json:"rank"` //排名
}

type HomestayBusinessListInfo struct {
	HomestayBusiness
	SellMonth     int64 `json:"sellMonth"`     //月销售
	PersonConsume int64 `json:"personConsume"` //个人消费
}

type HomestayBussinessDetailReq struct {
	Id int64 `json:"id"`
}

type HomestayBussinessDetailResp struct {
	Boss HomestayBusinessBoss `json:"boss"`
}

type HomestayBussinessListReq struct {
	LastId   int64 `json:"lastId"`
	PageSize int64 `json:"pageSize"`
}

type HomestayBussinessListResp struct {
	List []HomestayBusinessListInfo `json:"list"`
}

type HomestayComment struct {
	Id             int64  `json:"id"`
	HomestayId     int64  `json:"homestayId"`
	CommentTime    string `json:"commentTime"`
	Content        string `json:"content"`
	Star           string `json:"star"`
	TidyRating     string `json:"tidyRating"`
	TrafficRating  string `json:"trafficRating"`
	SecurityRating string `json:"securityRating"`
	FoodRating     string `json:"foodRating"`
	CostRating     string `json:"costRating"`
	UserId         int64  `json:"userId"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	ImageUrls      string `json:"imageUrls"`
	LikeCount      int64  `json:"likeCount"`
}

type HomestayDetailReq struct {
	HomestayId int64 `json:"homestayId"`
}

type HomestayDetailResp struct {
	Id                 int64   `json:"id"`
	Title              string  `json:"title"`
	RatingStars        float64 `json:"ratingStars"`
	CommentCount       int64   `json:"commentCount"`
	TitleTags          string  `json:"titleTags"`
	BannerUrls         string  `json:"bannerUrls"`
	Latitude           string  `json:"latitude"`
	Longitude          string  `json:"longitude"`
	Location           string  `json:"location"`
	Facilities         string  `json:"facilities"`
	Area               string  `json:"area"`
	RoomConfig         string  `json:"roomConfig"`
	CleanVideo         string  `json:"cleanVideo"`
	HomestayBusinessId int64   `json:"homestayBusinessId"`
	HostId             int64   `json:"hostId"`
	HostAvatar         string  `json:"hostAvatar"`   // 房东头像
	HostNickname       string  `json:"hostNickname"` // 房东昵称
	PriceBefore        int64   `json:"priceBefore"`  //民宿价格
	PriceAfter         int64   `json:"priceAfter"`
	IsCollected        bool    `json:"isCollected"`
}

type HomestayListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type HomestayListResp struct {
	List []Homestay `json:"list"`
}

type LikeCommentReq struct {
	CommentId int64 `json:"commentId"`
}

type LikeCommentResp struct {
	Success bool `json:"success"`
}

type MyHomestayListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type MyHomestayListResp struct {
	List []Homestay `json:"list"`
}

type SearchByLocationReq struct {
	Location string `json:"location"`
}

type SearchByLocationResp struct {
	List []Homestay `json:"list"`
}

type UpdateHomestayReq struct {
	HomestayId  int64  `json:"homestayId"`
	Title       string `json:"title"`
	TitleTags   string `json:"titleTags"`
	BannerUrls  string `json:"bannerUrls"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Location    string `json:"location"`
	Facilities  string `json:"facilities"`
	Area        string `json:"area"`
	RoomConfig  string `json:"roomConfig"`
	CleanVideo  string `json:"cleanVideo"`
	PriceBefore int64  `json:"priceBefore"` //民宿价格
	PriceAfter  int64  `json:"priceAfter"`
	RowState    int64  `json:"rowState"`
}

type UpdateHomestayResp struct {
	Success bool `json:"success"`
}
