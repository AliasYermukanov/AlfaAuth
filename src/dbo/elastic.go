package dbo

import (
	"context"
	"encoding/json"
	"github.com/AliasYermukanov/AlfaAuth/src/config"
	"github.com/AliasYermukanov/AlfaAuth/src/domain"
	"github.com/AliasYermukanov/AlfaAuth/src/errors"
	"github.com/olivere/elastic/v7"
)

var gclient *elastic.Client

func ElasticConStart() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(config.AllConfigs.Elastic.ConnectionUrl...), elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetElasticCon() (*elastic.Client, error) {
	if gclient == nil {
		client, err := ElasticConStart()
		if err != nil {
			return nil, err
		} else {
			gclient = client
			return gclient, nil
		}
	} else {
		return gclient, nil
	}
}

func SaveToElastic(index string, uid string, ct interface{}) error {

	client, err := GetElasticCon()

	if err != nil {
		return err
	}

	_, err = client.Index().
		Index(index).
		Id(uid).
		BodyJson(ct).
		Refresh("true").
		Do(context.TODO())

	return err
}

func FindUserByPhoneNumber(phoneNumber string) (*domain.User, error) {
	client, err := GetElasticCon()
	if err != nil {
		errors.ElasticConnectError.DeveloperMessage = err.Error()
		return nil, errors.ElasticConnectError

	}

	mainQuery := elastic.NewBoolQuery()
	mainQuery = mainQuery.Must(elastic.NewMatchQuery("mobile_phone", phoneNumber))

	searchQuery := client.Search()
	search, err := searchQuery.Index("users").Query(mainQuery).From(0).Size(1).Do(context.TODO())
	if err != nil {
		errors.ElasticConnectError.DeveloperMessage = err.Error()
		return nil, errors.ElasticConnectError
	}

	var resp domain.User

	if search.Hits.TotalHits != nil {

		for _, hit := range search.Hits.Hits {

			err := json.Unmarshal(hit.Source, &resp)
			if err != nil {
				return nil, err
			}
		}
	} else {
		errors.NoFound.DeveloperMessage = "no such user"
		return nil, err
	}

	return &resp,nil
}

func GetScopesByUID(uid string)(map[string]string,error)  {
	client, err := GetElasticCon()
	if err != nil {
		errors.ElasticConnectError.DeveloperMessage = err.Error()
		return nil, errors.ElasticConnectError

	}

	mainQuery := elastic.NewBoolQuery()
	mainQuery = mainQuery.Must(elastic.NewMatchQuery("uid", uid))

	searchQuery := client.Search()
	search, err := searchQuery.Index("users").Query(mainQuery).From(0).Size(1).Do(context.TODO())
	if err != nil {
		errors.ElasticConnectError.DeveloperMessage = err.Error()
		return nil, errors.ElasticConnectError
	}

	var resp domain.UsersScope

	if search.Hits.TotalHits != nil {

		for _, hit := range search.Hits.Hits {

			err := json.Unmarshal(hit.Source, &resp)
			if err != nil {
				return nil, err
			}
		}
	} else {
		errors.NoFound.DeveloperMessage = "no such user"
		return nil, err
	}

	return resp.Scope,nil
}