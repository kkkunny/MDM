package xltorrent

import (
	"maps"
	"os"
	"path/filepath"
	"sync"

	"github.com/anacrolix/torrent/metainfo"
	stlerr "github.com/kkkunny/stl/error"
	stlsync "github.com/kkkunny/stl/sync"

	"github.com/kkkunny/MDM/config"
)

var TorrentsCache = stlerr.MustWith(NewTorrentsCache())

type _TorrentsCache struct {
	locker stlsync.RWLocker
	data   map[string]*metainfo.MetaInfo
}

func NewTorrentsCache() (*_TorrentsCache, error) {
	tc := &_TorrentsCache{
		locker: new(sync.RWMutex),
		data:   make(map[string]*metainfo.MetaInfo),
	}
	return tc, tc.Scan()
}

func (tc *_TorrentsCache) Scan() error {
	tc.locker.Lock()
	defer tc.locker.Unlock()

	tc.data = make(map[string]*metainfo.MetaInfo)

	fileEntries, err := stlerr.ErrorWith(os.ReadDir(config.XLBtDir))
	if err != nil {
		return err
	}
	current := make(map[string]struct{}, len(fileEntries))
	for _, fileEntry := range fileEntries {
		fp := filepath.Join(config.XLBtDir, fileEntry.Name())
		current[fp] = struct{}{}
		mi, err := stlerr.ErrorWith(metainfo.LoadFromFile(fp))
		if err != nil {
			_ = config.Logger.Warn(err)
			continue
		}
		tc.data[fp] = mi
	}
	return nil
}

func (tc *_TorrentsCache) Get() map[string]*metainfo.MetaInfo {
	tc.locker.RLock()
	defer tc.locker.RUnlock()
	return maps.Clone(tc.data)
}
