package vpn

type Vpnglobalauthenticationlocalpolicybinding struct {
	Groupextraction bool   `json:"groupextraction,omitempty"`
	Policyname      string `json:"policyname,omitempty"`
	Priority        int    `json:"priority,omitempty"`
	Secondary       bool   `json:"secondary,omitempty"`
}
