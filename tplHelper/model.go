package tplhelper

import "text/template"

type tplData struct {
	Template         *template.Template `json:"template"`            // template
	LocationPath     string             `json:"location_path"`       // tpl location
	TempPath         string             `json:"temp_path"`           // temporary path location, will need this cause zip feature, we need to put it in same folder to zip
	AutoMoveFilePath string             `json:"auto_move_file_path"` // if auto move, file will move to this path
}
