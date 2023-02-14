package loggingx

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

// GRPCLogger wrapps a hclog.Logger and implements the grpclog.LoggerV2 interface
// allowing gRPC servers to log to the standard logger.
type GRPCLogger struct {
	level  string
	logger hclog.Logger
}

// NewGRPCLogger creates a grpclog.LoggerV2 that will output to the supplied
// logger with Severity/Verbosity level appropriate for the given config.
//
// Note that grpclog has Info, Warning, Error, Fatal severity levels AND integer
// verbosity levels for additional info. Verbose logs in hclog are always DEBUG
// severity so we map Info,V0 to INFO,V1 to DEBUG, and Info,V>1 to TRACE.
func NewGRPCLogger(logLevel string, logger hclog.Logger) *GRPCLogger {
	return &GRPCLogger{
		level:  logLevel,
		logger: logger,
	}
}

func (g *GRPCLogger) Info(args ...interface{}) {
	// gRPC's INFO level is more akin to TRACE level
	g.logger.Trace(fmt.Sprint(args...))
}

func (g *GRPCLogger) Infoln(args ...interface{}) {
	g.Info(fmt.Sprint(args...))
}

func (g *GRPCLogger) Infof(format string, args ...interface{}) {
	g.Info(fmt.Sprintf(format, args...))
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (g *GRPCLogger) Warning(args ...interface{}) {
	g.logger.Warn(fmt.Sprint(args...))
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (g *GRPCLogger) Warningln(args ...interface{}) {
	g.Warning(fmt.Sprint(args...))
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (g *GRPCLogger) Warningf(format string, args ...interface{}) {
	g.Warning(fmt.Sprintf(format, args...))
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (g *GRPCLogger) Error(args ...interface{}) {
	g.logger.Error(fmt.Sprint(args...))
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (g *GRPCLogger) Errorln(args ...interface{}) {
	g.Error(fmt.Sprint(args...))
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (g *GRPCLogger) Errorf(format string, args ...interface{}) {
	g.Error(fmt.Sprintf(format, args...))
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (g *GRPCLogger) Fatal(args ...interface{}) {
	g.logger.Error(fmt.Sprint(args...))
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (g *GRPCLogger) Fatalln(args ...interface{}) {
	g.Fatal(fmt.Sprint(args...))
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (g *GRPCLogger) Fatalf(format string, args ...interface{}) {
	g.Fatal(fmt.Sprintf(format, args...))
}

// V reports whether verbosity level l is at least the requested verbose level.
func (g *GRPCLogger) V(l int) bool {
	switch g.level {
	case "TRACE":
		// Enable ALL the verbosity!
		return true
	case "DEBUG":
		return l < 2
	case "INFO":
		return l < 1
	default:
		return false
	}
}
