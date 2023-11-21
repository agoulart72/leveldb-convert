package convert

import "strings"

func getOutputName(filename string, output string, suffix string) string {

	if len(strings.TrimSpace(output)) > 0 {
		if !strings.HasSuffix(output, suffix) {
			output = output + suffix
		}
		return output
	}

	path := ""
	outputFileName := filename

	if strings.Contains(filename, "/") {
		p := strings.Split(filename, "/")
		path = strings.Join(p[:len(p)-1], "/")
		outputFileName = p[len(p)-1]
	}

	if strings.Contains(outputFileName, ".") {
		fn := strings.Split(outputFileName, ".")
		outputFileName = strings.Join(fn[:len(fn)-1], ".")
	}

	if len(suffix) > 0 {
		outputFileName = outputFileName + suffix
	}

	if len(path) > 0 {
		outputFileName = path + "/" + outputFileName
	}

	return outputFileName
}
