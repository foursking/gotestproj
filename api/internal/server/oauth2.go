package server

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/aiscrm/redisgo"
	"github.com/foursking/ztgo/core/log"
	"github.com/foursking/ztgo/core/net/http"
	"github.com/gin-gonic/gin"
)

type TfToken struct {
	Token               string `redis:"token"`
	TokenExpTime        string `redis:"token_exptime"`
	ReflashToken        string `redis:"reflashtoken"`
	ReflashTokenExpTime string `redis:"reflashtoken_exptime"`
}

//var Redis *

func redisObj() *redisgo.Cacher {
	c, err := redisgo.New(
		redisgo.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			//Prefix:   "bossapi_",
		})
	if err != nil {
		panic(err)
	}
	return c
}

/**
 * @description:
 * @param {type}
 * @return:
 */
func genToken(ctx *gin.Context) {

	appid := ctx.PostForm("appid")
	appkey := ctx.PostForm("appkey")

	if !validAccount(appid, appkey) {
		http.JSON(ctx, "wrong", nil)
		log.Errorf("gentoken validaccount error appid:(%v) appkey:(%v)", appid, appkey)
		var err error = errors.New("this is a new error")
		http.JSON(ctx, nil, err)
		return
	}

	var resp = make(map[string]interface{})

	c := redisObj()
	key := "bossapi_" + appid
	if ok, _ := c.Exists(key); ok {
		mhash := TfToken{}
		err := c.HGetAll(key, &mhash)
		if err != nil {
			log.Errorf("hgetall key:(%v) error:(%v)", key, err)
			var err error = errors.New("this is a new error")
			http.JSON(ctx, nil, err)
			return
		}

		resp["token"] = mhash.Token
		resp["token_exptime"] = mhash.TokenExpTime
		resp["reflashtoken"] = mhash.ReflashToken
		resp["reflashtoken_exptime"] = mhash.ReflashTokenExpTime

	} else {
		token := createToken()
		rtoken := createToken()
		resp["token"] = token
		resp["token_exptime"] = aheadNowTime(3600 * 2)
		resp["reflashtoken"] = rtoken
		resp["reflashtoken_exptime"] = aheadNowTime(3600 * 24)
		if err := c.HMSet(key, resp, 86400*3); err != nil {
			log.Errorf("hmset error (%v)", err)
			http.JSON(ctx, nil, err)
			return
		}
	}
	http.JSON(ctx, resp, nil)
}

//验证token
func vaildToken(ctx *gin.Context) {
	appid := ctx.PostForm("appid")
	c := redisObj()
	key := "bossapi_" + appid
	if ok, _ := c.Exists(key); ok {
		var mhash = make(map[string]string)
		c.HGetAll(key, &mhash)
	} else {
		fmt.Println("key not exists")
	}
}

/**
 * @description:
 * @param {type}
 * @return:
 */
func validAccount(appid string, appkey string) bool {
	fmt.Println(appid)
	fmt.Println(appkey)
	return true
}

/**
 * @description: 刷新token
 * @param {type}
 * @return:
 */
func reflashToken(ctx *gin.Context) {
	appid := ctx.PostForm("appid")
	fromReflashToken := ctx.PostForm("reflashToken")
	c := redisObj()
	key := "bossapi_" + appid
	fmt.Println("aaaaa")
	if ok, _ := c.Exists(key); ok {
		//mhash := new(TfToken)
		mhash := TfToken{}
		if err := c.HGetAll(key, &mhash); err != nil {
			log.Errorf("hgetall key:(%v) error:(%v)", key, err)
			http.JSON(ctx, nil, err)
			return
		}

		selfReflashToken := mhash.ReflashToken
		if fromReflashToken == selfReflashToken {
			fmt.Println("key equeue")
			//说明key相等 则可以进行token刷新
			newtoken := createToken()
			t := aheadNowTime(3600 * 2)
			c.HSet(key, "token", newtoken)
			c.HSet(key, "token_exptime", t)
			var resp = make(map[string]interface{})
			resp["token"] = newtoken
			resp["token_exptime"] = t
			http.JSON(ctx, resp, nil)
			return
		}

	} else {
		fmt.Println("key not exists")
		var err error = errors.New("appid hkey not exists")
		log.Errorf("hget appid:(%v) :error (%v)", appid, err)
		http.JSON(ctx, nil, err)
	}
}

//生成40位长度token
func createToken() string {
	i := createRand()
	string := strconv.FormatInt(int64(i), 10)
	sha512 := SHA512Str(string)
	token := sha512[1:40]
	return token
}

func SHA512Str(src string) string {
	h := sha512.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}

func createRand() int {
	rand.Seed(int64(time.Now().UnixNano()))
	return rand.Int()
}

//添加超前时间
func aheadNowTime(i int) int64 {
	i64 := int64(i)
	return time.Now().Unix() + i64
}
