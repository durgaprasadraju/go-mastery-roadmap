package main

import (
	"context"
	"os"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/111-clean-architecture/examples"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

func main() {
	log := logger.Must("111-clean-architecture")
	ctx := logger.WithContext(context.Background(), log)

	if err := examples.BasicDemo(ctx); err != nil {
		log.Error("basic demo failed")
		os.Exit(1)
	}
	if err := examples.AdvancedDemo(ctx); err != nil {
		log.Error("advanced demo failed")
		os.Exit(1)
	}
}
