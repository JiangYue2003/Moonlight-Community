package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	knowpostmodel "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	relationmodel "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
)

func main() {
	var (
		datasource    = flag.String("datasource", os.Getenv("MYSQL_DATASOURCE"), "MySQL data source name")
		afterAuthorID = flag.Int64("after-author-id", 0, "resume after this author ID")
		batchSize     = flag.Int("batch-size", 500, "authors per page (1-10000)")
		maxPages      = flag.Int("max-pages", 100, "maximum pages processed in this run")
		pageInterval  = flag.Duration("page-interval", 100*time.Millisecond, "delay between pages")
		apply         = flag.Bool("apply", false, "persist promotions; default is dry-run")
	)
	flag.Parse()

	if *datasource == "" {
		fmt.Fprintln(os.Stderr, "-datasource or MYSQL_DATASOURCE is required")
		os.Exit(2)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	conn := sqlx.NewMysql(*datasource)
	report, err := runTierBackfill(
		ctx,
		relationmodel.NewActiveFollowerCountModel(conn),
		knowpostmodel.NewFeedAuthorDeliveryTierModel(conn),
		tierBackfillOptions{
			AfterAuthorID: *afterAuthorID,
			BatchSize:     *batchSize,
			MaxPages:      *maxPages,
			Apply:         *apply,
			PageInterval:  *pageInterval,
		},
	)
	if encodeErr := json.NewEncoder(os.Stdout).Encode(report); encodeErr != nil {
		fmt.Fprintf(os.Stderr, "encode report: %v\n", encodeErr)
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
