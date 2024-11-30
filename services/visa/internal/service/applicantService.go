package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/marelinaa/visa-api/services/visa/internal/domain"
	"github.com/marelinaa/visa-api/services/visa/internal/repository"
)

type ApplicantService struct {
	repo     *repository.Applicant
	validate *validator.Validate
}

func NewApplicantService(repo *repository.Applicant) *ApplicantService {
	gatewayService := &ApplicantService{
		repo:     repo,
		validate: validator.New(),
	}

	return gatewayService
}

func (service *ApplicantService) Apply(ctx context.Context, application domain.Application) error {
	if err := ValidateApplication(application, service.validate); err != nil {

		return err
	}

	return service.repo.AddApplication(ctx, application)

}
