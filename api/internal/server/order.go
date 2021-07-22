package server

import (
	"crypto/hmac"
	"crypto/sha1"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/service"
	"git.code.oa.com/yifenglu/qdgo-bossapi/api/internal/helper"
	bosstsf "git.code.oa.com/yifenglu/bosstsfcall"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	ffmt "gopkg.in/ffmt.v1"
)

type Item struct {
	name string
	//Type  string `json:"type"`
	value interface{}
}


var (
	sv *service.Service
)

type Head1 struct {
	AppID    int    `json:"appId"`
	Sign     string `json:"sign"`
	Time     int    `json:"time"`
	Stuff    string `json:"stuff"`
	Rawstuff string `json:"rawstuff"`
}

type Head2 struct {
	OrderInfo struct {
		Enterprise []struct {
			Name     string `json:"name"`
			Type     string `json:"type"`
			Value    string `json:"value"`
			RawValue string `json:"rawvalue"`
		} `json:"enterprise"`
		Product []struct {
			Name  string `json:"name"`
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"product"`
		Base []struct {
			Name  string `json:"name"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"base"`
	} `json:"orderInfo"`
}

func parsejson(ctx *gin.Context) {

	json := `{"product":[{"name":"FNumber","type":"string","value":"2852199891"},{"name":"FMealId","type":"int","value":"413"},{"name":"FMonthCount","type":"int","value":"12"},{"name":"FHCOrderId","type":"string","value":"443"}]}`

	m, ok := gjson.Parse(json).Value().(map[string]Item)

	if !ok {
		log.Fatalf("%s", "can not map")
	}
	fmt.Println(m)

	if !gjson.Valid(json) {
		log.Fatalf("%s", "invalid json")
	}

	value := gjson.Get(json, "name.product")

	fmt.Println(value)

}

func submit(ctx *gin.Context) {
	jsonstr := `{"appId":20032,"sign":"6SVt1qGbLz+hzaX\/vmAxXgRJ\/uk=","time":1593314192,"stuff":"173e4c42222172225e22903ee862bccba9734c545f3435c880f166e3598bd12dbe5a33461e73c77dce2bad476ac9d962a3b26f288826b93220a5634f90663ee5059eeb890390c45f08bd8e5ef6e42486c4a7ca92056b6833affe2e0fd67976227c0c891cc6b83e56af8e9c20efc7218bf33abb0bd61ed9848b1beea7ca9a5f82a1bdf89378dc2c2a741c6cca9869a9f56f7cc0995d01dbf141a70fd31747b40195ed506a064c914ce3b21c0acf7da04d63e71eb4097922d8cc8a95224e54fa3a3d985ed44911bb58e43cf0dd19d33dfefe1ce4f6f279dd159d543a3d2aba133ff6c38526900bd1c67c0c891cc6b83e56abe73392c43771c52b0031698d524b2507d34e5dba62f103b8322beb9906cfd67154d702e9b2018ea4b390ce2cb2292d22f83bd9d0c273c62e46b39ce3eadf047f38be7ea6559fa3e4400a7241d4d04d1de87ff21a6b15682b0031698d524b2507d34e5dba62f103b8322beb9906cfd67154d702e9b2018eed92d623b12379256c727a55ee424f364098caf749481c8f1fb7c545ea8a7ee09e76afe0b8a94e0ee7a0e1c1d4950ce251e205851ab4899e73f180591184b16c357df5116bebaf9772f134731725e052507eb6dd7756d9ebcc2bad497207efc6c5d98bd5140c6917474479fca2edc16333d3bc52149e0c21894f35e5767e88076f8857e2886f696838b3a83c1999c44fbe5a33461e73c77dce2bad476ac9d962cbef34dc5a2617a29e76afe0b8a94e0ee7a0e1c1d4950ce251e205851ab4899e2eeadc9a571eb605a56e876604444ec002fc50b6fa5b6286c5c7f86956b90507d88ec283858b5e8fc8c891dadddbc39bbbaa8835cdfc8ad5134504e45f9424d5e4400a7241d4d04d5516b1f8319cdb53dfc170ae661013baf347c34333f89cf8adfbb410310b95373bf6ae1540d82a84c39179565eff19e1d3c9f6e49fb24300961854f01c5370be1a1b9e5f6ce2bb71b95bc06fd9a6e122c4f41135eaad430a36d2aa0bcc22fb4de6c27ddc886720e2be5a33461e73c77d58d8032878d6ca2e8478f2e0a059166c06143b245c0fb1306f7cc0995d01dbf10c056241a8942c4840167f14743f7010029c42be4171b55a2d0427b0915489adf2ee284066879fc3540df0fb49d17f200f43b3f7cd09e38311ae1378880b1f53e495f544d251b0e09e76afe0b8a94e0e0f516437be71f2e52a1a03b1d617ab0e0a59e1d74e1eb13578b9a3d8ac2e2511fafc6d8ccfec58b105490b28d0b256389e76afe0b8a94e0ee7a0e1c1d4950ce251e205851ab4899eb12de7c69a447935d5cb56de0a34de3443e89e999294ff7e69b6e06c171201e7f347c34333f89cf8adfbb410310b95373bf6ae1540d82a8485b58cfbac580b84fe3b4834abaf4e32bd820fc3b520db4a9b7ccf0ac584d55c8a80457d551c1e3d9e76afe0b8a94e0ee7a0e1c1d4950ce251e205851ab4899ee217896ba0233a96b98a98f2fed630d71db3b43beb0524fefafc6d8ccfec58b1ddfa068b951dd38da1bdf89378dc2c2a741c6cca9869a9f56f7cc0995d01dbf17858d5908188773f41c9878afa377ed7e9c7b2a0420900e47ea222326be27c89","rawstuff":"{\"orderInfo\":{\"enterprise\":[{\"name\":\"FLicenceNumber\",\"type\":\"string\",\"value\":\"91441881MA4UJ6UTX5\"},{\"name\":\"FEnterpriseFullName\",\"type\":\"string\",\"value\":\"\\u82f1\\u5fb7\\u5e02\\u9e3f\\u7279\\u7a7a\\u8c03\\u8bbe\\u5907\\u6709\\u9650\\u516c\\u53f8\"},{\"name\":\"FCategory\",\"type\":\"int\",\"value\":\"\\u5de5\\u4e1a\\\/\\u5de5\\u4e1a\\u54c1\"},{\"name\":\"FSubCategory\",\"type\":\"int\",\"value\":\"\\u5176\\u4ed6\"},{\"name\":\"FProvince\",\"type\":\"string\",\"value\":\"\\u5e7f\\u4e1c\\u7701\"},{\"name\":\"FCity\",\"type\":\"string\",\"value\":\"\\u82f1\\u5fb7\\u5e02\"},{\"name\":\"FLicenceNumberUrl\",\"type\":\"file\",\"value\":\"http:\\\/\\\/download.psc.globalsources.com\\\/lic\\\/8855011022.jpg\"},{\"name\":\"FEnterpriseName\",\"type\":\"string\",\"value\":\"YINGDE VENTECH AIR CONDITIONING CO., LTD. \"}],\"product\":[{\"name\":\"FLicenceCount\",\"type\":\"int\",\"value\":1},{\"name\":\"FMealId\",\"type\":\"int\",\"value\":10017},{\"name\":\"FMonthCount\",\"type\":\"int\",\"value\":1}],\"base\":[{\"name\":\"FContactName\",\"type\":\"string\",\"value\":\"Daisy Li\"},{\"name\":\"FContactTel\",\"type\":\"string\",\"value\":13425728677},{\"name\":\"FAgentPosition\",\"type\":\"string\",\"value\":\"Sales Manager\"},{\"name\":\"FContactEmail\",\"type\":\"string\",\"value\":\"sales@chinaventech.com\"}]}}"}`

	var mm Head1
	//m, ok := gjson.Parse(json).Value().(&mm)
	json.Unmarshal([]byte(jsonstr), &mm)
	//fmt.Println(mm)
	rawstuff := mm.Rawstuff

	appinfo := sv.AppInfo(mm.AppID)
	//fmt.Println(rawstuff)
	var head2 Head2
	json.Unmarshal([]byte(rawstuff), &head2)

	//ffmt.Puts(head2)

	for k, v := range head2.OrderInfo.Enterprise {
		if v.Type == "file" {
			head2.OrderInfo.Enterprise[k].RawValue = v.Value
		}
	}

	params := make(map[string]interface{})
	params["params"] = 1
	ret := bosstsf.Call("1010300031", params)
	retstr := helper.Byte2str(ret)
	fmt.Println(retstr)
	//	$FIdData = TsfHelper::callAction(TSF_SERVICE_CMD_COMMON_GETONEIDENTIFICATION, array('HCORDER'));
	ffmt.Puts(ret)

	reqbase := make(map[string]interface{})

	reqbase["FAppId"] = mm.AppID
	reqbase["FUninCode"] = "a"
	reqbase["FMealId"] = 603
	reqbase["FNameAccount"] = appinfo["FNameAccount"]
	reqbase["FKfext"] = 1050
	reqbase["FMasterNameAccount"] = 10
	reqbase["FNumber"] = 1

	tsfreq := make(map[string]interface{})
	tsfreq["request"] = head2
	tsfreq["base"] = reqbase
	ret = bosstsf.Call("1020300006", params)

	//	$this->responseJson(ErrorHelper::SUCCESS, $data);

}

func createsign(params map[string]interface{}, key string) string {
	delete(params, "sign")
	bodyQuery := ksortAndHttpBuildQuery(params)
	str := hmacSha1(bodyQuery, key)
	return b64.StdEncoding.EncodeToString(str)
}

func ksortAndHttpBuildQuery(params map[string]interface{}) string {
	var dataParams string
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		coverstr := fmt.Sprintf("%v", params[k])
		dataParams = dataParams + k + "=" + coverstr + "&"
	}
	return dataParams[0 : len(dataParams)-1]

}

func hmacSha1(data string, secret string) []byte {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
	//return hex.EncodeToString(h.Sum(nil))
}
