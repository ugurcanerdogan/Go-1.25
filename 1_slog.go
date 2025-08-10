package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Old method: Direct attributes with Group
	logger.Info("content cannot be bought",
		slog.Bool("isSellable", false),
		slog.Group("price",
			slog.Int("value", 0),
			slog.String("currency", "TRY"),
		),
	)

	// In the old method, using (slice...) is not possible:
	attrs := []slog.Attr{
		slog.Int("value", 1000),
		slog.String("currency", "TRY"),
	}
	logger.Info("content cannot be bought",
		slog.Bool("isSellable", false),
		slog.Group("price", attrs...), // delete triple dot
	) // Compile error!

	// New method with Go 1.25:
	// GroupAttrs is a more efficient version of [Group]
	// that accepts only [Attr] values.

	logger.Info("content cannot be bought",
		slog.Bool("isSellable", false),
		slog.GroupAttrs("price", attrs...),
	)
}
