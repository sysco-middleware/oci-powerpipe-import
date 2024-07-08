package db

import (
	"context"
	"database/sql"
	"fmt"
)

func (s *Service) GetRunID(ctx context.Context) (int64, error) {
	var runId RunID

	tx, err := s.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer HandleTxRollback(tx)

	if err := tx.GetContext(ctx, &runId, sqlGetRunID); err != nil {
		return 0, fmt.Errorf("failed to get run id: %w", err)
	}

	if _, err := tx.ExecContext(ctx, sqlTruncateRunID); err != nil {
		return 0, fmt.Errorf("failed to truncate run id: %w", err)
	}

	if _, err := tx.ExecContext(ctx, sqlInsertRunID, runId.ID+1); err != nil {
		return 0, fmt.Errorf("failed to insert run id: %w", err)
	}

	return runId.ID + 1, tx.Commit()
}

func (s *Service) BulkInsert(ctx context.Context, data []Report) error {
	tx, err := s.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer HandleTxRollback(tx)

	_, err = tx.NamedExecContext(ctx, sqlBulkInsertReport, data)
	if err != nil {
		return err
	}
	return tx.Commit()
}
