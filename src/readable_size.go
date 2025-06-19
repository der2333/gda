package gda

import "fmt"

func ReadableSize(bytes int64) string {
	const unit = 1000
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	// 使用浮点数来保持精度
	value := float64(bytes)
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	index := 0

	for value >= unit && index < len(units)-1 {
		value /= unit
		index++
	}

	return fmt.Sprintf("%.2f %s", value, units[index])
}
