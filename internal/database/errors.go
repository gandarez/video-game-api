package database

// Level is a logging priority. Higher levels are more important.
type Level uint32

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in production.
	DebugLevel = iota
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

// ErrOpenTransaction is returned when transaction is not opened.
type ErrOpenTransaction string

// Error method to implement error interface.
func (e ErrOpenTransaction) Error() string {
	return string(e)
}

// Level method to implement LogLevel interface.
func (ErrOpenTransaction) Level() Level {
	return ErrorLevel
}

// ErrRollbackTransaction is returned when transaction failed to rollback.
type ErrRollbackTransaction string

// Error method to implement error interface.
func (e ErrRollbackTransaction) Error() string {
	return string(e)
}

// Level method to implement LogLevel interface.
func (ErrRollbackTransaction) Level() Level {
	return ErrorLevel
}

// ErrCommitTransaction is returned when transaction failed to commit.
type ErrCommitTransaction string

// Error method to implement error interface.
func (e ErrCommitTransaction) Error() string {
	return string(e)
}

// Level method to implement LogLevel interface.
func (ErrCommitTransaction) Level() Level {
	return ErrorLevel
}

// ErrNotFound is returned when no rows are found.
type ErrNotFound string

// Error method to implement error interface.
func (e ErrNotFound) Error() string {
	return string(e)
}

// Level method to implement LogLevel interface.
func (ErrNotFound) Level() Level {
	return WarnLevel
}

// ErrConflict is returned when a conflict or duplicate entry is found.
type ErrConflict string

// Error method to implement error interface.
func (e ErrConflict) Error() string {
	return string(e)
}

// Level method to implement LogLevel interface.
func (ErrConflict) Level() Level {
	return ErrorLevel
}
