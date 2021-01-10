package builder

type APIs []*API
type Libraries []*APILibrary

type API struct {
	Name          string   `yaml:"name"`
	Slug          string   `yaml:"slug"`
	Categories    []string `yaml:"categories"`
	Description   string   `yaml:"description"`
	URI           string   `yaml:"uri"`
	IsPaid        bool     `yaml:"is_paid"`
	Logo          string   `yaml:"icon"`
	DiscussionURI string   `yaml:"discussion_uri"`
	Type          string   `yaml:"type"`
	Contact       string   `yaml:"contact"`
	IsDead        bool     `yaml:"is_dead"`

	Libraries Libraries `yaml:"libraries"`
}

type APILibrary struct {
	Name          string `yaml:"name"`
	Description   string `yaml:"-"`
	HomepageURI   string `yaml:"homepage_uri" yaml:"homepage_uri"`
	SourceCodeURI string `yaml:"source_code_uri" yaml:"source_code_uri"`
	Version       string `yaml:"version"`
	Platform      string `yaml:"platform"`
}

func (a APIs) ByCategory() map[string][]*API {
	cm := make(map[string][]*API)
	for _, v := range a {
		for _, vv := range v.Categories {
			if _, ok := cm[vv]; !ok {
				cm[vv] = make([]*API, 0)
			}
			cm[vv] = append(cm[vv], v)
		}
	}

	return cm
}

func (a APIs) Graveyard() []*API {
	cm := make([]*API, 0)
	for _, v := range a {
		if v.IsDead == true {
			cm = append(cm, v)
		}
	}

	return cm
}

func (a Libraries) ByPlatform() map[string][]*APILibrary {
	cm := make(map[string][]*APILibrary)
	for _, v := range a {
		if _, ok := cm[v.Platform]; !ok {
			cm[v.Platform] = make([]*APILibrary, 0)
		}
		cm[v.Platform] = append(cm[v.Platform], v)
	}

	return cm
}
