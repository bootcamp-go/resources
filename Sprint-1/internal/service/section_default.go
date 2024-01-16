package service

import "github.com/usuario/repositorio/internal"

// NewSectionDefault creates a new instance of the section service
func NewSectionDefault(rp internal.SectionRepository) *SectionDefault {
	return &SectionDefault{
		rp: rp,
	}
}

// SectionDefault is the default implementation of the section service
type SectionDefault struct {
	// rp is the repository used by the service
	rp internal.SectionRepository
}

// FindAll returns all sections
func (s *SectionDefault) FindAll() (sections []internal.Section, err error) {
	return
}

// FindByID returns a section
func (s *SectionDefault) FindByID(id int) (section internal.Section, err error) {
	return
}

// Save creates a new section
func (s *SectionDefault) Save(section *internal.Section) (err error) {
	return
}

// Update updates a section
func (s *SectionDefault) Update(section *internal.Section) (err error) {
	return
}

// Delete deletes a section
func (s *SectionDefault) Delete(id int) (err error) {
	return
}