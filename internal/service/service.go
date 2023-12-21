package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Falokut/admin_movies_persons_service/internal/repository"
	movies_persons_service "github.com/Falokut/admin_movies_persons_service/pkg/admin_movies_persons_service/v1/protos"
	"github.com/Falokut/grpc_errors"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MoviesPersonsService struct {
	movies_persons_service.UnimplementedMoviesPersonsServiceV1Server
	logger        *logrus.Logger
	imagesService ImagesService
	repo          repository.PersonsRepository
	errorHandler  errorHandler
}

func NewMoviesPersonsService(logger *logrus.Logger, repo repository.PersonsRepository,
	imagesService ImagesService) *MoviesPersonsService {
	errorHandler := newErrorHandler(logger)
	return &MoviesPersonsService{
		logger:        logger,
		repo:          repo,
		errorHandler:  errorHandler,
		imagesService: imagesService,
	}
}

func (s *MoviesPersonsService) GetPersons(ctx context.Context,
	in *movies_persons_service.GetPersonsRequest) (*movies_persons_service.Persons, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.GetPersons")
	defer span.Finish()

	offset := in.Limit * (in.Page - 1)
	if err := validateLimitAndPage(in.Page, in.Limit); err != nil {
		span.SetTag("grpc.status", status.Code(err))
		ext.LogError(span, err)
		return nil, err
	}

	var persons []repository.Person
	var err error
	if in.PersonsIDs == "" {
		persons, err = s.repo.GetAllPersons(ctx, in.Limit, offset)
	} else {
		in.PersonsIDs = strings.TrimSpace(strings.ReplaceAll(in.PersonsIDs, `"`, ""))
		if err := checkParam(in.PersonsIDs); err != nil {
			return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidParam, "")
		}

		ids := strings.Split(in.PersonsIDs, ",")
		persons, err = s.repo.GetPersons(ctx, ids, in.Limit, offset)
	}

	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	} else if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return s.convertPersons(ctx, persons), nil
}

func (s *MoviesPersonsService) SearchPerson(ctx context.Context,
	in *movies_persons_service.SearchPersonRequest) (*movies_persons_service.Persons, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.SearchPerson")
	defer span.Finish()

	in.PersonName = strings.ReplaceAll(in.PersonName, "'", "")

	if err := validateLimitAndPage(in.Page, in.Limit); err != nil {
		span.SetTag("grpc.status", status.Code(err))
		ext.LogError(span, err)
		return nil, err
	}

	offset := in.Limit * (in.Page - 1)

	persons, err := s.repo.SearchPerson(ctx, in.PersonName, in.Limit, offset)

	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	} else if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return s.convertPersons(ctx, persons), nil
}

func (s *MoviesPersonsService) UpdatePersonFields(ctx context.Context, in *movies_persons_service.UpdatePersonFieldsRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.UpdatePersonFields")
	defer span.Finish()

	exists, err := s.IsPersonExists(ctx, &movies_persons_service.IsPersonExistsRequest{PersonID: in.ID})
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	} else if !exists.PersonExists {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	}

	var photoID = ""
	if len(in.Photo) > 0 {
		photoID, err = s.imagesService.UploadImage(ctx, in.Photo)
		if err != nil {
			span.SetTag("grpc.status", grpc_errors.GetGrpcCode(err))
			ext.LogError(span, err)
			return nil, err
		}
	}

	err = s.repo.UpdatePerson(ctx, in.ID, repository.UpdatePersonParam{
		FullnameRU: in.GetFullnameRU(),
		FullnameEN: in.GetFullnameEN(),
		Birthday:   getTimeFromTimestamp(in.GetBirthday()),
		Sex:        in.GetSex(),
		PhotoID:    photoID,
	}, true)

	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	}
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *MoviesPersonsService) DeletePersons(ctx context.Context,
	in *movies_persons_service.DeletePersonsRequest) (*movies_persons_service.DeletePersonsResponce, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.DeletePerson")
	defer span.Finish()

	in.PersonsIDs = strings.TrimSpace(strings.ReplaceAll(in.PersonsIDs, `"`, ""))
	if err := checkParam(in.PersonsIDs); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidParam, "")
	}
	ids := strings.Split(in.PersonsIDs, ",")

	if len(ids) == 0 {
		return &movies_persons_service.DeletePersonsResponce{DeletedPersonIDs: ids},
			s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, "there is no")
	}

	ids, err := s.repo.DeletePersons(ctx, ids)
	if err != nil {
		return &movies_persons_service.DeletePersonsResponce{DeletedPersonIDs: ids},
			s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &movies_persons_service.DeletePersonsResponce{DeletedPersonIDs: ids}, nil
}

func (s *MoviesPersonsService) IsPersonExists(ctx context.Context,
	in *movies_persons_service.IsPersonExistsRequest) (*movies_persons_service.IsPersonExistsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.IsPersonExists")
	defer span.Finish()

	exists, err := s.repo.IsPersonWithIDExist(ctx, in.PersonID)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	return &movies_persons_service.IsPersonExistsResponse{PersonExists: exists}, nil
}

func (s *MoviesPersonsService) CreatePerson(ctx context.Context,
	in *movies_persons_service.CreatePersonRequest) (*movies_persons_service.CreatePersonResponce, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.CreatePerson")
	defer span.Finish()

	exists, ids, err := s.IsPersonWithFieldsExists(ctx, in)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	} else if exists {
		msg := fmt.Sprintf("finded persons with ids: %s,"+
			"If this list does not contain the id of the person you"+
			"want to add, add more information about the person", strings.Join(ids, ", "))
		return nil, s.errorHandler.createExtendedErrorResponceWithSpan(span, ErrAlreadyExists, "", msg)
	}

	var photoID = ""
	if len(in.Photo) > 0 {
		photoID, err = s.imagesService.UploadImage(ctx, in.Photo)
		if err != nil {
			span.SetTag("grpc.status", status.Code(err))
			ext.LogError(span, err)
			return nil, err
		}
	}

	id, err := s.repo.CreatePerson(ctx, repository.CreatePersonParam{
		FullnameRU: in.GetFullnameRU(),
		FullnameEN: in.GetFullnameEN(),
		Birthday:   getTimeFromTimestamp(in.GetBirthday()),
		Sex:        in.GetSex(),
		PhotoID:    photoID,
	})
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &movies_persons_service.CreatePersonResponce{PersonID: id}, nil
}

func (s *MoviesPersonsService) IsPersonWithFieldsExists(ctx context.Context,
	in *movies_persons_service.CreatePersonRequest) (bool, []string, error) {
	return s.repo.IsPersonAlreadyExists(ctx, repository.IsPersonExistParam{
		FullnameRU: in.FullnameRU,
		FullnameEN: in.GetFullnameEN(),
		Birthday:   getTimeFromTimestamp(in.GetBirthday()),
		Sex:        in.GetSex(),
	})
}

func (s *MoviesPersonsService) UpdatePerson(ctx context.Context, in *movies_persons_service.UpdatePersonRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.UpdatePerson")
	defer span.Finish()

	exists, err := s.IsPersonExists(ctx, &movies_persons_service.IsPersonExistsRequest{PersonID: in.ID})
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	} else if !exists.PersonExists {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	}

	var photoID = ""
	if len(in.Photo) > 0 {
		photoID, err = s.imagesService.UploadImage(ctx, in.Photo)
		if err != nil {
			span.SetTag("grpc.status", status.Code(err))
			ext.LogError(span, err)
			return nil, err
		}
	}
	err = s.repo.UpdatePerson(ctx, in.ID, repository.UpdatePersonParam{
		FullnameRU: in.GetFullnameRU(),
		FullnameEN: in.GetFullnameEN(),
		Birthday:   in.GetBirthday().AsTime(),
		Sex:        in.GetSex(),
		PhotoID:    photoID,
	}, false)

	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	} else if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *MoviesPersonsService) convertPersons(ctx context.Context,
	persons []repository.Person) *movies_persons_service.Persons {
	converted := &movies_persons_service.Persons{}
	converted.Persons = make(map[string]*movies_persons_service.Person, len(persons))
	for _, p := range persons {
		birthday := ""
		if p.Birthday.Valid {
			birthday = p.Birthday.Time.Format("2006-01-02")
		}
		converted.Persons[p.ID] = &movies_persons_service.Person{
			FullnameRU: p.FullnameRU,
			FullnameEN: p.FullnameEN.String,
			Birthday:   birthday,
			Sex:        p.Sex.String,
			PhotoUrl:   s.imagesService.GetPictureURL(ctx, p.PhotoID.String),
		}
	}

	return converted
}

func getTimeFromTimestamp(t *timestamppb.Timestamp) time.Time {
	if t != nil {
		return t.AsTime()
	}

	return time.Time{}
}
