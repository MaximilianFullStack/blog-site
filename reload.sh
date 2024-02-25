# /bin/bash

wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run main.go & \
yarn browser-sync start \
  --files './**/*.go, ./**/*.templ' \
  --ignore '*_templ.go' \
  --port 3001 \
  --proxy 'localhost:3000' \
  --reloadThrottle 50 \
  --middleware 'function(req, res, next) { \
    res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
    return next(); \
  }'

# nodemon
# yarn nodemon -e go,templ --ignore '*_templ.go' --exec 'templ generate && go run .'