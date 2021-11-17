package config

type JwtConfig struct {
	Secret string
}

func LoadJwtConfiguration() JwtConfig {
	return JwtConfig{
		Secret: RequireEnvString("JWT_SECRET"),
	}
}
