package model

// CacheConfig 缓存配置
type CacheConfig struct {
	Adapter string `json:"adapter"`
	FileDir string `json:"fileDir"`
}

// TokenConfig 登录令牌配置
type TokenConfig struct {
	SecretKey       string `json:"secretKey"`
	Expires         int64  `json:"expires"`
	AutoRefresh     bool   `json:"autoRefresh"`
	RefreshInterval int64  `json:"refreshInterval"`
	MaxRefreshTimes int64  `json:"maxRefreshTimes"`
	MultiLogin      bool   `json:"multiLogin"`
}
