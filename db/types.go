package db

import (
	"database/sql"
	"time"
)

type Report struct {
	ID                 string         `db:"id"`
	RunID              int64          `db:"run_id"`
	GroupID            string         `db:"group_id"`
	Title              string         `db:"title"`
	Description        sql.NullString `db:"description"`
	ControlID          string         `db:"control_id"`
	ControlTitle       string         `db:"control_title"`
	ControlDescription string         `db:"control_description"`
	Reason             string         `db:"reason"`
	Resource           string         `db:"resource"`
	Status             string         `db:"status"`
	Severity           sql.NullString `db:"severity"`
	Compartment        sql.NullString `db:"compartment"`
	Name               sql.NullString `db:"name"`
	Region             sql.NullString `db:"region"`
	ReportingRegion    sql.NullString `db:"reporting_region"`
	Tenant             string         `db:"tenant"`
	Category           string         `db:"category"`
	Cis                bool           `db:"cis"`
	CisItemID          float64        `db:"cis_item_id"`
	CisLevel           int64          `db:"cis_level"`
	CisSectionID       int64          `db:"cis_section_id"`
	CisType            string         `db:"cis_type"`
	CisVersion         string         `db:"cis_version"`
	Plugin             string         `db:"plugin"`
	Service            string         `db:"service"`
	Type               sql.NullString `db:"type"`
	Created            time.Time      `db:"created"`
}

type RunID struct {
	ID int64 `db:"id"`
}
