module gopackage

go 1.14

require (
	github.com/casbin/casbin v1.9.1
	github.com/casbin/gorm-adapter v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/jonboulle/clockwork v0.2.0 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/onsi/ginkgo v1.14.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.6.0
	github.com/tebeka/strftime v0.1.5 // indirect
	golang.org/x/text v0.3.3
	gopkg.in/yaml.v2 v2.3.0
)

replace gopackage => ./gopackage
