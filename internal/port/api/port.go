package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"io"
	"navy/internal/common/nerrors"
	"navy/internal/port/domain"
	"navy/internal/port/domain/redis"
	"navy/internal/port/service"
	"net/http"
)

func Run(ctx context.Context) error {
	router := mux.NewRouter()

	// TODO: make configurable
	redisAddress := "redis:6379"
	ports, err := service.NewService(service.WithRepository(redis.NewRedisRepository(redisAddress)))
	if err != nil {
		return errors.Wrapf(err, "failed to create service")
	}

	router.HandleFunc("/ports", handlePutPorts(ports)).Methods("PUT")

	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		select {
		case <-ctx.Done():
			log.Print("Server Stopping")
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatalf("Server Shutdown Failed:%+v", err)
			}

			log.Print("Server Stopped")
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

	return nil
}

func handlePutPorts(ports *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		batchSize := 100 // TODO: make configurable
		batch := make([]domain.Port, 0, batchSize)
		stream, err := NewPortStream(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		for {
			next, err := stream.Next()
			if err != nil {
				if err == io.EOF {
					break
				}

				w.WriteHeader(http.StatusBadRequest)

				return
			}

			batch = append(batch, next)
			if len(batch) == batchSize {
				log.Info("batch upsert")
				err := ports.BatchUpsert(r.Context(), batch)
				if err != nil { // TODO: handle errors in middleware
					if err == nerrors.ErrInvalidInput {
						w.WriteHeader(http.StatusBadRequest)
					} else {
						w.WriteHeader(http.StatusInternalServerError)
					}
				}
				batch = []domain.Port{}
			}
		}

		if len(batch) > 0 {
			log.Info("batch upsert")
			err := ports.BatchUpsert(r.Context(), batch)
			if err != nil {
				if err == nerrors.ErrInvalidInput {
					w.WriteHeader(http.StatusBadRequest)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}

type PortStream struct {
	decoder *json.Decoder
}

func NewPortStream(r io.Reader) (*PortStream, error) {
	decoder := json.NewDecoder(r)
	_, err := decoder.Token() // read the opening bracket
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read opening bracket")
	}

	return &PortStream{
		decoder: decoder,
	}, nil
}
func (j *PortStream) Next() (domain.Port, error) {
	if !j.decoder.More() {
		return domain.Port{}, io.EOF
	}

	p := domain.Port{}

	id, err := j.decoder.Token() // read the field name
	if err != nil {
		return p, errors.Wrapf(err, "failed to read field name")
	}
	err = j.decoder.Decode(&p)
	if err != nil {
		return p, errors.Wrapf(err, "failed to decode port")
	}

	p.ID = id.(string)

	return p, nil
}
