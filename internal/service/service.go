package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Falokut/admin_movies_persons_service/internal/events"
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
	eventsMQ      events.PersonsEventsMQ
	errorHandler  errorHandler
}

func NewMoviesPersonsService(logger *logrus.Logger,
	repo repository.PersonsRepository,
	imagesService ImagesService,
	eventsMQ events.PersonsEventsMQ) *MoviesPersonsService {
	errorHandler := newErrorHandler(logger)
	return &MoviesPersonsService{
		logger:        logger,
		repo:          repo,
		errorHandler:  errorHandler,
		imagesService: imagesService,
		eventsMQ:      eventsMQ,
	}
}

func (s *MoviesPersonsService) GetPersons(ctx context.Context,
	in *movies_persons_service.GetPersonsRequest) (*movies_persons_service.Persons, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.GetPersons")
	defer span.Finish()

	offset := in.Limit * (in.Page - 1)
	if err := validateLimitAndPage(in.Page, in.Limit); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, err.Error())
	}

	var persons []repository.Person
	var err error
	if in.PersonsIDs == "" {
		persons, err = s.repo.GetAllPersons(ctx, in.Limit, offset)
	} else {
		in.PersonsIDs = strings.TrimSpace(strings.ReplaceAll(in.PersonsIDs, `"`, ""))
		if err := checkParam(in.PersonsIDs); err != nil {
			return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, err.Error())
		}

		ids := strings.Split(in.PersonsIDs, ",")
		persons, err = s.repo.GetPersons(ctx, convertStringsSlice(ids), in.Limit, offset)
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

	offset := in.Limit * (in.Page - 1)
	if err := validateLimitAndPage(in.Page, in.Limit); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, err.Error())
	}

	persons, err := s.repo.SearchPerson(ctx, repository.SearchPersonParam{
		FullnameRU: in.GetFullnameRU(),
		FullnameEN: in.GetFullnameEN(),
		Birthday:   getTimeFromTimestamp(in.Birthday),
		Sex:        in.GetSex(),
	}, in.Limit, offset)

	switch err {
	case repository.ErrInvalidArgument:
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, "")
	case repository.ErrNotFound:
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	case nil:
		span.SetTag("grpc.status", codes.OK)
		return s.convertPersons(ctx, persons), nil
	default:
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}
}

func (s *MoviesPersonsService) UpdatePersonFields(ctx context.Context,
	in *movies_persons_service.UpdatePersonFieldsRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx,
		"MoviesPersonsService.UpdatePersonFields")
	defer span.Finish()

	exists, err := s.IsPersonWithIDExists(ctx,
		&movies_persons_service.IsPersonWithIDExistsRequest{PersonID: in.ID})
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
	} else if err != nil {
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
	if in.PersonsIDs == "" {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, "persons_ids mustn't be empty")
	} else if err := checkParam(in.PersonsIDs); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, err.Error())
	}
	ids := strings.Split(in.PersonsIDs, ",")

	deletedIDs, err := s.repo.DeletePersons(ctx, convertStringsSlice(ids))
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	go func(s *MoviesPersonsService, deletedIDs []int32) {
		for _, id := range deletedIDs {
			err := s.eventsMQ.PersonDeleted(context.Background(), id)
			if err != nil {
				s.logger.Error(err)
			}
		}
	}(s, deletedIDs)
	span.SetTag("grpc.status", codes.OK)
	return &movies_persons_service.DeletePersonsResponce{DeletedPersonIDs: deletedIDs}, nil
}

func (s *MoviesPersonsService) IsPersonWithIDExists(ctx context.Context,
	in *movies_persons_service.IsPersonWithIDExistsRequest) (*movies_persons_service.IsPersonWithIDExistsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.IsPersonWithIDExists")
	defer span.Finish()

	exists, err := s.repo.IsPersonWithIDExist(ctx, in.PersonID)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &movies_persons_service.IsPersonWithIDExistsResponse{PersonExists: exists}, nil
}

func (s *MoviesPersonsService) IsPersonExists(ctx context.Context,
	in *movies_persons_service.IsPersonExistsRequest) (*movies_persons_service.IsPersonExistsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.IsPersonExists")
	defer span.Finish()

	exists, ids, err := s.repo.IsPersonAlreadyExists(ctx, repository.SearchPersonParam{
		FullnameRU: in.GetFullnameRU(),
		FullnameEN: in.GetFullnameEN(),
		Birthday:   getTimeFromTimestamp(in.Birthday),
		Sex:        in.GetSex(),
	})

	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &movies_persons_service.IsPersonExistsResponse{PersonExists: exists, FindedPersonsIDs: ids}, nil
}

func (s *MoviesPersonsService) CreatePerson(ctx context.Context,
	in *movies_persons_service.CreatePersonRequest) (*movies_persons_service.CreatePersonResponce, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.CreatePerson")
	defer span.Finish()

	res, err := s.IsPersonExists(ctx, &movies_persons_service.IsPersonExistsRequest{
		FullnameRU: &in.FullnameRU,
		FullnameEN: in.FullnameEN,
		Birthday:   in.Birthday,
		Sex:        in.Sex,
	})
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	} else if res.PersonExists {
		msg := fmt.Sprintf("finded persons with ids: %s,"+
			"If this list does not contain the id of the person you"+
			"want to add, add more information about the person", formatSlice(res.FindedPersonsIDs))
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

func formatSlice[T any](nums []T) string {
	var str = make([]string, 0, len(nums))
	for _, num := range nums {
		str = append(str, fmt.Sprint(num))
	}
	return strings.Join(str, ",")
}

func convertStringsSlice(str []string) []int32 {
	var nums = make([]int32, 0, len(str))
	for _, s := range str {
		num, err := strconv.Atoi(s)
		if err == nil {
			nums = append(nums, int32(num))
		}
	}
	return nums
}

func (s *MoviesPersonsService) UpdatePerson(ctx context.Context,
	in *movies_persons_service.UpdatePersonRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MoviesPersonsService.UpdatePerson")
	defer span.Finish()

	exists, err := s.IsPersonWithIDExists(ctx,
		&movies_persons_service.IsPersonWithIDExistsRequest{PersonID: in.ID})
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

func (s *MoviesPersonsService) SearchPersonByName(ctx context.Context,
	in *movies_persons_service.SearchPersonByNameRequest) (*movies_persons_service.Persons, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx,
		"MoviesPersonsService.SearchPersonByName")
	defer span.Finish()

	if in.Name == "" {
		return nil, s.errorHandler.createExtendedErrorResponceWithSpan(span, ErrInvalidArgument, "", "name mustn't be empty")
	}

	offset := in.Limit * (in.Page - 1)
	if err := validateLimitAndPage(in.Page, in.Limit); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, err.Error())
	}

	persons, err := s.repo.SearchPersonByName(ctx, in.Name, in.Limit, offset)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	} else if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return s.convertPersons(ctx, persons), nil
}

func (s *MoviesPersonsService) IsPersonsExists(ctx context.Context,
	in *movies_persons_service.IsPersonsExistsRequest) (*movies_persons_service.IsPersonsExistsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx,
		"MoviesPersonsService.IsPersonsExists")
	defer span.Finish()

	in.PersonsIDs = strings.TrimSpace(strings.ReplaceAll(in.PersonsIDs, `"`, ""))
	if in.PersonsIDs == "" {
		return nil, s.errorHandler.createExtendedErrorResponceWithSpan(span, ErrInvalidArgument, "",
			"persons_ids mustn't be empty")
	}

	needCheckIDs := strings.Split(in.PersonsIDs, ",")
	ids, exists, err := s.repo.IsPersonsExists(ctx, convertStringsSlice(needCheckIDs))
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	if exists {
		span.SetTag("grpc.status", codes.OK)
		return &movies_persons_service.IsPersonsExistsResponse{PersonsExists: true}, nil
	}
	findedIDs := make(map[int32]struct{}, len(ids))
	for _, id := range ids {
		if _, ok := findedIDs[id]; !ok {
			findedIDs[id] = struct{}{}
		}
	}

	var notFoundedIDs = make([]int32, 0, len(needCheckIDs))
	for _, id := range needCheckIDs {
		needCheckID, _ := strconv.Atoi(id)
		if _, ok := findedIDs[int32(needCheckID)]; !ok {
			notFoundedIDs = append(notFoundedIDs, int32(needCheckID))
		}
	}

	span.SetTag("grpc.status", codes.OK)
	return &movies_persons_service.IsPersonsExistsResponse{PersonsExists: false, NotExistIDs: notFoundedIDs}, nil
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
