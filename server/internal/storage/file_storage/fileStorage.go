package file_storage

import (
	"log/slog"
	"os"
	"server/internal/lib/logger/sl"
)

func InitFileStorage(storagePath string, logger *slog.Logger) {
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		err := os.Mkdir(storagePath, 0755)
		if err != nil {
			logger.Error("Can't create storage directory: %s", sl.Err(err))
		}
	}
}
