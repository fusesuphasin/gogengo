package service

import (
	"html/template"
	"os"
)

type CredentialServiceTemplate struct{}

func (ct *CredentialServiceTemplate) CreateCredentialServiceTemplate(file *os.File) {
	credentialserviceTemplate.Execute(file, struct {
		}{
	})
}

var credentialserviceTemplate = template.Must(template.New("").Parse(
	`package service

	import (
		"github.com/fusesuphasin/go-fiber/api/3rd/courier/thailand/scg/doministic/repository"
		"github.com/fusesuphasin/go-fiber/app/domain/response"
	)
	
	type CredentialService struct {
		CredentialRepository repository.CredentialRepository
	}
	
	func (rs *CredentialService) GetCredential(credentialAccountID string) response.Credential{
		Credential := rs.CredentialRepository.GetCredential(credentialAccountID)
	
		return Credential
	}
`))