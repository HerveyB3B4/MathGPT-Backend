package env

import "MATHB/config/config"

var JwtSecret = []byte("122443")
var TokenDuration = 2 * 60 * 60
var Port = config.Config.GetInt("server.port")
