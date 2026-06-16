# bks

> **Author**: 12qwaszx3edc123 (bks)  
> **License**: MIT License  
> **GitHub**: https://github.com/12qwaszx3edc123/bks

Go 寰湇鍔¤剼鎵嬫灦鐢熸垚宸ュ叿 鈥?涓€閿敓鎴愬畬鏁寸殑寰湇鍔￠」鐩粨鏋勶紝鏀寔鏁版嵁搴撳瓧娈佃嚜鍔ㄦ槧灏?

## 瀹夎

```bash
go install github.com/12qwaszx3edc123/bks/cmd/mygen@v1.0.0
```

## 蹇€熷紑濮?

```bash
# 鐢熸垚鍚嶄负 myshop 鐨勯」鐩紝鍖呭惈 user 鍜?order 妯″潡
mygen --name myshop --modules user,order --db-name myshop_db

# 鎸囧畾鏁版嵁搴撹繛鎺ワ紝鑷姩璇诲彇琛ㄥ瓧娈电敓鎴?model/proto/server
mygen --name myproject --modules user,order,doctor \
  --db-host 127.0.0.1 --db-port 3306 \
  --db-user root --db-password root \
  --db-name mydb

# 浣跨敤澶栭儴妯℃澘
mygen --name myproject --modules user --template /path/to/templates
```

## 鍙傛暟璇存槑

| 鍙傛暟 | 榛樿鍊?| 璇存槑 |
|------|--------|------|
| `--name` | `bks` | 椤圭洰鍚嶇О锛堢粷瀵硅矾寰勬垨鐩稿璺緞锛?|
| `--modules` | - | 妯″潡鍒楄〃锛岄€楀彿鍒嗛殧锛屽 `user,order,doctor`锛堝繀濉級 |
| `--bff` | `h5` | BFF 绫诲瀷: `h5`, `web`, `applet`, `app` |
| `--db-host` | `127.0.0.1` | 鏁版嵁搴撲富鏈?|
| `--db-port` | `3306` | 鏁版嵁搴撶鍙?|
| `--db-user` | `root` | 鏁版嵁搴撶敤鎴峰悕 |
| `--db-password` | `root` | 鏁版嵁搴撳瘑鐮?|
| `--db-name` | - | 鏁版嵁搴撳悕绉帮紙濉啓鍚庤嚜鍔ㄨ鍙栬〃瀛楁锛?|
| `--template` | 鍐呭祵妯℃澘 | 澶栭儴妯℃澘鐩綍璺緞 |

## 鐢熸垚鐨勯」鐩粨鏋?

```
<椤圭洰鍚?/
鈹溾攢鈹€ common/
鈹?  鈹溾攢鈹€ config/          # 閰嶇疆绠＄悊锛坈onfig.yaml / config.go / global.go锛?
鈹?  鈹溾攢鈹€ init/            # 鍒濆鍖栵紙MySQL / Redis / ES / gRPC锛?
鈹?  鈹斺攢鈹€ model/           # 鏁版嵁妯″瀷锛堟瘡涓ā鍧椾竴涓枃浠讹級
鈹溾攢鈹€ proto/               # Protobuf 瀹氫箟锛堟瘡涓ā鍧椾竴涓洰褰曪級
鈹溾攢鈹€ <module>-server/     # 鍚勬ā鍧楃殑 gRPC 鏈嶅姟
鈹?  鈹溾攢鈹€ basic/cmd/       # 鏈嶅姟鍏ュ彛 main.go
鈹?  鈹斺攢鈹€ server/          # gRPC 鏈嶅姟瀹炵幇
鈹溾攢鈹€ bff<type>/           # BFF 灞傦紙h5/web/applet/app锛?
鈹?  鈹溾攢鈹€ basic/           # 鍩虹閰嶇疆锛坢ain / middlewares / router锛?
鈹?  鈹斺攢鈹€ handler/         # 璇锋眰澶勭悊锛坅pi / request锛?
鈹斺攢鈹€ pkg/                 # 鍏叡宸ュ叿鍖咃紙jwt / upload / cart / alipay / sendsms / ordersn锛?
```

## 鏍稿績鐗规€?

### 鏁版嵁搴撳瓧娈佃嚜鍔ㄦ槧灏?

鎸囧畾 `--db-name` 鍚庯紝mygen 浼氳嚜鍔ㄨ繛鎺ユ暟鎹簱璇诲彇琛ㄧ粨鏋勶紝鐢熸垚瀹屾暣鐨勫瓧娈垫槧灏勶細

- **Model** 鈥?GORM 妯″瀷锛屽寘鍚纭殑瀛楁绫诲瀷鍜?JSON 鏍囩
- **Proto** 鈥?Protobuf 娑堟伅瀹氫箟锛孧ySQL 绫诲瀷鑷姩杞崲涓?Proto 绫诲瀷
- **Server** 鈥?gRPC 鏈嶅姟鏂规硶
- **Handler** 鈥?BFF 璇锋眰澶勭悊
- **Request** 鈥?璇锋眰/鍝嶅簲缁撴瀯浣?

### 澶氭ā鍧楁敮鎸?

鏀寔涓€娆＄敓鎴愬涓井鏈嶅姟妯″潡锛実RPC 绔彛浠?50051 鑷姩閫掑锛?

```bash
mygen --name myproject --modules user,order,doctor
# user-server   鈫?:50051
# order-server  鈫?:50052
# doctor-server 鈫?:50053
```

### 鍐呭祵妯℃澘

妯℃澘閫氳繃 `//go:embed` 鍐呭祵鍒颁簩杩涘埗鏂囦欢涓紝鏃犻渶棰濆閰嶇疆銆備篃鏀寔 `--template` 鍙傛暟鎸囧畾澶栭儴妯℃澘鐩綍杩涜鑷畾涔夈€?

### MySQL 绫诲瀷鑷姩杞崲

| MySQL 绫诲瀷 | Go 绫诲瀷 | Proto 绫诲瀷 |
|-----------|---------|-----------|
| varchar / char / text | string | string |
| int / tinyint / smallint | int32 | int32 |
| bigint | int64 | int64 |
| float | float32 | float |
| double / decimal | float64 | double |
| date / datetime / timestamp | string | string |
| bit / bool | bool | bool |

## 妯℃澘鏂囦欢

妯℃澘浣嶄簬 `templates/bks/` 鐩綍锛?

```
templates/bks/
鈹溾攢鈹€ bff/
鈹?  鈹溾攢鈹€ cmd/main.go.tmpl
鈹?  鈹溾攢鈹€ handler/api/handler.go.tmpl
鈹?  鈹溾攢鈹€ handler/api/upload.go.tmpl
鈹?  鈹溾攢鈹€ handler/request/request.go.tmpl
鈹?  鈹溾攢鈹€ handler/request/upload.go.tmpl
鈹?  鈹溾攢鈹€ middlewares/middlewares.go.tmpl
鈹?  鈹斺攢鈹€ router/router.go.tmpl
鈹溾攢鈹€ common/
鈹?  鈹溾攢鈹€ config/config.go.tmpl
鈹?  鈹溾攢鈹€ config/config.yaml.tmpl
鈹?  鈹溾攢鈹€ config/global.go.tmpl
鈹?  鈹溾攢鈹€ init/init.go.tmpl
鈹?  鈹斺攢鈹€ model/model.go.tmpl
鈹溾攢鈹€ pkg/
鈹?  鈹溾攢鈹€ alipay.go.tmpl
鈹?  鈹溾攢鈹€ cart.go.tmpl
鈹?  鈹溾攢鈹€ jwt.go.tmpl
鈹?  鈹溾攢鈹€ ordersn.go.tmpl
鈹?  鈹溾攢鈹€ sendsms.go.tmpl
鈹?  鈹斺攢鈹€ upload.go.tmpl
鈹溾攢鈹€ proto/proto.tmpl
鈹斺攢鈹€ srv/
    鈹溾攢鈹€ cmd/main.go.tmpl
    鈹斺攢鈹€ server/server.go.tmpl
```

## 鍚庣画姝ラ

鐢熸垚椤圭洰鍚庯細

1. 淇敼 `common/config/config.yaml` 涓殑绌哄€奸厤缃紙MySQL / Redis / ES 绛夛級
2. 杩愯 `protoc` 鐢熸垚 Go 浠ｇ爜
3. 杩愯 `go mod tidy` 涓嬭浇渚濊禆
4. 鍚姩鏈嶅姟

## 鎶€鏈爤

鐢熸垚鐨勯」鐩娇鐢ㄤ互涓嬫妧鏈爤锛?

- **Web 妗嗘灦**: Gin
- **ORM**: GORM + MySQL
- **缂撳瓨**: Redis (go-redis/v9)
- **鎼滅储寮曟搸**: ElasticSearch (olivere/elastic/v7)
- **RPC**: gRPC + Protobuf
- **閰嶇疆**: Viper
- **璁よ瘉**: JWT (dgrijalva/jwt-go)
- **鏀粯**: Alipay (smartwalle/alipay)
- **鐭俊**: 鑷畾涔夌煭淇℃湇鍔?
- **瀵硅薄瀛樺偍**: 涓冪墰浜?(qiniu/go-sdk)
## 核心功能代码示例

mygen 生成的文章自带以下功能，配置好 config.yaml 即可运行。以下代码可直接复制使用，只需替换包名 `bks` 为你的项目名。

---

### 1. Redis 购物车

`pkg/cart.go` — 用 Redis Hash 存储，Key 格式 `cart:{userId}:{drugId}`

```go
package pkg

import (
	"fmt"
	"bks/common/config"
)

// 加入购物车
func CartAdd(UserId int64, DrugId int64, drugMap map[string]interface{}) error {
	key := fmt.Sprintf("cart:%d:%d", UserId, DrugId)
	return config.RDB.HMSet(config.Ctx, key, drugMap).Err()
}

// 判断购物车里某商品是否存在
func CartExist(UserId int64, DrugId int64) bool {
	key := fmt.Sprintf("cart:%d:%d", UserId, DrugId)
	return config.RDB.Exists(config.Ctx, key).Val() > 0
}

// 更新数量（传入正数增加，负数减少）
func CartUpdate(UserId int64, DrugId int64, quantity int64) error {
	key := fmt.Sprintf("cart:%d:%d", UserId, DrugId)
	return config.RDB.HIncrBy(config.Ctx, key, "quantity", quantity).Err()
}

// 列出用户购物车所有商品 Key
func CartList(UserId int64) []string {
	key := fmt.Sprintf("cart:%d:*", UserId)
	return config.RDB.Keys(config.Ctx, key).Val()
}

// 删除单个商品
func CartDelete(UserId int64, DrugId int64) error {
	key := fmt.Sprintf("cart:%d:%d", UserId, DrugId)
	return config.RDB.Del(config.Ctx, key).Err()
}

// 清空购物车（下单后调用）
func CartClear(UserId int64) error {
	key := fmt.Sprintf("cart:%d:*", UserId)
	keys := config.RDB.Keys(config.Ctx, key).Val()
	for _, k := range keys {
		if err := config.RDB.Del(config.Ctx, k).Err(); err != nil {
			return err
		}
	}
	return nil
}
```

---

### 2. ES 全文搜索 + 多条件筛选

`doctor-server/server/doctor.go` — 医生搜索，支持关键词、科室、价格区间、排序

```go
import (
	"github.com/olivere/elastic/v7"
)

func (s *DoctorService) DoctorList(ctx context.Context, req *pbdoctor.DoctorListRequest) (*pbdoctor.DoctorListResponse, error) {
	boolQuery := elastic.NewBoolQuery()

	// 关键词模糊搜索
	if req.Name != "" {
		boolQuery.Must(elastic.NewMatchQuery("name", req.Name))
	}
	// 科室精确筛选
	if req.DepartmentId != 0 {
		boolQuery.Must(elastic.NewTermQuery("department_id", req.DepartmentId))
	}
	// 价格区间
	if req.MinPrice >= 0 && req.MaxPrice > req.MinPrice {
		boolQuery.Must(elastic.NewRangeQuery("price").Gte(req.MinPrice).Lte(req.MaxPrice))
	}
	// 排序：price/rating/month_sales + Asc/Desc
	sortQuery := elastic.NewFieldSort("_score").Desc()
	if req.Sort != "" {
		slice := strings.Split(req.Sort, "_")
		sortQuery = elastic.NewFieldSort(slice[0])
		if slice[1] == "Desc" { sortQuery.Desc() } else { sortQuery.Asc() }
	}

	do, err := config.Esc.Search().
		Index("doctor").
		Query(boolQuery).
		SortBy(sortQuery).
		From(offset).Size(size).
		Do(context.Background())
	if err != nil {
		return nil, errors.New("ES 搜索失败")
	}

	var list []*pbdoctor.CommodityList
	for _, hit := range do.Hits.Hits {
		item := new(pbdoctor.CommodityList)
		json.Unmarshal(hit.Source, item)
		list = append(list, item)
	}
	return &pbdoctor.CommodityListResponse{
		Count: do.TotalHits(),
		List:  list,
	}, nil
}
```

**ES 写入同步**（新增医生时同步到 ES）：

```go
_, err = config.Esc.Index().
	Index("doctor").
	BodyJson(map[string]interface{}{
		"id":            doctor.ID,
		"doctor_name":   doctor.DoctorName,
		"department_id": doctor.DepartmentId,
		"title":         doctor.Title,
		"server_price":  doctor.ServerPrice,
	}).
	Do(context.Background())
```

---

### 3. 药品下单 — 事务 + 行锁 + 库存扣减

`drug-server/server/drug.go` — 创建订单时对药品表加行锁 `SELECT ... FOR UPDATE`，防止超卖

```go
import (
	"gorm.io/gorm/clause"
)

func (s *DrugService) DrugOrderAdd(ctx context.Context, req *pbDrug.DrugOrderAddReq) (*pbDrug.DrugOrderAddResp, error) {
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil { tx.Rollback() }
	}()
	total := 0.0
	for _, i := range req.List {
		var drug model.Drug
		// ★ 行级锁：SELECT ... FOR UPDATE
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&drug, i.DrugId).Error; err != nil {
			tx.Rollback()
			return nil, errors.New("药品不存在")
		}
		if i.Quantity > drug.DrugNum {
			tx.Rollback()
			return nil, errors.New("库存不足")
		}
		// 扣减库存
		drug.DrugNum -= i.Quantity
		if err := tx.Save(&drug).Error; err != nil {
			tx.Rollback()
			return nil, errors.New("库存扣减失败")
		}
		total += float64(i.Quantity) * drug.DrugPrice
	}
	// 创建订单
	expireTime := time.Now().Add(15 * time.Minute)
	drugOrder := model.DrugOrder{
		DrugSn:     pkg.DrugOrder(),
		UserId:     req.UserId,
		Total:      total,
		PayType:    1,
		Status:     1, // 1=待支付
		ExpireTime: &expireTime,
	}
	if err := tx.Create(&drugOrder).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("订单创建失败")
	}
	tx.Commit()
	// 生成支付宝支付链接
	payURL := pkg.Alipay(drugOrder.DrugSn, total)
	return &pbDrug.DrugOrderAddResp{DrugSn: drugOrder.DrugSn, Total: total, PayUrl: payURL}, nil
}
```

---

### 4. 排班预约 + 号源扣减

`schedule-server/server/schedule.go` — 预约前校验：医生存在、排班存在、时间不冲突、不重复预约

```go
func (s *ScheduleService) AppointmentAdd(ctx context.Context, req *pbSchedule.AppointmentAddReq) (*pbSchedule.AppointmentAddResp, error) {
	// 1. 校验医生存在
	var doctor model.Doctor
	if err := doctor.DoctorFirstId(config.DB, req.DoctorId); err != nil {
		return nil, errors.New("医生不存在")
	}
	// 2. 校验排班存在
	var schedule model.Schedule
	if err := schedule.ScheduleFirstId(config.DB, req.ScheduleId, req.AppointmentTime); err != nil {
		return nil, errors.New("排班不存在")
	}
	// 3. 校验时间合理性
	parse, _ := time.Parse("2006-01-02", req.AppointmentDate)
	nowDay := time.Now().Truncate(24 * time.Hour)
	if parse.Before(nowDay) {
		return nil, errors.New("不能预约过去时间")
	}
	// 4. 防重复预约
	var appointment model.Appointment
	if err := appointment.AppointmentFirst(config.DB,
		req.UserId, req.DoctorId, req.ScheduleId, &parse, req.AppointmentTime); err == nil {
		return nil, errors.New("不能重复预约")
	}
	// 5. 创建预约单（状态=1 待支付，15分钟过期）
	expireTime := time.Now().Add(15 * time.Minute)
	appointment = model.Appointment{
		UserId:      req.UserId,
		DoctorId:    req.DoctorId,
		ScheduleId:  req.ScheduleId,
		Status:      1,
		Total:       doctor.ServerPrice,
		ExpireTime:  &expireTime,
	}
	if err := appointment.AppointmentAdd(config.DB); err != nil {
		return nil, errors.New("预约添加失败")
	}
	return &pbSchedule.AppointmentAddResp{Id: int64(appointment.ID), Total: float32(appointment.Total)}, nil
}
```

---

### 5. 支付宝异步回调通知 + 订单状态更新

`bffH5/handler/api/notify.go` — 接收支付宝 POST 回调，解析表单，按订单前缀分发处理

```go
func Notify(c *gin.Context) {
	c.Request.ParseForm()
	notifyMap := make(map[string]string)
	for k, v := range c.Request.PostForm {
		notifyMap[k] = v[0]
	}
	// 校验订单号
	outTradeNo := notifyMap["out_trade_no"]
	if outTradeNo == "" {
		c.String(200, "fail")
		return
	}
	// 校验支付状态
	if notifyMap["trade_status"] != "TRADE_SUCCESS" {
		c.String(200, "fail")
		return
	}
	// 药品订单回调（订单号前缀 do）
	if strings.Contains(outTradeNo, "do") {
		var drugOrder model.DrugOrder
		if err := drugOrder.FindDrugOrderByDrugSn(config.DB, outTradeNo); err != nil {
			c.String(200, "fail")
			return
		}
		drugOrder.Status = 2 // 已支付
		drugOrder.PayType = 2
		drugOrder.Save(config.DB)
		c.String(200, "success")
		return
	}
	// 预约订单回调（订单号前缀 YY）
	if strings.Contains(outTradeNo, "YY") {
		var appointment model.Appointment
		if err := appointment.FindAppointmentBySn(config.DB, outTradeNo); err != nil {
			c.String(200, "fail")
			return
		}
		appointment.Status = 2
		appointment.Save(config.DB)
		c.String(200, "success")
		return
	}
}
```

---

### 6. 过期订单定时取消 + 库存恢复

`bffH5/handler/api/notify.go` — cron 定时任务，扫描超时未支付订单，恢复库存并置为已取消

```go
func HandlerExpireOrder() {
	var orders []model.DrugOrder
	// 查待支付且已过期的订单
	config.DB.Where("status = 1 AND expire_time < ?", time.Now()).Find(&orders)
	if len(orders) == 0 { return }

	for _, o := range orders {
		tx := config.DB.Begin()
		// 查订单明细
		var items []model.OrderItem
		if err := tx.Where("order_id = ?", o.ID).Find(&items).Error; err != nil {
			tx.Rollback()
			continue
		}
		// ★ 逐行加锁恢复库存
		for _, item := range items {
			var drug model.Drug
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				First(&drug, item.DrugId).Error; err != nil {
				tx.Rollback()
				continue
			}
			drug.DrugNum += item.Quantity
			tx.Save(&drug)
		}
		// 订单状态 → 3(已取消)
		tx.Model(&o).Update("status", 3)
		tx.Commit()
	}
}

// 在 main.go 中启动定时任务：
// cron.New() → cron.AddFunc("@every 30m", HandlerExpireOrder) → cron.Start()
```

---

### 7. 用户注册/登录 + 短信验证码

`users-server/server/users.go` — 图形验证码 + 短信验证 + 自动注册

```go
import "github.com/mojocn/base64Captcha"

// 图形验证码（Redis 限流：1分钟1次）
func (s *UserService) CaptchaAdd(ctx context.Context, req *pbUser.CaptchaAddReq) (*pbUser.CaptchaAddResp, error) {
	count, _ := config.RDB.Incr(config.Ctx, "captcha"+req.Mobile).Result()
	if count > 1 {
		return nil, errors.New("一分钟只能执行一次")
	}
	config.RDB.Expire(config.Ctx, "captcha"+req.Mobile, 1*time.Minute)

	c := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)
	_, base64Img, answer, _ := c.Generate()
	config.RDB.Set(config.Ctx, "captcha"+req.Source+req.Mobile, answer, 5*time.Minute)

	return &pbUser.CaptchaAddResp{Image: base64Img, Success: true}, nil
}

// 短信验证码（Redis 限流：1分钟1条）
func (s *UserService) SendSmsAdd(ctx context.Context, req *pbUser.SendSmsAddReq) (*pbUser.SendSmsAddResp, error) {
	count, _ := config.RDB.Get(config.Ctx, "sendsms"+req.Mobile).Int64()
	if count >= 1 {
		return nil, errors.New("一分钟只能发送一次")
	}
	code := rand.Intn(9000) + 1000
	sms, err := pkg.SendSms(req.Mobile, strconv.Itoa(code)) // 调用互亿短信 API
	if sms.Code != 2 {
		return nil, errors.New("短信发送失败")
	}
	config.RDB.Set(config.Ctx, "sendsms"+req.Source+req.Mobile, code, 5*time.Minute)
	config.RDB.Incr(config.Ctx, "sendsms"+req.Mobile) // 限流计数
	return &pbUser.SendSmsAddResp{Success: true}, nil
}

// 用户注册（校验验证码 → 创建用户 → 创建实名记录）
func (s *UserService) UserAdd(ctx context.Context, req *pbUser.UserAddReq) (*pbUser.UserAddResp, error) {
	// 验证码校验
	captcha, _ := config.RDB.Get(config.Ctx, "captcha"+"login"+req.Mobile).Result()
	if req.Captcha != captcha {
		return nil, errors.New("图形验证码错误")
	}
	sms, _ := config.RDB.Get(config.Ctx, "sendsms"+"login"+req.Mobile).Result()
	if req.SendSms != sms {
		return nil, errors.New("短信验证码错误")
	}
	// 防重复注册
	var user model.User
	if err := user.UserFirstNameOrMobile(config.DB, req.UserName, req.Mobile); err == nil {
		return nil, errors.New("用户已注册")
	}
	user = model.User{
		UserName: req.UserName,
		Password: req.Password,
		Mobile:   req.Mobile,
		Account:  pkg.Account(req.Mobile, req.UserName),
		IsReal:   1,
	}
	if err := user.UserAdd(config.DB); err != nil {
		return nil, errors.New("用户注册失败")
	}
	// 同时创建实名记录
	real := model.UserReal{
		UserId:   user.IsReal,
		UserName: user.UserName,
	}
	real.Add(config.DB)
	return &pbUser.UserAddResp{Id: int64(user.ID)}, nil
}
```

---

### 8. 实名认证（调用腾讯云 API）

`users-server/server/users.go` + `pkg/real.go`

```go
// pkg/real.go — 调用腾讯云市场身份证实名验证 API
func Real(cardNo string, realName string) bool {
	data := config.Cfg.Real
	auth, _, _ := calcAuthorization(data.Id, data.Key)

	bodyParams := urlencode(map[string]string{"cardNo": cardNo, "realName": realName})
	resp, _ := http.Post(
		"https://ap-beijing.cloudmarket-apigw.com/service-18c38npd/idcard/VerifyIdcardv2",
		"application/x-www-form-urlencoded",
		strings.NewReader(bodyParams),
	)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var result struct {
		Result struct { Isok bool `json:"isok"` } `json:"result"`
	}
	json.Unmarshal(bodyBytes, &result)
	return result.Result.Isok
}

// server/ 调用
func (s *UserService) RealAdd(ctx context.Context, req *pbUser.RealAddReq) (*pbUser.RealAddResp, error) {
	var user model.User
	user.UserFirstId(config.DB, req.UserId) // 查用户
	if user.IsReal == 2 { return nil, errors.New("无需重复实名") }

	if !pkg.Real(req.IdCard, req.UserName) { // 调 API 验证
		return nil, errors.New("实名认证失败")
	}
	user.IsReal = 2
	user.Save(config.DB)
	return &pbUser.RealAddResp{Success: true}, nil
}
```

### 9. JWT 鉴权

`pkg/jwt.go` — HS256 签名，支持生成和解析

```go
import "github.com/dgrijalva/jwt-go"

func TokenHandler(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userId,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})
	return token.SignedString([]byte(config.Cfg.Jwt.APP_KEY))
}

func GetToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Jwt.APP_KEY), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("登录异常")
}
```

---

> 💡 以上代码由 mygen 模板自动生成，替换 `bks` 为你的项目名即可编译运行。配置好 `config.yaml` 中的 MySQL/Redis/ES/支付宝/短信连接信息后，直接 `go run`。


## License

MIT License
