package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (resultError error) {
	openedResource, err := o()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return err
	}
	defer openedResource.Close()
	defer func() {
		if err := recover(); err != nil {
			if ferr, ok := err.(FrobError); ok {
				openedResource.Defrob(ferr.defrobTag)
			}
			resultError = err.(error)
		}
	}()
	openedResource.Frob(input)
	return nil
}