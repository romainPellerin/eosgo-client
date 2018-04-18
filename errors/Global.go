/**
 *  @copyright defined in eosgo-client/LICENSE.txt
 *  @author Romain Pellerin - romain@slyseed.com
 *
 *  Donation appreciated :)
 *  EOS wallet: 0x8d39307d9a0687c894058115365f6ad0f03b9ff9
 *	ETH wallet: 0x317b60152f0a90c10cea52d751ccb4dfad2fe9e7
 *  BTC wallet: 3HdFx8C3WNA6RyyGYKygECGbLD1BxyeqTN
 *  BCH wallet: 14e9fnmcHSZxxd3fnrkaG6EYph9JuWcF9T
 */

package errors

type AppError struct {
	Error   error
	Message string
	Code    int
	Custom  interface{}
}

func NewAppError(error error, message string, code int64, custom interface{}) *AppError {
	return &AppError{
		error,
		message,
		int(code),
		custom,
	}
}

func MarshallingError(err error) *AppError {
	return NewAppError(err, "", -1, nil)
}

func UnsupportedOperation() *AppError {
	return &AppError{nil, "unsupported operation", 14, nil}
}
