module github.com/lxxccc/tfxingApiProxy

go 1.12

require (
	github.com/astaxie/beego v1.10.0
	github.com/debber/debber-v0.3 v0.0.0-20160226200915-cf10fd052975 // indirect
	github.com/laher/argo v0.0.0-20140722103944-11d91c83cc0f // indirect
	github.com/laher/goxc v0.18.1 // indirect
	golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a
	gopkg.in/libgit2/git2go.v27 v27.0.0-20190104134018-ecaeb7a21d47
	gopkg.in/yaml.v2 v2.2.2
	lxxccc.top/Library/ToolkitsGo v0.1.6
)

replace (
	golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	lxxccc.top/Library/ToolkitsGo v0.1.6 => home.lxxccc.top/Library/ToolkitsGo.git v0.1.6
)
