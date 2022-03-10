package container

var container Container

func GetContainer() Container {
	if container == nil {
		c, err := NewDigWrapper()
		if err != nil {
			panic(err)
		}

		container = c
	}

	return container
}
