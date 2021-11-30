package company

import (
	"context"
	"github.com/b-entangled/bazel-integration/pkg/design"
)

type Client = design.CompanyServiceServer

type companyServer struct {

}

func NewCompanyServer() Client {
	return &companyServer{}
}

func (cs *companyServer) GetCompany(ctx context.Context, getCompanyRequest *design.GetCompanyRequest) (getCompanyResponse *design.GetCompanyResponse, err error) {
	getCompanyResponse = &design.GetCompanyResponse{
		Company: &design.Company{
			CompanyName:"test", 
			CompanyId: "test", 
			CompanyUrl:"test"}}
	return getCompanyResponse, nil
}
