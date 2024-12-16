package apply

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/marelinaa/visa-api/services/visa/internal/domain"
	"github.com/marelinaa/visa-api/services/visa/internal/repository"
	"github.com/marelinaa/visa-api/services/visa/internal/service"
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

func (s *ApplicantService) Apply(ctx context.Context, application domain.Application) error {
	if err := service.ValidateInput(application, s.validate); err != nil {

		return err
	}

	return s.repo.AddApplication(ctx, application)

}
