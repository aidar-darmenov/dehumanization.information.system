package db

import (
	"fmt"
	"github.com/aidar-darmenov/dehumanization.information.system/model"
	"github.com/jinzhu/gorm"
)

type Manager struct {
	DB *gorm.DB
	//Cache *redis.Client
}

func NewDB(cfg *model.Configuration) *gorm.DB {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password =%v sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPass)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	db.LogMode(cfg.DBLogMode)
	db.DB().SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.DB().SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.DB().SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	return db
}

func NewDBManager(db *gorm.DB /*, cache *redis.Client*/) *Manager {
	return &Manager{
		DB: db,
		//Cache:cache,
	}
}

/*
func NewCache(config *model.Configuration) *redis.Client{
	cache:=redis.NewClient(&redis.Options{
		Addr:config.RedisHost,
		DB:config.RedisDB,
		Password: config.RedisPassword,})
	if err:=cache.Ping().Err();err!=nil{
		panic(err)
	}
	return cache
}
*/

//AutoMigrate
/*
func (m *Manager) AutoMigrate(){
	m.DB.AutoMigrate(

	)
}
*/
