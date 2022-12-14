package service

import (
	"api-star-wars-golang/internal/model"
	"api-star-wars-golang/internal/service/mocks"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_planet_Save(t *testing.T) {
	type fields struct {
		dao   *mocks.PlanetRepository
		swapi SwapiInterface
	}
	type args struct {
		ctx context.Context
		in  *model.PlanetIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
		mock    func(repository *mocks.PlanetRepository)
	}{
		{
			name: "save sucess",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				in: &model.PlanetIn{
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				},
			},
			want:    mock.Anything,
			wantErr: false,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("Save", mock.Anything, mock.Anything).Return(mock.Anything, nil).Once()
			},
		},
		{
			name: "save error",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				in: &model.PlanetIn{
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				},
			},
			want:    "",
			wantErr: true,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("Save", mock.Anything, mock.Anything).Return("", errors.New("error to save"))
			},
		},
		{
			name: "save error2",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				in: &model.PlanetIn{
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				},
			},
			want:    "",
			wantErr: true,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("Save", mock.Anything, mock.Anything).Return("", errors.New("error to save"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				repository: tt.fields.dao,
				swapi:      tt.fields.swapi,
			}
			got, err := s.Save(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

func Test_planet_DeleteById(t *testing.T) {
	type fields struct {
		dao   *mocks.PlanetRepository
		swapi SwapiInterface
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
		mock    func(repository *mocks.PlanetRepository)
	}{
		{
			name: "delete sucess",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			want:    nil,
			wantErr: false,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("DeleteById", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				repository: tt.fields.dao,
				swapi:      tt.fields.swapi,
			}
			if err := s.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

func Test_planet_FindById(t *testing.T) {
	idd := primitive.NewObjectID()
	type fields struct {
		dao   *mocks.PlanetRepository
		swapi SwapiInterface
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.PlanetOut
		wantErr bool
		mock    func(repository *mocks.PlanetRepository)
	}{
		{name: "success",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				id:  mock.Anything,
			},
			want: &model.PlanetOut{
				ID:                      idd,
				Name:                    mock.Anything,
				Climate:                 mock.Anything,
				Terrain:                 mock.Anything,
				NumberOfFilmAppearances: 0,
			},
			wantErr: false,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("FindById", mock.Anything, mock.Anything).Return(&model.Planet{
					ID:      idd,
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				}, nil).Once()
			},
		},
		{
			name: "error",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				id:  mock.Anything,
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("FindById", mock.Anything, mock.Anything).Return(nil, errors.New("error ao buscar"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				repository: tt.fields.dao,
				swapi:      tt.fields.swapi,
			}
			got, err := s.FindById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

func Test_planet_UpdateById(t *testing.T) {
	id := primitive.NewObjectID().Hex()
	//idd, _ := primitive.ObjectIDFromHex(id)

	type fields struct {
		dao   *mocks.PlanetRepository
		swapi SwapiInterface
	}
	type args struct {
		ctx context.Context
		p   model.PlanetIn
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
		mock    func(repository *mocks.PlanetRepository)
	}{
		{name: "success",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				p: model.PlanetIn{
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				},
				id: id,
			},
			want:    nil,
			wantErr: false,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "error",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				p: model.PlanetIn{
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				},
				id: id,
			},
			want:    errors.New("erro ao fazer update"),
			wantErr: true,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("erro ao fazer update"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				repository: tt.fields.dao,
				swapi:      tt.fields.swapi,
			}
			if err := s.Update(tt.args.ctx, tt.args.p, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateById() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

//
func Test_planet_FindByParam(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		dao   *mocks.PlanetRepository
		swapi SwapiInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.PlanetOut
		wantErr bool
		mock    func(repository *mocks.PlanetRepository)
	}{
		{
			name: "findALl success",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: []model.PlanetOut{

				model.PlanetOut{
					ID:                      id,
					Name:                    mock.Anything,
					Climate:                 mock.Anything,
					Terrain:                 mock.Anything,
					NumberOfFilmAppearances: 0,
				},
				model.PlanetOut{
					ID:                      id,
					Name:                    mock.Anything,
					Climate:                 mock.Anything,
					Terrain:                 mock.Anything,
					NumberOfFilmAppearances: 0,
				},
			},
			wantErr: false,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("FindByParam", mock.Anything, mock.Anything).Return([]model.Planet{
					model.Planet{

						ID:      id,
						Name:    mock.Anything,
						Climate: mock.Anything,
						Terrain: mock.Anything,
					},
					model.Planet{
						ID:      id,
						Name:    mock.Anything,
						Climate: mock.Anything,
						Terrain: mock.Anything,
					},
				}, nil).Once()
			},
		},
		{name: "findALl error",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("FindByParam", mock.Anything, mock.Anything).Return(nil, errors.New("erro ao retornar"))
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				repository: tt.fields.dao,
				swapi:      tt.fields.swapi,
			}
			got, err := s.FindByParam(tt.args.ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByParam() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}
func Test_planet_FindByParam_of_arguments(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		dao   *mocks.PlanetRepository
		swapi SwapiInterface
	}
	type args struct {
		ctx  context.Context
		plan *model.PlanetIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.PlanetOut
		wantErr bool
		mock    func(repository *mocks.PlanetRepository)
	}{
		{
			name: "findByParam success",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				plan: &model.PlanetIn{
					Name:    mock.Anything,
					Terrain: mock.Anything,
					Climate: mock.Anything,
				},
			},
			want: []model.PlanetOut{

				model.PlanetOut{
					ID:                      id,
					Name:                    mock.Anything,
					Climate:                 mock.Anything,
					Terrain:                 mock.Anything,
					NumberOfFilmAppearances: 0,
				},
				model.PlanetOut{
					ID:                      id,
					Name:                    mock.Anything,
					Climate:                 mock.Anything,
					Terrain:                 mock.Anything,
					NumberOfFilmAppearances: 0,
				},
			},
			wantErr: false,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("FindByParam", mock.Anything, mock.Anything).Return([]model.Planet{
					model.Planet{

						ID:      id,
						Name:    mock.Anything,
						Climate: mock.Anything,
						Terrain: mock.Anything,
					},
					model.Planet{
						ID:      id,
						Name:    mock.Anything,
						Climate: mock.Anything,
						Terrain: mock.Anything,
					},
				}, nil).Once()
			},
		},
		{
			name: "findByParam success 2",
			fields: fields{
				dao:   new(mocks.PlanetRepository),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				plan: &model.PlanetIn{
					Name:    "Tatooine",
					Terrain: "",
					Climate: "",
				},
			},
			want: []model.PlanetOut{

				model.PlanetOut{
					ID:                      id,
					Name:                    mock.Anything,
					Climate:                 mock.Anything,
					Terrain:                 mock.Anything,
					NumberOfFilmAppearances: 0,
				},
				model.PlanetOut{
					ID:                      id,
					Name:                    mock.Anything,
					Climate:                 mock.Anything,
					Terrain:                 mock.Anything,
					NumberOfFilmAppearances: 0,
				},
			},
			wantErr: false,
			mock: func(repository *mocks.PlanetRepository) {
				repository.On("FindByParam", mock.Anything, mock.Anything).Return([]model.Planet{
					model.Planet{

						ID:      id,
						Name:    mock.Anything,
						Climate: mock.Anything,
						Terrain: mock.Anything,
					},
					model.Planet{
						ID:      id,
						Name:    mock.Anything,
						Climate: mock.Anything,
						Terrain: mock.Anything,
					},
				}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				repository: tt.fields.dao,
				swapi:      tt.fields.swapi,
			}
			got, err := s.FindByParam(tt.args.ctx, tt.args.plan)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByParam() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}
