package xgenny

import (
	"sort"
	"strings"

	"github.com/ignite/cli/v28/ignite/pkg/cliui/colors"
	"github.com/ignite/cli/v28/ignite/pkg/xfilepath"
)

var (
	modifyPrefix = colors.Modified("modify ")
	createPrefix = colors.Success("create ")
	removePrefix = func(s string) string {
		return strings.TrimPrefix(strings.TrimPrefix(s, modifyPrefix), createPrefix)
	}
)

// SourceModification describes modified and created files in the source code after a run.
type SourceModification struct {
	modified map[string]struct{}
	created  map[string]struct{}
}

func NewSourceModification() SourceModification {
	return SourceModification{
		make(map[string]struct{}),
		make(map[string]struct{}),
	}
}

// ModifiedFiles returns the modified files of the source modification.
func (sm SourceModification) ModifiedFiles() (modifiedFiles []string) {
	for modified := range sm.modified {
		modifiedFiles = append(modifiedFiles, modified)
	}
	return
}

// CreatedFiles returns the created files of the source modification.
func (sm SourceModification) CreatedFiles() (createdFiles []string) {
	for created := range sm.created {
		createdFiles = append(createdFiles, created)
	}
	return
}

// AppendModifiedFiles appends modified files in the source modification that are not already documented.
func (sm *SourceModification) AppendModifiedFiles(modifiedFiles ...string) {
	for _, modifiedFile := range modifiedFiles {
		_, alreadyModified := sm.modified[modifiedFile]
		_, alreadyCreated := sm.created[modifiedFile]
		if !alreadyModified && !alreadyCreated {
			sm.modified[modifiedFile] = struct{}{}
		}
	}
}

// AppendCreatedFiles appends a created files in the source modification that are not already documented.
func (sm *SourceModification) AppendCreatedFiles(createdFiles ...string) {
	for _, createdFile := range createdFiles {
		_, alreadyModified := sm.modified[createdFile]
		_, alreadyCreated := sm.created[createdFile]
		if !alreadyModified && !alreadyCreated {
			sm.created[createdFile] = struct{}{}
		}
	}
}

// Merge merges new source modification to an existing one.
func (sm *SourceModification) Merge(newSm SourceModification) {
	sm.AppendModifiedFiles(newSm.ModifiedFiles()...)
	sm.AppendCreatedFiles(newSm.CreatedFiles()...)
}

// String convert to string value.
func (sm *SourceModification) String() (string, error) {
	// get file names and add prefix
	files := make([]string, 0)
	for _, modified := range sm.ModifiedFiles() {
		// get the relative app path from the current directory
		relativePath, err := xfilepath.RelativePath(modified)
		if err != nil {
			return "", err
		}
		files = append(files, modifyPrefix+relativePath)
	}
	for _, created := range sm.CreatedFiles() {
		// get the relative app path from the current directory
		relativePath, err := xfilepath.RelativePath(created)
		if err != nil {
			return "", err
		}
		files = append(files, createPrefix+relativePath)
	}

	// sort filenames without prefix
	sort.Slice(files, func(i, j int) bool {
		s1 := removePrefix(files[i])
		s2 := removePrefix(files[j])

		return strings.Compare(s1, s2) == -1
	})

	return "\n" + strings.Join(files, "\n"), nil
}
