package customerror

import "errors"

var MongoStoreException = errors.New("Error writing to mongo db")
var TimeParseError = errors.New("Error parsing Duration")
