package initConfig

// 一些重要配置
const (
	Ipv4Address             = "自己的ipv4地址"
	ServerPort              = "8080"
	ScourceStaticPath       = "http://" + Ipv4Address + ":" + ServerPort
	ScourceStaticPublicPath = "http://" + Ipv4Address + ":" + ServerPort + "/public/"

	TimeFormat = "2006-01-02 15:04:05"
	AUTH_KEY   = "gindouyin" // jwt鉴权 秘钥
)
