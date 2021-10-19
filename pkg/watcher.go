package pkg

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"time"
)

type FsWatcher interface {
	Close() error
	Add(name string) error
	Remove(name string) error
	Events() chan map[string]fsnotify.Event
}

type fsWatcher struct {
	watcher  *fsnotify.Watcher
	events   chan map[string]fsnotify.Event
	done     chan struct{}
	interval time.Duration
}

func NewFsWatcher(interval time.Duration) (res FsWatcher, err error) {
	fs, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	fw := &fsWatcher{
		interval: interval,
		events:   make(chan map[string]fsnotify.Event, 1),
		done:     make(chan struct{}, 1),
		watcher:  fs,
	}
	go fw.run()
	res = fw
	return
}

func (f *fsWatcher) run() {
	tick := time.Tick(f.interval)
	events := make(map[string]fsnotify.Event, 0)
	for {
		select {
		case e, ok := <-f.watcher.Events:
			if !ok {
				return
			}
			key := fmt.Sprintf("%s-%v", e.Name, e.Op)
			_, has := events[key]
			if !has {
				events[key] = e
			}
		case <-tick:
			if len(events) == 0 {
				continue
			}
			f.events <- events
			events = make(map[string]fsnotify.Event, 0)
		case <-f.done:
			return
		}
	}
}

func (f *fsWatcher) Close() error {
	f.done <- struct{}{}
	return f.watcher.Close()
}

func (f *fsWatcher) Add(name string) error {
	return f.watcher.Add(name)
}

func (f *fsWatcher) Remove(name string) error {
	return f.watcher.Remove(name)
}

func (f *fsWatcher) Events() chan map[string]fsnotify.Event {
	return f.events
}
