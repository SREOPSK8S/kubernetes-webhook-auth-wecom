package repository

import (
	"context"
	"fmt"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/ent"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/ent/migrate"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/url"
)

type RepoStore struct {
	DB *ent.Client
}

type DataClient struct {}
func NewRepoStore(db *ent.Client) *RepoStore {
	return &RepoStore{DB: db}
}

type DataI interface {
	GetConnection() *ent.Client
}

func (DataClient)GetConnection() *ent.Client {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	loc := viper.GetString("datasource.loc")
	rst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		url.QueryEscape(loc),
	)
	client, err := ent.Open(driverName, rst)
	if err != nil {
		logs.Logger.Fatal("failed open connecting to mysql",zap.Any("error",err),zap.String("url",rst))
	}
	return client
}

func (DataClient)CreateSchema(ctx context.Context, c *ent.Client) error {
	err := c.Schema.Create(ctx,migrate.WithGlobalUniqueID(true))
	if err != nil {
		logs.Logger.Fatal("create database schema failure create table error",zap.Any("error",err))
	}
	return  nil
}