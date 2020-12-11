package config

type Server struct {
	JWT     JWT
	Zap     Zap
	Redis   Redis
	Captcha Captcha
	// gorm
	Mysql Mysql
	// oss
	Local  Local
	Qinniu Qinniu
}
