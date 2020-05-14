package removing

type store interface {
	RemoveTil(string) error
}

func Run(s store, uuid string) error {
	err := s.RemoveTil(uuid)
	if err != nil {
		return err
	}

	return nil
}
