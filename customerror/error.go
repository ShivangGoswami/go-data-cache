package customerror

import "errors"

var MongoStoreException = errors.New("Error writing to mongo db")
var TimeParseError = errors.New("Error parsing Duration")
var InvalidInput = errors.New("Either key or index is a mandatory input")
var MongoNotFound = errors.New("Cache key not found in mongo")
