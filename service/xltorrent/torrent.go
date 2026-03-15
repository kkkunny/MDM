package xltorrent

import (
	"os"
	"path/filepath"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/fsnotify/fsnotify"
	"github.com/kkkunny/stl/container/hashmap"
	stlmaps "github.com/kkkunny/stl/container/maps"
	stlerr "github.com/kkkunny/stl/error"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/util"
)

var TorrentsCache = stlerr.MustWith(NewTorrentsCache())

type _TorrentsCache struct {
	watcher *fsnotify.Watcher
	data    stlmaps.MapObj[string, *metainfo.MetaInfo]
}

func NewTorrentsCache() (*_TorrentsCache, error) {
	watcher, err := stlerr.ErrorWith(fsnotify.NewWatcher())
	if err != nil {
		return nil, err
	}
	err = stlerr.ErrorWrap(watcher.Add(config.XLBtDir))
	if err != nil {
		return nil, err
	}
	tc := &_TorrentsCache{
		watcher: watcher,
		data:    hashmap.ThreadSafeStdWith[string, *metainfo.MetaInfo](),
	}
	return tc, tc.init()
}

func (tc *_TorrentsCache) init() error {
	fileEntries, err := stlerr.ErrorWith(os.ReadDir(config.XLBtDir))
	if err != nil {
		return err
	}
	for _, fileEntry := range fileEntries {
		fp := filepath.Join(config.XLBtDir, fileEntry.Name())
		mi, err := stlerr.ErrorWith(metainfo.LoadFromFile(fp))
		if err != nil {
			_ = config.Logger.Warn(err)
			continue
		}
		tc.data.Set(fp, mi)
	}

	go func() {
		for {
			select {
			case event, ok := <-tc.watcher.Events:
				if !ok {
					return
				}
				_ = config.Logger.Debug(util.ToJson[string](event))
				var err error
				switch {
				case event.Has(fsnotify.Create):
					err = tc.onCreate(event.Name)
				case event.Has(fsnotify.Remove):
					err = tc.onRemove(event.Name)
				}
				if err != nil {
					_ = config.Logger.Error(err)
				}
			case err, ok := <-tc.watcher.Errors:
				if !ok {
					return
				}
				_ = config.Logger.Error(stlerr.ErrorWrap(err))
			}
		}
	}()

	return nil
}

func (tc *_TorrentsCache) Close() error {
	return stlerr.ErrorWrap(tc.watcher.Close())
}

func (tc *_TorrentsCache) Get() stlmaps.MapObj[string, *metainfo.MetaInfo] {
	return tc.data
}

func (tc *_TorrentsCache) onCreate(path string) error {
	mi, err := stlerr.ErrorWith(metainfo.LoadFromFile(path))
	if err != nil {
		return err
	}
	tc.data.Set(path, mi)
	return nil
}

func (tc *_TorrentsCache) onRemove(path string) error {
	tc.data.Remove(path)
	return nil
}
