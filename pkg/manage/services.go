package manage

func (m *Manage) GetServices() error {
	services, err := m.Lister.ListS()
	if err != nil {
		return err
	}
	for _, s := range services {
		m.Incidents[s.Name] = nil
	}
	return nil
}
