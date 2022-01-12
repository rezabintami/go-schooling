package teachers

import (
	"context"
	"go-schooling/app/middleware"
	"go-schooling/business"
	"go-schooling/helper/encrypt"
	"go-schooling/helper/logging"
	"time"
)

type TeacherUsecase struct {
	teacherRepository Repository
	userRepository    interface{}
	contextTimeout    time.Duration
	jwtAuth           *middleware.ConfigJWT
	logger            logging.Logger
}

func NewTeacherUsecase(tr Repository, ur interface{}, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &TeacherUsecase{
		teacherRepository: tr,
		userRepository:    ur,
		jwtAuth:           jwtauth,
		contextTimeout:    timeout,
		logger:            logger,
	}
}

func (tu *TeacherUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	tu.logger.LogEntry("get all data teachers", nil).Info("incoming request")

	result, err := tu.teacherRepository.GetAll(ctx)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		tu.logger.LogEntry("can't get all data teachers", result).Error(err.Error())

		return []Domain{}, err
	}

	tu.logger.LogEntry("success to get all data teachers", nil).Info("incoming request")

	return result, nil

}

func (tu *TeacherUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	request := map[string]interface{}{
		"id": id,
	}

	teachers, err := tu.teacherRepository.GetByID(ctx, id)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		tu.logger.LogEntry(request, result).Error(err.Error())

		return Domain{}, err
	}

	result := map[string]interface{}{
		"id":    teachers.ID,
		"name":  teachers.Name,
		"email": teachers.Email,
		"nip":   teachers.NIP,
		"photo": teachers.Photo,
	}
	tu.logger.LogEntry(request, result).Info("incoming request")

	return teachers, nil
}

func (tu *TeacherUsecase) Update(ctx context.Context, teacherDomain *Domain, id int) error {
	request := map[string]interface{}{
		"id":    id,
		"name":  teacherDomain.Name,
		"email": teacherDomain.Email,
		"nip":   teacherDomain.NIP,
		"photo": teacherDomain.Photo,
	}

	err := tu.teacherRepository.Update(ctx, teacherDomain, id)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		tu.logger.LogEntry(request, result).Error(err.Error())
		return err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	tu.logger.LogEntry(request, result).Info("incoming request")

	return nil
}

func (tu *TeacherUsecase) Store(ctx context.Context, teacherDomain *Domain) error {
	request := map[string]interface{}{
		"name":  teacherDomain.Name,
		"email": teacherDomain.Email,
		"nip":   teacherDomain.NIP,
		"photo": teacherDomain.Photo,
	}

	err := tu.teacherRepository.Store(ctx, teacherDomain)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		tu.logger.LogEntry(request, result).Error(err.Error())
		return err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	tu.logger.LogEntry(request, result).Info("incoming request")

	return nil
}

func (tu *TeacherUsecase) Login(ctx context.Context, email, password string) (string, error) {
	request := map[string]interface{}{
		"email": email,
	}
	existedUser, err := tu.teacherRepository.GetByEmail(ctx, email)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		tu.logger.LogEntry(request, result).Error(err.Error())
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) {
		return "", business.ErrEmailPasswordNotFound
	}

	token := tu.jwtAuth.GenerateToken(existedUser.ID, existedUser.Roles)

	result := map[string]interface{}{
		"success": "true",
	}
	tu.logger.LogEntry(request, result).Info("incoming request")

	return token, nil
}
