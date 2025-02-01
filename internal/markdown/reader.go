package markdown

// FileReader t.b.d. until API stable
type FileReader interface {
	Read(target string) ([]byte, error)
	List(target string) ([]string, error)
}
