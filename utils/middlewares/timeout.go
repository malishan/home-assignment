package middlewares

// // Timeout wraps the request context with a timeout
// func Timeout(timeout time.Duration, errTimeout *apperrors.Error) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// set Gin's writer as our custom writer
// 		tw := &timeoutWriter{ResponseWriter: c.Writer, h: make(http.Header)}
// 		c.Writer = tw

// 		// wrap the request context with a timeout
// 		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
// 		defer cancel()

// 		// update gin request context
// 		c.Request = c.Request.WithContext(ctx)

// 		finished := make(chan struct{})        // to indicate handler finished
// 		panicChan := make(chan interface{}, 1) // used to handle panics if we can't recover

// 		go func() {
// 			defer func() {
// 				if p := recover(); p != nil {
// 					panicChan <- p
// 				}
// 			}()

// 			c.Next() // calls subsequent middleware(s) and handler
// 			finished <- struct{}{}
// 		}()

// 		select {
// 		case <-panicChan:
// 			// if we cannot recover from panic,
// 			// send internal server error
// 			e := apperrors.NewInternal()
// 			tw.ResponseWriter.WriteHeader(e.Status())
// 			eResp, _ := json.Marshal(gin.H{
// 				"error": e,
// 			})
// 			tw.ResponseWriter.Write(eResp)
// 		case <-finished:
// 			// if finished, set headers and write resp
// 			tw.mu.Lock()
// 			defer tw.mu.Unlock()
// 			// map Headers from tw.Header() (written to by gin)
// 			// to tw.ResponseWriter for response
// 			dst := tw.ResponseWriter.Header()
// 			for k, vv := range tw.Header() {
// 				dst[k] = vv
// 			}
// 			tw.ResponseWriter.WriteHeader(tw.code)
// 			// tw.wbuf will have been written to already when gin writes to tw.Write()
// 			tw.ResponseWriter.Write(tw.wbuf.Bytes())
// 		case <-ctx.Done():
// 			// timeout has occurred, send errTimeout and write headers
// 			tw.mu.Lock()
// 			defer tw.mu.Unlock()
// 			// ResponseWriter from gin
// 			tw.ResponseWriter.Header().Set("Content-Type", "application/json")
// 			tw.ResponseWriter.WriteHeader(errTimeout.Status())
// 			eResp, _ := json.Marshal(gin.H{
// 				"error": errTimeout,
// 			})
// 			tw.ResponseWriter.Write(eResp)
// 			c.Abort()
// 			tw.SetTimedOut()
// 		}
// 	}
// }
