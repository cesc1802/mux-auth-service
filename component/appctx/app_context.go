package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetAssetDomain() string
	SecretKey() string
}

type appContext struct {
	db          *gorm.DB
	assetDomain string
	secretKey   string
}

func NewAppContext(db *gorm.DB, assetDomain string, secretKey string) *appContext {
	return &appContext{db: db, assetDomain: assetDomain, secretKey: secretKey}
}

func (appCtx *appContext) GetMainDBConnection() *gorm.DB {
	return appCtx.db.Session(&gorm.Session{NewDB: true})
}

func (appCtx *appContext) GetAssetDomain() string {
	return appCtx.assetDomain
}

func (appCtx *appContext) SecretKey() string {
	return appCtx.secretKey
}
