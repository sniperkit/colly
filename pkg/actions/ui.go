package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/kk/pkg/kk"
)

// ToggleViewDebug returns a function that will toggle the debug view
func ToggleViewDebug(s *k.State) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenDebug {
			s.UI.SetActiveScreen(g, k.ScreenTable)
			return nil
		}
		s.UI.SetActiveScreen(g, k.ScreenDebug)
		return nil

	}
}

// ToggleResources returns a function that will toggle the resources modal
func ToggleResources(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenModal {
			s.UI.SetActiveScreen(g, k.ScreenTable)
			return nil
		}
		k.Debugln("Toggle: resources")
		lines := s.Entities.Resources.Resources
		s.UI.SetActiveScreen(g, k.ScreenModal)
		s.UI.Modal.SetKind(g, k.KindModalResources)
		s.UI.Modal.SetLines(g, lines)
		s.UI.Modal.SetTitle(g, "Resources")
		s.UI.Modal.SetSize(g, k.ModalSizeMedium)
		return nil

	}
}

// ToggleNamespaces returns a function that will toggle the namespaces modal
func ToggleNamespaces(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenModal {
			s.UI.SetActiveScreen(g, k.ScreenTable)
			return nil
		}
		k.Debugln("Toggle: namespaces")
		lines := []string{}
		for _, ns := range s.Entities.Namespaces.Namespaces.Items {
			lines = append(lines, ns.ObjectMeta.Name)
		}
		s.UI.SetActiveScreen(g, k.ScreenModal)
		s.UI.Modal.SetKind(g, k.KindModalNamespaces)
		s.UI.Modal.SetLines(g, lines)
		s.UI.Modal.SetTitle(g, "Namespaces")
		s.UI.Modal.SetSize(g, k.ModalSizeMedium)
		return nil
	}
}
