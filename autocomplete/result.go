package autocomplete

import (
	"io"
	"os"

	"log/slog"

	"github.com/adrg/xdg"
)

const (
	cacheFilePath = "clac/clac.cache"
)

func CacheResult(result string) {
	path, err := xdg.CacheFile(cacheFilePath)
	if err != nil {
		slog.Error("unable to create cache file", slog.String("error", err.Error()))
		return
	}
	f, err := os.Create(path)
	if err != nil {
		slog.Error("unable to open cache file", slog.String("error", err.Error()))
		return
	}
	if _, err := f.WriteString(result); err != nil {
		slog.Error("unable to write to cache file", slog.String("error", err.Error()))
	}
}

func ReadResult() string {
	path, err := xdg.CacheFile(cacheFilePath)
	if err != nil {
		slog.Error("unable to create cache file", slog.String("error", err.Error()))
		return ""
	}
	f, err := os.Open(path)
	if err != nil {
		slog.Error("unable to open cache file", slog.String("error", err.Error()))
		return ""
	}
	content, err := io.ReadAll(f)
	if err != nil {
		slog.Error("unable to read from cache file", slog.String("error", err.Error()))
		return ""
	}
	return string(content)
}