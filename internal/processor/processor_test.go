package processor

import (
	"os"
	"testing"
)

func TestFindMarkerPositions(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		startMarker string
		endMarker   string
		wantStart   int
		wantEnd     int
		wantErr     bool
	}{
		{
			name:        "valid markers",
			content:     "Hello\n<!-- BEGIN -->\nold content\n<!-- END -->\nWorld",
			startMarker: "<!-- BEGIN -->",
			endMarker:   "<!-- END -->",
			wantStart:   20,
			wantEnd:     33,
			wantErr:     false,
		},
		{
			name:        "missing start marker",
			content:     "Hello\nold content\n<!-- END -->\nWorld",
			startMarker: "<!-- BEGIN -->",
			endMarker:   "<!-- END -->",
			wantStart:   -1,
			wantEnd:     -1,
			wantErr:     true,
		},
		{
			name:        "missing end marker",
			content:     "Hello\n<!-- BEGIN -->\nold content\nWorld",
			startMarker: "<!-- BEGIN -->",
			endMarker:   "<!-- END -->",
			wantStart:   -1,
			wantEnd:     -1,
			wantErr:     true,
		},
		{
			name:        "custom markers",
			content:     "Start\n// CUSTOM START\ncode here\n// CUSTOM END\nEnd",
			startMarker: "// CUSTOM START",
			endMarker:   "// CUSTOM END",
			wantStart:   21,
			wantEnd:     32,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd, err := FindMarkerPositions(tt.content, tt.startMarker, tt.endMarker)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMarkerPositions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStart != tt.wantStart {
				t.Errorf("FindMarkerPositions() gotStart = %v, want %v", gotStart, tt.wantStart)
			}
			if gotEnd != tt.wantEnd {
				t.Errorf("FindMarkerPositions() gotEnd = %v, want %v", gotEnd, tt.wantEnd)
			}
		})
	}
}

func TestReplaceContentBetweenMarkers(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		newContent string
		startPos   int
		endPos     int
		want       string
	}{
		{
			name:       "simple replacement",
			content:    "Hello\n<!-- BEGIN -->\nold content\n<!-- END -->\nWorld",
			newContent: "\nnew content\n",
			startPos:   20,
			endPos:     33,
			want:       "Hello\n<!-- BEGIN -->\nnew content\n<!-- END -->\nWorld",
		},
		{
			name:       "empty new content",
			content:    "Hello\n<!-- BEGIN -->\nold content\n<!-- END -->\nWorld",
			newContent: "\n",
			startPos:   20,
			endPos:     33,
			want:       "Hello\n<!-- BEGIN -->\n<!-- END -->\nWorld",
		},
		{
			name:       "multiline replacement",
			content:    "Start\n// BEGIN\nold\ncode\n// END\nEnd",
			newContent: "\nnew\ncode\nhere\n",
			startPos:   14,
			endPos:     24,
			want:       "Start\n// BEGIN\nnew\ncode\nhere\n// END\nEnd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceContentBetweenMarkers(tt.content, tt.newContent, tt.startPos, tt.endPos)
			if got != tt.want {
				t.Errorf("ReplaceContentBetweenMarkers() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestProcessFile(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		startMarker string
		endMarker   string
		newContent  string
		want        string
		wantErr     bool
	}{
		{
			name:        "successful file processing",
			fileContent: "Header\n<!-- BEGIN -->\nold content\n<!-- END -->\nFooter",
			startMarker: "<!-- BEGIN -->",
			endMarker:   "<!-- END -->",
			newContent:  "new content",
			want:        "Header\n<!-- BEGIN -->\nnew content\n<!-- END -->\nFooter",
			wantErr:     false,
		},
		{
			name:        "missing start marker",
			fileContent: "Header\nold content\n<!-- END -->\nFooter",
			startMarker: "<!-- BEGIN -->",
			endMarker:   "<!-- END -->",
			newContent:  "new content",
			want:        "",
			wantErr:     true,
		},
		{
			name:        "missing end marker",
			fileContent: "Header\n<!-- BEGIN -->\nold content\nFooter",
			startMarker: "<!-- BEGIN -->",
			endMarker:   "<!-- END -->",
			newContent:  "new content",
			want:        "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpFile, err := os.CreateTemp("", "test-*.txt")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer func() {
				tmpFile.Close()
				os.Remove(tmpFile.Name())
			}()

			// Write test content to file
			if _, err := tmpFile.WriteString(tt.fileContent); err != nil {
				t.Fatalf("Failed to write to temp file: %v", err)
			}
			tmpFile.Close()

			// Test ProcessFile
			err = ProcessFile(tmpFile.Name(), tt.startMarker, tt.endMarker, tt.newContent)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return // Don't check file content if we expected an error
			}

			// Read file content back
			content, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				t.Fatalf("Failed to read temp file: %v", err)
			}

			if string(content) != tt.want {
				t.Errorf("ProcessFile() file content = %q, want %q", string(content), tt.want)
			}
		})
	}
}