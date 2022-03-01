package config

type Jwt struct {
	SigningKey    string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`          // jwt签名
	ExpiresSecond int64  `mapstructure:"expires-second" json:"expiresSecond" yaml:"expires-second"` // 过期时间（秒）
	BufferSecond  int64  `mapstructure:"buffer-second" json:"bufferSecond" yaml:"buffer-second"`    // 缓冲时间（秒）
	Issuer        string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                        // 签发者
}
