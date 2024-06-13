package service

import (
	"context"
	"database/sql"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServieImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func (service *CategoryServieImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse{

	err:=service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category:=domain.Category{
		Name: request.Name,	
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServieImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse{

	err:=service.Validate.Struct(request)
	helper.PanicIfError(err)
	


	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name
	

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)

}

func (service *CategoryServieImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)


	service.CategoryRepository.Delete(ctx, tx, category)

}

func (service *CategoryServieImpl) FindAll(ctx context.Context, categoryId int) web.CategoryResponse{
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	return helper.ToCategoryResponse(category)


}

func (service *CategoryServieImpl) FindById(ctx context.Context, request web.CategoryCreateRequest) []web.CategoryResponse{
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories:= service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)

}


