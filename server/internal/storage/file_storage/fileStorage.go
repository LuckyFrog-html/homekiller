package file_storage

import (
	"log/slog"
	"os"
	"path"
	"server/internal/lib/logger/sl"
)

func InitFileStorage(storagePath string, logger *slog.Logger) {
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		err := os.Mkdir(storagePath, 0755)
		if err != nil {
			logger.Error("Can't create storage directory: %s", sl.Err(err))
		}
	}
	if _, err := os.Stat(path.Join(storagePath, "teachers")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(storagePath, "teachers"), 0755)
		if err != nil {
			logger.Error("Can't create storage directory: %s", sl.Err(err))
		}
	}
	if _, err := os.Stat(path.Join(storagePath, "students")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(storagePath, "students"), 0755)
		if err != nil {
			logger.Error("Can't create storage directory: %s", sl.Err(err))
		}
	}
}
