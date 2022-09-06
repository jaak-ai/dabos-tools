package setup

import (
	"fmt"
	"gopkg.in/ini.v1"
	"reflect"
)

func Populate(model interface{}, path string) error {
	config, err := readFileINI(path)
	if err != nil {
		return err
	}

	vModel := getReflectValue(model)
	populateModel(vModel, config, "")

	return nil
}

func readFileINI(path string) (*ini.File, error) {
	config, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func getReflectValue(model interface{}) reflect.Value {
	vModel := reflect.ValueOf(model)

	if vModel.Type().Kind() == reflect.Ptr {
		vModel = vModel.Elem()
	}

	return vModel
}

func populateModel(vModel reflect.Value, config *ini.File, section string) {
	numField := vModel.NumField()

	for index := 0; index < numField; index++ {
		tField := vModel.Type().Field(index)

		tagList := tField.Tag
		tagCFG := tagList.Get("cfg")

		if tagCFG == "" || tagCFG == "-" {
			continue
		}

		if tField.Type.Kind() == reflect.Struct {
			vField := vModel.Field(index)
			var newSection string
			if section == "" {
				newSection = tagCFG
			} else {
				newSection = fmt.Sprintf("%s.%s", section, tagCFG)
			}

			populateModel(vField, config, newSection)
			continue
		}

		if tField.Type.Kind() == reflect.String {
			vConfig := config.Section(section).Key(tagCFG).String()
			vField := vModel.Field(index)

			if vField.CanSet() {
				vField.SetString(vConfig)
			}
		}

		if tField.Type.Kind() == reflect.Bool {
			vConfig, err := config.Section(section).Key(tagCFG).Bool()
			if err != nil {
				continue
			}

			vField := vModel.Field(index)
			if vField.CanSet() {
				vField.SetBool(vConfig)
			}
		}

		if tField.Type.Kind() == reflect.Int64 {
			vConfig, err := config.Section(section).Key(tagCFG).Int64()
			if err != nil {
				continue
			}

			vField := vModel.Field(index)
			if vField.CanSet() {
				vField.SetInt(vConfig)
			}
		}

		if tField.Type.Kind() == reflect.Float64 {
			vConfig, err := config.Section(section).Key(tagCFG).Float64()
			if err != nil {
				continue
			}

			vField := vModel.Field(index)
			if vField.CanSet() {
				vField.SetFloat(vConfig)
			}
		}
	}
}
