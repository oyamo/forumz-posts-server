package handlers

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"posts-server/internal/domain/persons"
	"time"
)

type PersonHandler struct {
	logger   *zap.SugaredLogger
	personUC *persons.UseCase
}

func (h *PersonHandler) PutPerson(b []byte) {
	var person persons.Person

	err := json.Unmarshal(b, &person)
	if err != nil {
		h.logger.Errorw("error unmarshalling person", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = h.personUC.Upsert(ctx, &person)
	if err != nil {
		h.logger.Errorw("error upserting person", "error", err)
	}
}

func NewPersonHandler(logger *zap.SugaredLogger, personUC *persons.UseCase) *PersonHandler {
	return &PersonHandler{
		logger:   logger,
		personUC: personUC,
	}
}
