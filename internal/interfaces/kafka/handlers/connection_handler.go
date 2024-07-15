package handlers

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"posts-server/internal/domain/connections"
	"time"
)

type ConnectionHandler struct {
	connectionUC *connections.UseCase
	logger       *zap.SugaredLogger
}

func (h *ConnectionHandler) Connect(b []byte) {
	var connection connections.CreateConnectionDTO
	err := json.Unmarshal(b, &connection)
	if err != nil {
		h.logger.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = h.connectionUC.Connect(ctx, &connection)
	if err != nil {
		h.logger.Error(err)
	}
}

func (h *ConnectionHandler) Disconnect(b []byte) {
	var connection connections.CreateConnectionDTO
	err := json.Unmarshal(b, &connection)
	if err != nil {
		h.logger.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = h.connectionUC.Disconnect(ctx, &connection)
	if err != nil {
		h.logger.Error(err)
	}
}

func NewConnectionHandler(connectionUC *connections.UseCase, logger *zap.SugaredLogger) *ConnectionHandler {
	return &ConnectionHandler{
		connectionUC: connectionUC,
		logger:       logger,
	}
}
