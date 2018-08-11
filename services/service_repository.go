package services

import (
	"github.com/go-squads/reuni-server/helper"
)

const (
	getAllServicesQuery       = "SELECT id,name,created_at,created_by FROM services WHERE organization_id = $1"
	createServiceQuery        = "INSERT INTO services(name, organization_id,authorization_token, created_by) VALUES ($1,$2,$3,$4)"
	deleteServiceQuery        = "DELETE FROM services WHERE name = $1"
	findOneServiceByNameQuery = "SELECT id, name, created_by FROM services WHERE name = $1"
	getServiceTokenQuery      = "SELECT authorization_token FROM services WHERE name = $1"
	translateNameToIdQuery    = "SELECT id FROM organization WHERE name = $1"
)

func getAll(q helper.QueryExecuter, organizationId int) ([]service, error) {
	data, err := q.DoQuery(getAllServicesQuery, organizationId)
	if err != nil {
		return nil, err
	}
	var services []service
	err = helper.ParseMap(data, &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func createService(q helper.QueryExecuter, servicestore service) error {
	_, err := q.DoQuery(createServiceQuery, servicestore.Name, servicestore.OrganizationId, servicestore.AuthorizationToken, servicestore.CreatedBy)
	return err
}

func deleteService(q helper.QueryExecuter, servicestore service) error {
	_, err := q.DoQuery(deleteServiceQuery, servicestore.Name)
	return err
}

func findOneServiceByName(q helper.QueryExecuter, name string) (*service, error) {
	data, err := q.DoQuery(findOneServiceByNameQuery, name)
	if err != nil {
		return nil, err
	}
	var dest service
	if len(data) < 1 {
		return nil, helper.NewHttpError(404, "Not Found")
	}
	err = helper.ParseMap(data[0], &dest)
	if err != nil {
		return nil, err
	}
	return &dest, err
}

func getServiceToken(q helper.QueryExecuter, name string) (*serviceToken, error) {
	var token serviceToken
	data, err := q.DoQuery(getServiceTokenQuery, name)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, helper.NewHttpError(404, "Not Found")
	}
	err = helper.ParseMap(data[0], &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func translateNameToIdRepository(q helper.QueryExecuter, organizationName string) (int, error) {
	data, err := q.DoQueryRow(translateNameToIdQuery, organizationName)
	if err != nil {
		return 0, err
	}
	id := int(data["id"].(int64))
	return id, nil
}
