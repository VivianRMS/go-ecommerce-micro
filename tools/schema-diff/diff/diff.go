package diff

import (
	"fmt"
	"github.com/jhump/protoreflect/desc/protoparse"
)

func CompareProtoFiles(oldPath, newPath string) ([]string, error) {
	parser := protoparse.Parser{}

	oldFiles, err := parser.ParseFiles(oldPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse old files: %w", err)
	}

	newFiles, err := parser.ParseFiles(newPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse new files: %w", err)
	}

	oldMsgs := oldFiles[0].GetMessageTypes()
	newMsgs := newFiles[0].GetMessageTypes()

	var diffs []string

	for _, oldMsg := range oldMsgs {
		for _, newMsg := range newMsgs {
			if oldMsg.GetName() == newMsg.GetName() {
				oldFields := oldMsg.GetFields()
				newFields := newMsg.GetFields()

				oldMap := map[int32]string{}
				newMap := map[int32]string{}

				for _, f := range oldFields {
					oldMap[f.GetNumber()] = fmt.Sprintf("%s %s", f.GetType().String(), f.GetName())
				}

				for _, f := range newFields {
					newMap[f.GetNumber()] = fmt.Sprintf("%s %s", f.GetType().String(), f.GetName())
				}

				for num, oldDef := range oldMap {
					if newDef, ok := newMap[num]; !ok {
						diffs = append(diffs, fmt.Sprintf("Field #%d (%s) removed", num, oldDef))
					} else if oldDef != newDef {
						diffs = append(diffs, fmt.Sprintf("Field #%d changed: %s -> %s", num, oldDef, newDef))
					}
				}

				for num, newDef := range newMap {
					if _, ok := oldMap[num]; !ok {
						diffs = append(diffs, fmt.Sprintf("Field #%d (%s) added", num, newDef))
					}
				}
			}
		}
	}
	return diffs, nil
}
