package repositories

import (
	"DonationBE/src/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type ConfigurationsRepo struct {
	dB *sql.DB
}

func NewConfigurationsRepo() ConfigurationsRepo {
	var host = os.Getenv("CONFIG_REPO_HOST")
	var user = os.Getenv("CONFIG_REPO_USER")
	var password = os.Getenv("CONFIG_REPO_PASSWORD")
	var dbName = os.Getenv("CONFIG_REPO_DBNAME")

	connStr := fmt.Sprintf("user='%s' password='%s'  host='%s' dbname='%s'", user, password, host, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	return ConfigurationsRepo{dB: db}
}

func (repo *ConfigurationsRepo) GetConfigurations(id int) (models.Configs, error) {
	rows, err := repo.dB.Query("SELECT * FROM configs WHERE channelid = $1", id)

	if err != nil {
		return models.Configs{}, err
	}

	rows.Next()
	var channelId int
	var stripeKey string
	var streamlabsToken string
	var streamlabsRefreshToken string

	err = rows.Scan(&channelId, &stripeKey, &streamlabsToken, &streamlabsRefreshToken)

	if err != nil {
		println(err.Error())
		return models.Configs{}, err
	}

	config := models.Configs{
		ChannelId:              channelId,
		StripeToken:            stripeKey,
		StreamlabsToken:        streamlabsToken,
		StreamlabsRefreshToken: streamlabsRefreshToken,
	}

	return config, nil

}

func (repo *ConfigurationsRepo) PostConfigurations(configs models.Configs) error {
	_, err := repo.dB.Exec(
		`INSERT INTO configs(channelid, stripekey, streamlabstoken, streamlabsrefreshtoken) VALUES($1, $2, $3, $4)
			ON CONFLICT (channelid) DO UPDATE SET stripekey = $2, streamlabstoken = $3, streamlabsrefreshtoken = $4`,
		configs.ChannelId,
		configs.StripeToken,
		configs.StreamlabsToken,
		configs.StreamlabsRefreshToken,
	)

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
