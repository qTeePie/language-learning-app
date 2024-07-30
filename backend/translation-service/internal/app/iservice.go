// internal/app/service.go
package app

type IService interface {
	GetTranslation(sourceLang, targetLang, word, pos string) (string, error)
}
