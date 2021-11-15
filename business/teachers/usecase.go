package teachers

import (
	"go-schooling/app/middleware"
	"time"
)

type TeacherUsecase struct {
	teacherRepository Repository
	contextTimeout    time.Duration
	jwtAuth           *middleware.ConfigJWT
}

func NewTeacherUsecase(tr Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &TeacherUsecase{
		teacherRepository: tr,
		jwtAuth:           jwtauth,
		contextTimeout:    timeout,
	}
}

