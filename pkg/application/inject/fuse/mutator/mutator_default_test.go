package mutator

import (
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDefaultMutatorHelper_GenerateUniqueHostPath(t *testing.T) {
	// Test case struct
	type args struct {
		input               string
		metaObjName         string
		metaObjGenerateName string
	}

	tests := []struct {
		name                    string
		args                    args
		expectedHostPathPattern string
	}{
		{
			name: "with-metaobj-name",
			args: args{
				input:       "/runtime-mnt/juicefs/default/jfsdemo",
				metaObjName: "test",
			},
			expectedHostPathPattern: `^/runtime-mnt/juicefs/default/test/\d{16}-[a-zA-Z0-9]{8}/jfsdemo$`,
		},
		{
			name: "with-input-has-se-suffix",
			args: args{
				input:       "/runtime-mnt/juicefs/default/jfsdemo/",
				metaObjName: "test",
			},
			expectedHostPathPattern: `^/runtime-mnt/juicefs/default/test/\d{16}-[a-zA-Z0-9]{8}/jfsdemo$`,
		},
		{
			name: "with-metaobj-generate-name",
			args: args{
				input:               "/runtime-mnt/juicefs/default/jfsdemo",
				metaObjName:         "",
				metaObjGenerateName: "test",
			},
			expectedHostPathPattern: `^/runtime-mnt/juicefs/default/test--generate-name/\d{16}-[a-zA-Z0-9]{8}/jfsdemo$`,
		},
	}

	// Execute test cases
	for _, tt := range tests {
		mutatorHelper := &defaultMutatorHelper{
			Specs: &MutatingPodSpecs{
				MetaObj: metav1.ObjectMeta{
					Name:         tt.args.metaObjName,
					GenerateName: tt.args.metaObjGenerateName,
				},
			},
		}
		t.Run(tt.name, func(t *testing.T) {
			result, err := mutatorHelper.generateUniqueHostPath(tt.args.input)
			components := strings.Split(strings.TrimSuffix(result, string(filepath.Separator)), string(filepath.Separator))
			assert.EqualExportedValues(t, mutatorHelper.ctx.generatedUniquePathElem, filepath.Join(components[len(components)-3:len(components)-1]...))
			re, err := regexp.Compile(tt.expectedHostPathPattern)
			if err != nil {
				t.Errorf("regexp compile failed: %v", err)
			}
			assert.Equal(t, re.MatchString(result), true)
		})
	}
}
