package injection

func HTTPServer() {
	CompositionRoot{}.HTTPServer()
}

type CompositionRoot struct{}

func (self CompositionRoot) HTTPServer() {

}
