package processor

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrStartMarkerNotFound = errors.New("start marker not found")
	ErrEndMarkerNotFound   = errors.New("end marker not found")
)

// FindMarkerPositions locates the start and end positions for content replacement
// between the specified markers in the given content string.
// Returns the position after the start marker and the position of the end marker.
func FindMarkerPositions(content, startMarker, endMarker string) (startPos, endPos int, err error) {
	startIdx := strings.Index(content, startMarker)
	if startIdx == -1 {
		return -1, -1, ErrStartMarkerNotFound
	}
	
	startPos = startIdx + len(startMarker)
	
	endIdx := strings.Index(content, endMarker)
	if endIdx == -1 {
		return -1, -1, ErrEndMarkerNotFound
	}
	
	endPos = endIdx
	
	return startPos, endPos, nil
}

// ReplaceContentBetweenMarkers replaces the content between the specified positions
// with the new content and returns the modified string.
func ReplaceContentBetweenMarkers(content, newContent string, startPos, endPos int) string {
	prefix := content[:startPos]
	suffix := content[endPos:]
	
	return prefix + newContent + suffix
}

// ProcessFile reads a file, finds content between markers, replaces it with new content,
// and writes the result back to the file.
func ProcessFile(filepath, startMarker, endMarker, newContent string) error {
	// Read file content
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	
	// Convert to string once and reuse
	contentStr := string(content)
	
	// Find marker positions
	startPos, endPos, err := FindMarkerPositions(contentStr, startMarker, endMarker)
	if err != nil {
		return err
	}
	
	// Replace content between markers
	newFileContent := ReplaceContentBetweenMarkers(contentStr, "\n"+newContent+"\n", startPos, endPos)
	
	// Write back to file
	return os.WriteFile(filepath, []byte(newFileContent), 0644)
}