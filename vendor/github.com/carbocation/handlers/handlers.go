// Copyright 2013 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package handlers is a collection of handlers for use with Go's net/http package.
*/
package handlers

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// MethodHandler is an http.Handler that dispatches to a handler whose key in the MethodHandler's
// map matches the name of the HTTP request's method, eg: GET
//
// If the request's method is OPTIONS and OPTIONS is not a key in the map then the handler
// responds with a status of 200 and sets the Allow header to a comma-separated list of
// available methods.
//
// If the request's method doesn't match any of its keys the handler responds with
// a status of 406, Method not allowed and sets the Allow header to a comma-separated list
// of available methods.
type MethodHandler map[string]http.Handler

func (h MethodHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := h[req.Method]; ok {
		handler.ServeHTTP(w, req)
	} else {
		allow := []string{}
		for k := range h {
			allow = append(allow, k)
		}
		sort.Strings(allow)
		w.Header().Set("Allow", strings.Join(allow, ", "))
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// loggingHandler is the http.Handler implementation for LoggingHandlerTo and its friends
type loggingHandler struct {
	writer  io.Writer
	handler http.Handler
}

// combinedLoggingHandler is the http.Handler implementation for LoggingHandlerTo and its friends
type combinedLoggingHandler struct {
	writer  io.Writer
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	var logger loggingResponseWriter
	if _, ok := w.(http.Hijacker); ok {
		logger = &hijackLogger{responseLogger: responseLogger{w: w}}
	} else {
		logger = &responseLogger{w: w}
	}
	h.handler.ServeHTTP(logger, req)
	writeLog(h.writer, req, t, logger.Status(), logger.Size())
}

func (h combinedLoggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	var logger loggingResponseWriter
	if _, ok := w.(http.Hijacker); ok {
		logger = &hijackLogger{responseLogger: responseLogger{w: w}}
	} else {
		logger = &responseLogger{w: w}
	}
	h.handler.ServeHTTP(logger, req)
	writeCombinedLog(h.writer, req, t, logger.Status(), logger.Size())
}

type loggingResponseWriter interface {
	http.ResponseWriter
	Status() int
	Size() int
}

// responseLogger is wrapper of http.ResponseWriter that keeps track of its HTTP status
// code and body size
type responseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (l *responseLogger) Header() http.Header {
	return l.w.Header()
}

func (l *responseLogger) Write(b []byte) (int, error) {
	if l.status == 0 {
		// The status will be StatusOK if WriteHeader has not been called yet
		l.status = http.StatusOK
	}
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *responseLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}

func (l *responseLogger) Status() int {
	return l.status
}

func (l *responseLogger) Size() int {
	return l.size
}

type hijackLogger struct {
	responseLogger
}

func (l *hijackLogger) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h := l.responseLogger.w.(http.Hijacker)
	conn, rw, err := h.Hijack()
	if err == nil && l.responseLogger.status == 0 {
		// The status will be StatusSwitchingProtocols if there was no error and WriteHeader has not been called yet
		l.responseLogger.status = http.StatusSwitchingProtocols
	}
	return conn, rw, err
}

// accounts for reverse proxies such as that by nginx
func buildRemoteAddr(req *http.Request) string {
	if x := req.Header.Get("X-Forwarded-For"); x != `` {
		return x
	}

	return req.RemoteAddr
}

// buildCommonLogLine builds a log entry for req in Apache Common Log Format.
// ts is the timestamp with which the entry should be logged.
// status and size are used to provide the response HTTP status and size.
func buildCommonLogLine(w io.Writer, req *http.Request, ts time.Time, status int, size int) {
	username := "-"
	if req.URL.User != nil {
		if name := req.URL.User.Username(); name != "" {
			username = name
		}
	}

	host, _, err := net.SplitHostPort(buildRemoteAddr(req))

	if err != nil {
		host = buildRemoteAddr(req)
	}

	io.WriteString(w, host+" - "+username+" ["+ts.Format("02/Jan/2006:15:04:05 -0700")+`] "`+req.Method+" "+req.URL.RequestURI()+" "+req.Proto+`" `+strconv.Itoa(status)+" "+strconv.Itoa(size))
}

// writeLog writes a log entry for req to w in Apache Common Log Format.
// ts is the timestamp with which the entry should be logged.
// status and size are used to provide the response HTTP status and size.
func writeLog(w io.Writer, req *http.Request, ts time.Time, status, size int) {
	buildCommonLogLine(w, req, ts, status, size)
	io.WriteString(w, "\n")
}

// writeCombinedLog writes a log entry for req to w in Apache Combined Log Format.
// ts is the timestamp with which the entry should be logged.
// status and size are used to provide the response HTTP status and size.
func writeCombinedLog(w io.Writer, req *http.Request, ts time.Time, status, size int) {
	buildCommonLogLine(w, req, ts, status, size)
	io.WriteString(w, ` "`+req.Referer()+`" "`+req.UserAgent()+`"`+"\n")
}

// CombinedLoggingHandler return a http.Handler that wraps h and logs requests to out in
// Apache Combined Log Format.
//
// See http://httpd.apache.org/docs/2.2/logs.html#combined for a description of this format.
//
// LoggingHandler always sets the ident field of the log to -
func CombinedLoggingHandler(out io.Writer, h http.Handler) http.Handler {
	return combinedLoggingHandler{out, h}
}

// LoggingHandler return a http.Handler that wraps h and logs requests to out in
// Apache Common Log Format (CLF).
//
// See http://httpd.apache.org/docs/2.2/logs.html#common for a description of this format.
//
// LoggingHandler always sets the ident field of the log to -
func LoggingHandler(out io.Writer, h http.Handler) http.Handler {
	return loggingHandler{out, h}
}
