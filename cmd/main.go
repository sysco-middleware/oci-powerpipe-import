package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/sysco-middleware/oci-powerpipe-import/db"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service, s, err := db.New()
	if err != nil {
		log.Fatalf("msg: %s , err : %v \n", s, err)
	}
	ctx := context.Background()
	id, err := service.GetRunID(ctx)
	if err != nil {
		log.Fatalf("msg: error.getting.run.id , err : %v \n", err)
	}

	log.Printf("id is %d \n", id)

	records, err := parseCsv("/Users/prakhar/dashboards/oci_compliance.benchmark.cis_v200.20240708T150811.csv", id)
	if err != nil {
		log.Fatalf("msg: error.parsing.csv , err : %v \n", err)
	}

	err = service.BulkInsert(ctx, records)
	if err != nil {
		log.Fatalf("msg: error.inserting.bulk.records , err : %v \n", err)
	}
}

func parseCsv(filepath string, runID int64) ([]db.Report, error) {

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var records []db.Report
	for {

		record, err := reader.Read()
		// if we've reached the end of the file, break
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		records = append(records, db.Report{
			RunID:              runID,
			GroupID:            record[0],
			Title:              record[1],
			Description:        parseSqlString(record[2]),
			ControlID:          record[3],
			ControlTitle:       record[4],
			ControlDescription: record[5],
			Reason:             record[6],
			Resource:           record[7],
			Status:             record[8],
			Severity:           parseSqlString(record[9]),
			Compartment:        parseSqlString(record[10]),
			Name:               parseSqlString(record[11]),
			Region:             parseSqlString(record[12]),
			ReportingRegion:    parseSqlString(record[13]),
			Tenant:             record[14],
			Category:           record[15],
			Cis:                parseBoolean(record[16]),
			CisItemID:          parseFloat(record[17]),
			CisLevel:           parseInt(record[18]),
			CisSectionID:       parseInt(record[19]),
			CisType:            record[20],
			CisVersion:         record[21],
			Plugin:             record[22],
			Service:            record[23],
			Type:               parseSqlString(record[24]),
			Created:            time.Now(),
		})
	}

	return records[1:], nil

}

func parseSqlString(str string) sql.NullString {
	ss := sql.NullString{}
	err := ss.Scan(str)
	if err != nil {
		return sql.NullString{}
	}
	return ss
}

func parseBoolean(str string) bool {
	if strings.TrimSpace(str) == "true" {
		return true
	}
	return false
}

func parseFloat(str string) float64 {
	ff, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return ff
}
func parseInt(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
