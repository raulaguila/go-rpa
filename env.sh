echo "
TWITTER_USER=''                                 # Twitter user
TWITTER_PASS=''                                 # Twitter pass

MONGO_USE='INT'                                 # Mongo HOST and PORT to use on application
MONGO_INT_HOST='80.30.1.52'                     # Mongo Container internal HOST
MONGO_EXT_HOST='127.0.0.1'                      # Mongo Container external HOST
MONGO_INT_PORT='27017'                          # Mongo Container internal PORT
MONGO_EXT_PORT='27018'                          # Mongo Container external PORT
MONGO_USER='admin'                              # Mongo USER
MONGO_PASS='password123'                        # Mongo PASS
MONGO_BASE='news'                               # Mongo BASE
" > .env