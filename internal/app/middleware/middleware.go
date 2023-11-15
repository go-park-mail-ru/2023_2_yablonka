package middleware

import "github.com/sirupsen/logrus"

func middlewareDebugLog(logger *logrus.Logger, function string, message string) {
	logger.
		WithFields(logrus.Fields{
			"route_node": "middleware",
			"function":   function,
		}).
		Debug(message)
}
