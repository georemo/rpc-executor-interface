/*
Testing RPC
By G. Oremo
10th March 2024
*/

// package shared contains the shared interface definition
package shared

// CdExecServer represents the server interface for executing commands.
type CdExecServer interface {
	// CdExec is a method that executes a command and returns the result.
	CdExec(jsonInput string) (string, error)
}

// CdExecClient represents the client interface for executing commands.
type CdExecClient interface {
	// CdExec is a method that executes a command and returns the result.
	CdExec(jsonInput string) (string, error)
}
