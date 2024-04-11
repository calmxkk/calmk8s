package cache

import (
	"calmk8s/internal/library/cache/file"
	"calmk8s/internal/model"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
)

// cache 缓存驱动
var cache *gcache.Cache

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		panic("cache uninitialized.")
	}
	return cache
}

// SetAdapter 设置缓存适配器
func SetAdapter(ctx context.Context) {
	var adapter gcache.Adapter
	var conf model.CacheConfig

	cfg, err := g.Cfg().Get(ctx, "cache.adapter")
	if err != nil {
		g.Log().Fatal(ctx, "get cache config error")
		return
	}

	conf.Adapter = cfg.String()
	if cfg.String() == "file" {
		filedir, err := g.Cfg().Get(ctx, "cache.fileDir")
		if err != nil {
			g.Log().Fatal(ctx, "file path must be configured for file caching.")
			return
		}
		conf.FileDir = filedir.String()
	}

	switch conf.Adapter {
	case "redis":
		adapter = gcache.NewAdapterRedis(g.Redis())
	case "file":
		if conf.FileDir == "" {
			g.Log().Fatal(ctx, "file path must be configured for file caching.")
			return
		}

		if !gfile.Exists(conf.FileDir) {
			if err := gfile.Mkdir(conf.FileDir); err != nil {
				g.Log().Fatalf(ctx, "failed to create the cache directory. procedure, err:%+v", err)
				return
			}
		}
		adapter = file.NewAdapterFile(conf.FileDir)
	default:
		adapter = gcache.NewAdapterMemory()
	}

	// 数据库缓存，默认和通用缓冲驱动一致，如果你不想使用默认的，可以自行调整
	g.DB().GetCache().SetAdapter(adapter)

	// 通用缓存
	cache = gcache.New()
	cache.SetAdapter(adapter)
}
