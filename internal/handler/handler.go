package handler

import (
	"context"
	"database/sql"
	"fmt"
	"go-data-read/internal/db/store"
	"go-data-read/internal/reader"
	"go-data-read/pkg"

	"github.com/google/uuid"
)

type Handler interface {
	GetRecords(context.Context) ([]store.Petroleira, error)
	AddNewRecord(ctx context.Context, arg store.CreateRecordParams) (sql.Result, error)
}

type handlerImpl struct {
	DB  *pkg.DB
	CSV []reader.CamposCsv
}

func NewHandler(db *pkg.DB, csv []reader.CamposCsv) Handler {
	return &handlerImpl{
		DB:  db,
		CSV: csv,
	}
}

func (h *handlerImpl) GetRecords(ctx context.Context) ([]store.Petroleira, error) {
	data, err := h.DB.Queries.GetAllData(ctx)
	if err != nil {
		return nil, fmt.Errorf("Ops! some error has occurred: %w", err)
	}

	return data, nil
}

func (h *handlerImpl) AddNewRecord(ctx context.Context, arg store.CreateRecordParams) (sql.Result, error) {
	data, err := h.DB.Queries.CreateRecord(ctx, store.CreateRecordParams{
		ID:            uuid.New().String(),
		Nome:          h.CSV[0].Nome,
		Organizacao:   h.CSV[0].Organizacao,
		Descricao:     pkg.NewNullString(h.CSV[0].Descricao),
		Tags:          h.CSV[0].Tags,
		Qtdrecursos:   pkg.NewNullInt32(int32(h.CSV[0].QuantidadeRecursos)),
		Qtdreusos:     pkg.NewNullInt32(int32(h.CSV[0].QuantidadeReusos)),
		Qtddownloads:  pkg.NewNullInt32(int32(h.CSV[0].QuantidadeDownloads)),
		Qtdseguidores: pkg.NewNullInt32(int32(h.CSV[0].QuantidadeSeguidores)),
	})
	if err != nil {
		return nil, fmt.Errorf("Ops! Some error has occurred: %w", err)
	}

	return data, nil
}
