package pkg

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"github.com/skratchdot/open-golang/open"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"reserver/utils"
	"strings"
	"time"
)

type reServer struct {
	log           utils.Log
	args          *utils.CmdArgs
	wsUpgrade     *websocket.Upgrader
	wsConnections []*websocket.Conn
}

func NewReServer(args *utils.CmdArgs) *reServer {
	return &reServer{
		args: args,
		log:  utils.NewLog(),
		wsUpgrade: &websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				if r.Method != "GET" || r.URL.Path != "/ws" {
					return false
				}
				return true
			},
		},
	}
}

func (l *reServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	p := request.URL.Path
	if p == "/" {
		p = "/index.html"
	}
	if p == "/ws" {
		l.WsHandler(w, request)
		return
	}
	body, err := os.ReadFile(fmt.Sprintf("%s/%s", l.args.Root, p))
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	contentType, b := utils.GetFileContentTypeBySuffix(p)
	if !b {
		contentType = " "
	}
	body, _ = l.joinWsReloadScript(body)
	w.Header().Set("Content-Type", contentType)
	_, _ = w.Write(body)
}

func (l *reServer) WsHandler(w http.ResponseWriter, request *http.Request) {
	conn, err := l.wsUpgrade.Upgrade(w, request, nil)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	l.wsConnections = append(l.wsConnections, conn)
	err = WriteMsg("hi", conn)
	if err != nil {
		l.log.Warning("write 'hi' string to client failed", err)
	}
}

func (l *reServer) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", l.args.Host, l.args.Port))
	if err != nil {
		panic(err)
	}
	host := l.args.Host
	if host == "0.0.0.0" {
		host = "127.0.0.1"
	}
	url := fmt.Sprintf("http://%s:%d", host, l.args.Port)
	if l.args.Path != "" {
		url += l.args.Path
	}
	l.log.Success("Start successful, at %s", url)
	if !l.args.NoBrowser {
		if l.args.Browser != "" {
			err = open.RunWith(url, l.args.Browser)
		} else {
			err = open.Run(url)
		}
		if err != nil {
			l.log.Warning("open browser failed")
		}
	}
	if !l.args.NoWatch {
		l.watch(l.args.Root)
	}
	err = http.Serve(listen, l)
	if err != nil {
		panic(err)
	}
}

func (l *reServer) joinWsReloadScript(body []byte) (res []byte, length int) {
	tag := "</body>"
	match, err := regexp.Match(tag, body)
	if err != nil {
		return body, len(body)
	}
	if match {
		bodyString := string(body)
		replace := strings.Replace(bodyString, tag, GetInjectScript(tag), 1)
		res = []byte(replace)
		return res, len(res)
	}
	return body, len(body)
}

func (l *reServer) watch(root string) {
	watcher, err := NewFsWatcher(time.Duration(l.args.Wait) * time.Millisecond)
	if err != nil {
		l.log.Error("watch root dir failed")
		return
	}
	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			select {
			case events, ok := <-watcher.Events():
				if !ok {
					return
				}
				t := RELOAD
				var hasCss, moreThanCss bool
				for _, e := range events {
					if !hasCss {
						hasCss = path.Ext(e.Name) == ".css"
						if !hasCss && !moreThanCss {
							moreThanCss = true
						}
					}
					if !l.args.HideLog {
						l.log.DD("[%s] %s", e.Op, e.Name)
					}
					switch e.Op {
					case fsnotify.Create:
						{
							stat, terr := os.Stat(e.Name)
							if terr == nil && stat.IsDir() {
								_ = watcher.Add(e.Name)
							}
						}
					case fsnotify.Remove:
						{
							stat, terr := os.Stat(e.Name)
							if terr == nil && stat.IsDir() {
								_ = watcher.Remove(e.Name)
							}
						}
					}
				}
				if !l.args.CssReload && hasCss && !moreThanCss {
					t = CSS
				}
				WriteMsgToManyConn(t, l.wsConnections)
			}
		}
	}()
	err = watcher.Add(root)
	if err != nil {
		l.log.Error("watch dir failed: %s", root)
	} else {
		l.log.Success("watch dir success: %s", root)
	}
	l.watchDir(root, watcher)
}

func (l *reServer) watchDir(dirPath string, watcher FsWatcher) {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		l.log.Error("read root dir failed")
		_ = watcher.Close()
		return
	}
	for _, entry := range dir {
		if entry.IsDir() {
			p := filepath.ToSlash(filepath.Join(dirPath, entry.Name()))
			if _, has := l.args.IgnoreMap[p]; has {
				continue
			}
			err = watcher.Add(p)
			if err != nil {
				l.log.Error("watch dir failed: %s", p)
			} else {
				l.log.Success("watch dir success: %s", p)
			}
			l.watchDir(p, watcher)
		}
	}
}
