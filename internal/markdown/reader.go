package markdown

type FileReader interface {
	Read(target string) ([]byte, error)
}
