// internal/app/service.go
package translate

type IService interface {
	GetTranslation(sourceLang, targetLang, word, pos string) (string, error)
}
