// COPYRIGHT(c) 2024 Trenova
//
// This file is part of Trenova.
//
// The Trenova software is licensed under the Business Source License 1.1. You are granted the right
// to copy, modify, and redistribute the software, but only for non-production use or with a total
// of less than three server instances. Starting from the Change Date (November 16, 2026), the
// software will be made available under version 2 or later of the GNU General Public License.
// If you use the software in violation of this license, your rights under the license will be
// terminated automatically. The software is provided "as is," and the Licensor disclaims all
// warranties and conditions. If you use this license's text or the "Business Source License" name
// and trademark, you must comply with the Licensor's covenants, which include specifying the
// Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
// Grant, and not modifying the license in any other way.

package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/emoss08/trenova/internal/server"
	"github.com/emoss08/trenova/pkg/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

// FleetCodeService handles business logic for FleetCode
type FleetCodeService struct {
	db     *bun.DB
	logger *zerolog.Logger
}

// NewFleetCodeService creates a new instance of FleetCodeService
func NewFleetCodeService(s *server.Server) *FleetCodeService {
	return &FleetCodeService{
		db:     s.DB,
		logger: s.Logger,
	}
}

// QueryFilter defines the filter parameters for querying FleetCode
type FleetCodeQueryFilter struct {
	Query          string
	OrganizationID uuid.UUID
	BusinessUnitID uuid.UUID
	Limit          int
	Offset         int
}

// filterQuery applies filters to the query
func (s FleetCodeService) filterQuery(q *bun.SelectQuery, f *FleetCodeQueryFilter) *bun.SelectQuery {
	q = q.Where("fl.organization_id = ?", f.OrganizationID).
		Where("fl.business_unit_id = ?", f.BusinessUnitID)

	if f.Query != "" {
		q = q.Where("fl.code = ? OR fl.description ILIKE ?", f.Query, "%"+strings.ToLower(f.Query)+"%")
	}

	q = q.OrderExpr("CASE WHEN fl.code = ? THEN 0 ELSE 1 END", f.Query).
		Order("fl.created_at DESC")

	return q.Limit(f.Limit).Offset(f.Offset)
}

// GetAll retrieves all FleetCode based on the provided filter
func (s FleetCodeService) GetAll(ctx context.Context, filter *FleetCodeQueryFilter) ([]*models.FleetCode, int, error) {
	var entities []*models.FleetCode

	q := s.db.NewSelect().
		Model(&entities)

	q = s.filterQuery(q, filter)

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch FleetCode")
		return nil, 0, fmt.Errorf("failed to fetch FleetCode: %w", err)
	}

	return entities, count, nil
}

// Get retrieves a single FleetCode by ID
func (s FleetCodeService) Get(ctx context.Context, id, orgID, buID uuid.UUID) (*models.FleetCode, error) {
	entity := new(models.FleetCode)
	err := s.db.NewSelect().
		Model(entity).
		Where("fl.organization_id = ?", orgID).
		Where("fl.business_unit_id = ?", buID).
		Where("fl.id = ?", id).
		Scan(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch FleetCode")
		return nil, fmt.Errorf("failed to fetch FleetCode: %w", err)
	}

	return entity, nil
}

// Create creates a new FleetCode
func (s FleetCodeService) Create(ctx context.Context, entity *models.FleetCode) (*models.FleetCode, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().
			Model(entity).
			Returning("*").
			Exec(ctx)
		return err
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to create FleetCode")
		return nil, fmt.Errorf("failed to create FleetCode: %w", err)
	}

	return entity, nil
}

// UpdateOne updates an existing FleetCode
func (s FleetCodeService) UpdateOne(ctx context.Context, entity *models.FleetCode) (*models.FleetCode, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err := entity.OptimisticUpdate(ctx, tx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to update FleetCode")
		return nil, fmt.Errorf("failed to update FleetCode: %w", err)
	}

	return entity, nil
}