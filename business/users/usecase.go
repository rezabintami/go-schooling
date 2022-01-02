package users

import (
	"context"
	"go-schooling/app/middleware"
	"go-schooling/business"
	"go-schooling/helper/encrypt"
	"go-schooling/helper/logging"
	"strings"
	"time"
)

type UserUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
	logger         logging.Logger
}

func NewUserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &UserUsecase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
		logger:         logger,
	}
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	request := map[string]interface{}{
		"email": email,
		"sso":   sso,
	}

	existedUser, err := uc.userRepository.GetByEmail(ctx, email)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		uc.logger.LogEntry(request, result).Error(err.Error())
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) && !sso {
		return "", business.ErrEmailPasswordNotFound
	}

	token := uc.jwtAuth.GenerateToken(existedUser.ID, existedUser.Roles)
	result := map[string]interface{}{
		"success": "true",
	}
	uc.logger.LogEntry(request, result)
	return token, nil
}

func (uc *UserUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	request := map[string]interface{}{
		"id": id,
	}

	users, err := uc.userRepository.GetByID(ctx, id)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		uc.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"id":        users.ID,
		"name":      users.Name,
		"email":     users.Email,
		"nisn":      users.NISN,
		"class":     users.Classes.Name,
		"graduated": users.Graduated,
		"status":    users.Status,
	}
	uc.logger.LogEntry(request, result)

	return users, nil
}

func (uc *UserUsecase) Update(ctx context.Context, userDomain *Domain, id int) error {
	request := map[string]interface{}{
		"id":        id,
		"name":      userDomain.Name,
		"email":     userDomain.Email,
		"nisn":      userDomain.NISN,
		"class_id":  userDomain.ClassID,
		"graduated": userDomain.Graduated,
		"status":    userDomain.Status,
	}

	err := uc.userRepository.Update(ctx, userDomain, id)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		uc.logger.LogEntry(request, result).Error(err.Error())
		return err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	uc.logger.LogEntry(request, result)

	return nil
}

func (uc *UserUsecase) Register(ctx context.Context, userDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	request := map[string]interface{}{
		"email": userDomain.Email,
		"name":  userDomain.Name,
	}

	existedUser, err := uc.userRepository.GetByEmail(ctx, userDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			result := map[string]interface{}{
				"success": "false",
				"error":   err.Error(),
			}
			uc.logger.LogEntry(request, result).Error(err.Error())
			return err
		}
	}
	if existedUser != (Domain{}) {
		return business.ErrDuplicateData
	}

	if !sso {
		userDomain.Password, _ = encrypt.Hash(userDomain.Password)
	}

	userDomain.Sso = sso
	userDomain.Roles = "USER"
	userDomain.Graduated = false
	err = uc.userRepository.Register(ctx, userDomain)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		uc.logger.LogEntry(request, result).Error(err.Error())
		return err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	uc.logger.LogEntry(request, result).Info("incoming request")

	return nil
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	uc.logger.LogEntry("get all data users", nil)

	result, err := uc.userRepository.GetAll(ctx)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		uc.logger.LogEntry("can't get all data users", result).Error(err.Error())
		return []Domain{}, err
	}

	uc.logger.LogEntry("success to get all data users", nil)

	return result, nil
}

func (uc *UserUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	request := map[string]interface{}{
		"page":     page,
		"per_page": perpage,
	}
	uc.logger.LogEntry(request, nil)

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.userRepository.Fetch(ctx, page, perpage)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		uc.logger.LogEntry("can't get all data users Fetch", result).Error(err.Error())
		return []Domain{}, 0, err
	}
	result := map[string]interface{}{
		"status": "success",
	}
	uc.logger.LogEntry(request, result)

	return res, total, nil
}
