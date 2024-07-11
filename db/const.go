package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

const (
	sqlGetRunID         string = "SELECT id FROM public.run_id"
	sqlTruncateRunID    string = "TRUNCATE TABLE public.run_id"
	sqlInsertRunID      string = "INSERT INTO  public.run_id (id) VALUES ($1)"
	sqlBulkInsertReport string = `
	INSERT INTO public.report (
		run_id,	group_id,title, description, control_id, control_title,
		control_description, reason, resource, status, severity, compartment,
		name, region, reporting_region, tenant,
		category, cis, cis_item_id, cis_level,
		cis_section_id, cis_type, cis_version, plugin,
		service, type, created
	)
	VALUES (
		:run_id,	:group_id,:title, :description, :control_id, :control_title,
		:control_description, :reason, :resource, :status, :severity, :compartment,
		:name, :region, :reporting_region, :tenant,
		:category, :cis, :cis_item_id, :cis_level,
		:cis_section_id, :cis_type, :cis_version, :plugin,
		:service, :type, :created
	)
`
)

func HandleTxRollback(txx *sqlx.Tx) {
	err := txx.Rollback()
	if err != nil && err.Error() != sql.ErrTxDone.Error() {
		return
	}
}
