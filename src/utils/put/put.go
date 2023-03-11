package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type M map[string]string

var wg = sync.WaitGroup{}

var Lock = make(chan struct{}, 500)

var kind = []string{"手机", "电脑", "家具", "资料", "食品", "文具", "冰箱", "饮料", "水果", "鞋子", "衣服", "裤子", "袜子", "帽子", "其他"}

var text = []string{
	"第二次购买，质量没毛病，外观也不丑。",
	"安装完以后，试了一下，效果非常棒，良心店家，售后安装师傅态度好，客服松松服务专业",
	"外观材质外观大气，时尚漂亮，颜色好看静音效果很静音哦。室内室外的静音效果都可以，1.5p制冷热效果快速，安装师傅上门安装，服务态度好。",
	"一直以来都在使用，保湿持久",
	"迫不及待已经用了，一下子买了两个颜色。",
	"多次购买，信赖苏宁，防晒，隔离，都是在苏宁易购，苏宁国际购买，一直回购，大品牌，值得信赖",
	"调解肤色，价格实惠，真心喜欢",
	"一直都是用这个牌子的隔离，提亮肤色，不油腻。有活动一下屯四瓶。真心跟你们推荐这个，苏宁易购买的正品行货，价格超级便宜。",
	"苏宁牛逼，昨天下午购买，今天上午就收到了。还是在疫情防控的时候。",
	"颜值高，物流包装很给力，常用的一款，经济实惠",
	"一直使用这款，感觉效果很不错。",
	"朋友推荐购买的，好用",
	"东西已经收到了，还是可以的。",
	"东西不错，就是发货不及时",
	"已经回购很多次了，还不错哦，有优惠还会购买的。",
	"好好好好好好好好好好好",
	"之前买的紫的。这回试下绿的适合敏感肌易红脸。到时候试试看咋样。看人家说效果挺好的",
	"非正品但是都是一分价一分货",
	"商品拿回来盒子已经拆开了，而且盒子超级脏，跟旧的一样。",
	"字都模模糊糊的，不知道真还是假的，跟以前买的有点不一样",
	"有点不服帖，有点假之前用过正品",
	"饥饿时的救世主哈哈哈",
	"发货很快，买过好多次，全家都喜欢吃。",
	"非常满意！很好吃，吃着很香，停不下来",
	"家里必备，小点心家人都喜欢的，甜度适中，口感细腻",
	"味道还不错，价格实惠，日期新鲜",
	"面包很好吃，早饭，哈哈，推荐大家，支持苏宁",
	"好吃，值得购买！好吃不贵",
	"非常喜欢，质量很好，做工很棒。穿上很潮流，尺寸合适。特别舒适，没有色差。好评",
	"穿着很舒适很好看，质量也特别的好，款式非常新颖。宝贝特别好实物与图片一致下次再来哦。",
	"质量很不错，款式新颖，特别百搭时尚，尺码也很合适。",
	"收到了，很满意，很喜欢做工精很精细，尺码很标准，款式很好看，百搭，穿起来很阳光，很大气，很时尚，点赞。",
	"价格真的很给力，物美价廉，发货速度也特别快，服务也特别好，非常满意。",
	"外观颜值很高，材质做工非常好，摸起来很舒服，大小合适，穿起来很轻便，价格也非常合理优惠。",
	"小米这款笔外观漂亮，书写流畅，握着很有手感，非常不错。发货速度快，服务好极了。",
	"便宜实惠，推荐大家使用购买",
	"好，很好，宝宝喜欢！",
	"宝宝很喜欢，很实惠哟！",
	"快递很给力，刚收到就给宝宝吃上了，还不错。",
	"宝宝10个月了，长牙齿，咬东西，咬咬就不哭了",
	"小朋友很喜欢吃，也很合适",
	"加购买的，很超值，希望苏宁多多搞活动，推出更多实惠的商品。还没给宝宝吃，准备明天试一试。",
	"到货很快，宝宝很爱喝。",
	"宝宝要长牙了，给他买来磨磨牙！",
	"宝宝6个月了，买给宝宝磨牙的，不担心整颗吃",
	"到货了。一直吃，不错，大品牌。可信啊。",
	"宝宝很爱吃，搞活动买的比较划算",
	"买回家后还没有用，等使用的时候再看看效果。",
	"太硬了，宝宝咬不动，不爱吃",
	"宝宝很喜欢吃，没一会就吃了一半，吃的到处都是，还乐的不行",
	"这个挺好的，孩子能吃好久",
	"小孩子很喜欢吃，太棒了",
}

type Poster struct {
	Url    string
	Fields map[string]string
	Files  map[string]string
	Header map[string]string
}

type Getter struct {
	Url    string
	Query  map[string]string
	Header map[string]string
}

type Putter struct {
	Url    string
	Fields map[string]string
	Header map[string]string
}

func main() {
	Manager{}.Default()
}

type Manager struct {
}

func (m Manager) Default() {
	go func() {
		t := 0
		for true {
			fmt.Println("Lasted ", t, " s")
			fmt.Println("goroutine num is ", len(Lock))
			t++
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 10; i++ {
		token, err := m.Login(m.RegisterShopkeeper(i))

		if err != nil {
			fmt.Println(err)
			return
		}
		go func(i int, token string) {
			m.BeShopKeeper(i, token)
			go func() {
				Lock <- struct{}{}
				m.AddGoods(i, 10, token)
				<-Lock
			}()
			go func() {
				Lock <- struct{}{}
				m.ShopKeeperCoupon(token)
				<-Lock
			}()

		}(i, token)
	}
	time.Sleep(time.Second)
	timeout := time.Now().Add(500 * time.Second)
	for {
		if len(Lock) == 0 || time.Now().After(timeout) {
			break
		}
	}
	for j := 0; j < 10; j++ {
		go func(j int) {
			token, err := m.Login(m.RegisterUser(j))
			if err != nil {
				fmt.Println(err)
				return
			}
			go func(token string) {
				for n := 0; n < 10; n++ {
					go func() {
						Lock <- struct{}{}
						m.UserShoppingCar(token)
						<-Lock
					}()
				}
				for l := 0; l < 10; l++ {
					go func() {
						Lock <- struct{}{}
						m.UserStar(token)
						<-Lock
					}()
				}
				for l := 0; l < 10; l++ {
					go func() {
						Lock <- struct{}{}
						m.UserCoupon(token)
						<-Lock
					}()
				}
				for k := 0; k < 200; k++ {
					go func() {
						Lock <- struct{}{}
						m.UserEvaluation(token)
						<-Lock
					}()
				}
			}(token)
		}(j)
	}
	time.Sleep(time.Second)
	timeout2 := time.Now().Add(500 * time.Second)
	for {
		if len(Lock) == 0 || time.Now().After(timeout2) {
			break
		}
	}
	fmt.Println("ok")
}

func (m Manager) AllCouponId() []string {
	var g = Getter{
		Url:    "http://39.101.72.18:9090/coupon/all",
		Query:  nil,
		Header: nil,
	}
	get, err := g.Get()
	if err != nil {
		return nil
	}
	var dst = map[string]interface{}{}
	_ = json.Unmarshal(get, &dst)
	i, ok := dst["data"].([]interface{})
	if !ok {
		return nil
	}
	var idSlice []string
	for _, slice := range i {
		if coupon, ok := slice.(map[string]interface{}); ok {
			idSlice = append(idSlice, fmt.Sprintf("%v", coupon["id"]))
		}
	}
	return idSlice
}

func (m Manager) RandomCouponId() string {
	rand.Seed(time.Now().UnixNano())
	ids := m.AllCouponId()
	return ids[rand.Intn(len(ids))]
}

func (m Manager) UserCoupon(token string) {
	p := Putter{
		Url: "http://39.101.72.18:9090/user/coupon/",
		Header: M{
			"Token": token,
		},
	}
	p.Url += m.RandomCouponId()
	data, err := p.Put()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(data))
	}
	fmt.Println("UserCoupon")
}

func (m Manager) UserEvaluation(token string) {
	p := Poster{
		Url: "http://39.101.72.18:9090/evaluation",
		Fields: M{
			"text":    text[rand.Intn(len(text))],
			"goodsId": m.RandomGoodsId(),
			"score":   strconv.Itoa(rand.Intn(3) + 3),
		},
		Files: M{
			"picture": "./goods/" + strconv.Itoa(1+rand.Intn(20)) + ".jpg",
		},
		Header: M{
			"Token": token,
		},
	}
	data, err := p.Post()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(data))
	}
	fmt.Println("userEvaluation")
}

func (m Manager) UserStar(token string) {
	p := Putter{
		Url: "http://39.101.72.18:9090/user/star/",
		Header: M{
			"Token": token,
		},
	}
	p.Url += m.RandomGoodsId()
	data, err := p.Put()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(data))
	}
	fmt.Println("userStar")
}

func (m Manager) UserShoppingCar(token string) {
	p := Putter{
		Url: "http://39.101.72.18:9090/user/car/",
		Header: M{
			"Token": token,
		},
	}
	p.Url += m.RandomGoodsId()
	data, err := p.Put()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(data))
	}
	fmt.Println("car")
}

func (m Manager) ShopKeeperCoupon(token string) {
	rand.Seed(time.Now().UnixNano())
	p := Poster{
		Url: "http://39.101.72.18:9090/coupon",
		Header: M{
			"Token": token,
		},
	}
	base := time.Now()
	begin := base.Format("2006-01-02 15:04:05")
	end := base.Add(time.Duration(rand.Intn(120)+24) * time.Hour).Format("2006-01-02 15:04:05")
	switch rand.Intn(3) {
	case 0:
		p.Fields = M{
			"kind":     "0",
			"name":     "我是一张现金券",
			"discount": strconv.Itoa(rand.Intn(100)) + "." + strconv.Itoa(rand.Intn(100)),
			"amount":   strconv.Itoa(rand.Intn(500) + 500),
			"beginAt":  begin,
			"endAt":    end,
		}
	case 1:
		p.Fields = M{
			"kind":     "1",
			"name":     "我是一张折扣券",
			"discount": "0." + strconv.Itoa(rand.Intn(10)*5+50),
			"amount":   strconv.Itoa(rand.Intn(500) + 500),
			"beginAt":  begin,
			"endAt":    end,
		}
	case 2:
		p.Fields = M{
			"kind":     "2",
			"name":     "我是一张满减券",
			"discount": strconv.Itoa(rand.Intn(100)+20) + "," + strconv.Itoa(rand.Intn(20)),
			"amount":   strconv.Itoa(rand.Intn(500) + 500),
			"beginAt":  begin,
			"endAt":    end,
		}
	}
	data, err := p.Post()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(data))
	}
	fmt.Println("coupon")
}

func (m Manager) Order(token string) {

}

func (m Manager) RandomGoodsId() string {
	rand.Seed(time.Now().UnixNano())
	ids := m.AllGoodsId()
	return ids[rand.Intn(len(ids))]
}

func (m Manager) Put(shopNum, goodsNum int) {
	for i := 0; i < shopNum; i++ {
		token, err := m.Login(m.RegisterShopkeeper(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		m.BeShopKeeper(i, token)
		m.AddGoods(i, goodsNum, token)
	}
}

func (m Manager) AllGoodsId() []string {
	var g = Getter{
		Url:    "http://39.101.72.18:9090/shop/goods/all",
		Query:  nil,
		Header: nil,
	}
	get, err := g.Get()
	if err != nil {
		return nil
	}
	var dst = map[string]interface{}{}
	_ = json.Unmarshal(get, &dst)
	i, ok := dst["data"].([]interface{})
	if !ok {
		return nil
	}
	var idSlice []string
	for _, slice := range i {
		if goods, ok := slice.(map[string]interface{}); ok {
			idSlice = append(idSlice, fmt.Sprintf("%v", goods["id"]))
		}
	}
	return idSlice
}
func main0() {
	token, err := Manager{}.Login(Manager{}.RegisterShopkeeper(4))
	if err != nil {
		fmt.Println(err)
		return
	}
	Manager{}.BeShopKeeper(4, token)
	Manager{}.AddGoods(4, 500, token)
}
func Get01() {
	var g = Getter{
		Url:    "http://39.101.72.18:9090/shop/goods/all",
		Query:  nil,
		Header: nil,
	}
	g.Get()

}

func (m Manager) BeShopKeeper(num int, token string) {
	var req = Poster{
		Url: "http://39.101.72.18:9090/shop",
		Fields: M{
			"shopName": "[" + strconv.Itoa(num) + "]号商店",
			"notice":   "我是[" + strconv.Itoa(num) + "]号商店的公告",
		},
		Files: nil,
		Header: M{
			"Token": token,
		},
	}
	req.Post()
}

func Post01() {
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
		}
	}()
	wg.Add(500)
	for j := 1; j < 11; j++ {
		time.Sleep(3 * time.Second)
		for i := 1500 + j*50; i < 1500+(j+1)*50; i++ {
			go func(i int) {
				p := Poster{
					Url: "http://39.101.72.18:9090/shop/goods",
					Fields: M{
						"name":  "我是[1]号商店的[" + strconv.Itoa(i) + "]号商品",
						"price": strconv.Itoa(50 + rand.Intn(500)),
						"kind":  kind[rand.Intn(len(kind))],
					},
					Files: M{
						"picture": "./goods/" + strconv.Itoa(1+rand.Intn(20)) + ".jpg",
					},
					Header: M{
						"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiLllYblrrYx5Y+3IiwiZXhwIjoiMTY3NTI3NjYyMiIsIm5iZiI6IjE2NzUyNzMwMjIifQ==.c477e5208319b205f8ee37e6a6d43c2153deb04af74c975595d80d14d24c6866",
					},
				}
				p.Post()
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
	fmt.Println("ok")
	select {}
}

func (m Manager) RegisterUser(num int) (u string, p string) {
	i := strconv.Itoa(num)
	u = "用户[" + i + "]号"
	p = "000000" + i
	var req = Poster{
		Url: "http://39.101.72.18:9090/user/register",
		Fields: M{
			"username":   u,
			"password":   p,
			"rePassword": p,
			"email":      "137" + strconv.Itoa(20000000+num) + "@qq.com",
			"phone":      "137" + strconv.Itoa(20000000+num),
		},
		Files:  nil,
		Header: nil,
	}
	b, err := req.Post()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	return
}

func (m Manager) RegisterShopkeeper(num int) (u string, p string) {
	i := strconv.Itoa(num)
	u = "商家[" + i + "]号"
	p = "000000" + i
	var req = Poster{
		Url: "http://39.101.72.18:9090/user/register",
		Fields: M{
			"username":   u,
			"password":   p,
			"rePassword": p,
			"email":      "137" + strconv.Itoa(10000000+num) + "@qq.com",
			"phone":      "137" + strconv.Itoa(10000000+num),
		},
		Files:  nil,
		Header: nil,
	}
	b, err := req.Post()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	return
}

func (m Manager) Login(u string, p string) (token string, err error) {
	var req = Poster{
		Url: "http://39.101.72.18:9090/user/login",
		Fields: M{
			"username": u,
			"password": p,
		},
		Files:  nil,
		Header: nil,
	}
	res, err := req.Post()
	if err != nil {
		return "", err
	}
	var des = map[string]interface{}{}
	_ = json.Unmarshal(res, &des)
	return des["token"].(string), nil
}

func (m Manager) AddGoods(shopNum int, amount int, token string) {
	wg.Add(amount)
	for i := 1000; i < 1000+amount; i++ {
		go func(i int) {
			Lock <- struct{}{}
			p := Poster{
				Url: "http://39.101.72.18:9090/shop/goods",
				Fields: M{
					"name":  "我是[" + strconv.Itoa(shopNum) + "]号商店的[" + strconv.Itoa(i) + "]号商品",
					"price": strconv.Itoa(100 + rand.Intn(900)),
					"kind":  kind[rand.Intn(len(kind))],
				},
				Files: M{
					"picture": "./goods/" + strconv.Itoa(1+rand.Intn(20)) + ".jpg",
				},
				Header: M{
					"Token": token,
				},
			}
			p.Post()
			<-Lock
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("ok")
}

func (g Getter) Get() ([]byte, error) {
	var urlSlice []string
	for k, v := range g.Query {
		urlSlice = append(urlSlice, k+"="+v)
	}
	var url = g.Url
	if len(urlSlice) != 0 {
		url = strings.Join(urlSlice, "&")
		url = g.Url + "?" + url
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range g.Header {
		req.Header.Add(k, v)
	}
	var client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("创建请求失败")
	}
	defer resp.Body.Close()
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("读取数据失败")
	}
	//fmt.Println(string(d))
	dst := map[string]interface{}{}
	if err := json.Unmarshal(d, &dst); err != nil {
		return nil, err
	}
	res, ok := dst["status"].(float64)
	if !ok {
		return nil, err
	}
	if res != 200 {
		return nil, errors.New(fmt.Sprintf("status=%.0f", res))
	}

	fmt.Println("Get")
	return d, nil
}

func (r Poster) Post() ([]byte, error) {
	var buff bytes.Buffer

	writer := multipart.NewWriter(&buff)

	for k, v := range r.Fields {
		writer.WriteField(k, v)
	}
	for k, v := range r.Files {
		w, err := writer.CreateFormFile(k, v)
		file, err := os.Open(v)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("")
		}
		io.Copy(w, file)

	}
	writer.Close()
	req, err := http.NewRequest(http.MethodPost, r.Url, &buff)
	if err != nil {
		return nil, errors.New("")
	}
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	var client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取失败")
		return nil, errors.New("")
	}
	dst := map[string]interface{}{}
	if err := json.Unmarshal(d, &dst); err != nil {
		return nil, err
	}
	res, ok := dst["status"].(float64)
	if !ok {
		return nil, err
	}
	if res != 200 {
		fmt.Println(dst["msg"])
		return nil, errors.New(fmt.Sprintf("status=%.0f", res))
	}

	fmt.Println("Post")
	return d, nil
}

func (p Putter) Put() ([]byte, error) {
	var urlSlice []string
	for k, v := range p.Fields {
		urlSlice = append(urlSlice, k+"="+v)
	}
	req, err := http.NewRequest(http.MethodPut, p.Url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range p.Header {
		req.Header.Add(k, v)
	}
	var client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("创建请求失败")
	}
	defer resp.Body.Close()
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("读取数据失败")
	}
	//fmt.Println(string(d))
	dst := map[string]interface{}{}
	if err := json.Unmarshal(d, &dst); err != nil {
		return nil, err
	}
	res, ok := dst["status"].(float64)
	if !ok {
		return nil, err
	}
	if res != 200 {
		return nil, errors.New(fmt.Sprintf("status=%.0f", res))
	}

	fmt.Println("Put")

	return d, nil
}
