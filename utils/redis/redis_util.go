package redis_util
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"beego_study/caches"
	"errors"
	"github.com/gogather/com/log"
	"beego_study/utils"
)

var redis *caches.MyRedisCache

func InitRedis() {
	cacheConfig := beego.AppConfig.String("cache")

	log.Greenf("cacheConfig:%v \n",cacheConfig)

	var cc cache.Cache
	if "redis" == cacheConfig {
		var err error

		defer utils.Recover("redis init falure")

		cc, err = cache.NewCache("redis", `{"conn":"` + beego.AppConfig.String("redis_host") + `"}`)

		if err != nil {
			log.Redf("%v", err)
		}
		cache, ok := cc.(*caches.MyRedisCache)
		if ok {
			redis = cache
		}else {
			log.Redf("parse cache to MyRedisCache failure !")
		}
	}
}

func Set(key string, val interface{}, expire int64) error {
	var err error
	data, err := utils.Encode(val)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("cc is nil")
	}

	defer utils.Recover("redis set falure")

	err = redis.Set(key, data, expire)
	if err != nil {
		log.Redf("%v", err)
	}

	return err;
}

func Hset(key string,field string, val interface{}, expire int64) error {
	var err error
	data, err := utils.Encode(val)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("cc is nil")
	}

	defer utils.Recover("redis set falure")

	err = redis.Hset(key, field,data, expire)
	if err != nil {
		log.Redf("%v", err)
	}

	return err;
}


func Get(key string, to interface{}) error {
	var err error
	defer utils.Recover("redis get falure")
	data := redis.Get(key)

	if data == nil {
		to = nil
		return errors.New("key point value is nil ")
	}
	err = utils.Decode(data.([]byte), to)
	if err != nil {
		log.Redf("decode failure", err)
	}
	return err
}

func Hget(key string,field string, to interface{}) error {
	var err error
	defer utils.Recover("redis get falure")
	data := redis.Hget(key,field)

	if data == nil {
		to = nil
		return errors.New("key point value is nil ")
	}
	err = utils.Decode(data.([]byte), to)
	if err != nil {
		log.Redf("decode failure", err)
	}
	return err
}
