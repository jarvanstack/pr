package util

import "regexp"

func Replace(data string) string {
	// 句号
	re := regexp.MustCompile(`([\p{Han}])\.[ ]*`)
	output := re.ReplaceAllString(data, `$1。`)

	// 逗号
	re = regexp.MustCompile(`([\p{Han}]),[ ]*`)
	output = re.ReplaceAllString(output, `$1，`)

	// 冒号
	re = regexp.MustCompile(`([\p{Han}]):[ ]*`)
	output = re.ReplaceAllString(output, `$1：`)

	// 分号
	re = regexp.MustCompile(`([\p{Han}]);[ ]*`)
	output = re.ReplaceAllString(output, `$1；`)

	// 问号
	re = regexp.MustCompile(`([\p{Han}])\?[ ]*`)
	output = re.ReplaceAllString(output, `$1？`)

	// 感叹号
	re = regexp.MustCompile(`([\p{Han}])![ ]*`)
	output = re.ReplaceAllString(output, `$1！`)

	// 左括号
	re = regexp.MustCompile(`[ ]*\(([\p{Han}])`)
	output = re.ReplaceAllString(output, `（$1`)

	// 右括号
	re = regexp.MustCompile(`([\p{Han}])\)[ ]*`)
	output = re.ReplaceAllString(output, `$1）`)

	// 左引号
	re = regexp.MustCompile(`[ ]*"([\p{Han}])`)
	output = re.ReplaceAllString(output, `“$1`)

	// 右引号
	re = regexp.MustCompile(`([\p{Han}])"[ ]*`)
	output = re.ReplaceAllString(output, `$1”`)

	// 左小引号
	re = regexp.MustCompile(`[ ]*'([\p{Han}])`)
	output = re.ReplaceAllString(output, `‘$1`)

	// 右小引号
	re = regexp.MustCompile(`([\p{Han}])'[ ]*`)
	output = re.ReplaceAllString(output, `$1’`)

	return output
}
