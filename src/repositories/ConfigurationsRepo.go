package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type ConfigurationsRepo struct {
	dB *sql.DB
}

func NewConfigurationsRepo() ConfigurationsRepo {
	var host = os.Getenv("REPO_CONFIG_HOST")
	var user = os.Getenv("REPO_CONFIG_USER")
	var password = os.Getenv("REPO_CONFIG_PASSWORD")
	var dbName = os.Getenv("REPO_CONFIG_DBNAME")

	connStr := fmt.Sprintf("user='%s' password='%s'  host='%s' dbname='%s'", user, password, host, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	return ConfigurationsRepo{dB: db}
}

func (repo *ConfigurationsRepo) GetConfigurations(id int) (int, string, string, string, error) {
	rows, err := repo.dB.Query("SELECT * FROM configs WHERE channelid = $1", id)

	if err != nil {
		return 0, "", "", "", err
	}

	rows.Next()
	var channelId int
	var stripeKey string
	var streamlabsToken string
	var streamlabsRefreshToken string

	err = rows.Scan(&channelId, &stripeKey, &streamlabsToken, &streamlabsRefreshToken)

	if err != nil {
		println(err.Error())
		return 0, "", "", "", err
	}

	return channelId, stripeKey, streamlabsToken, streamlabsRefreshToken, nil

}

func (repo *ConfigurationsRepo) PostConfigurations(id int, stripeKey string, streamlabsToken string, streamlabsRefreshToken string) error {
	_, err := repo.dB.Exec("INSERT INTO configs(channelid, stripekey, streamlabstoken, streamlabsrefreshtoken) VALUES($1, $2, $3, $4)", id, stripeKey, streamlabsToken, streamlabsRefreshToken)

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
