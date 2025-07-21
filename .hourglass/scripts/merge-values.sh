#!/usr/bin/env bash
set -e

DEFAULTS_FILE="$1"
USER_VALUES_FILE="$2"
CONTEXT_DATA="$3"
OUTPUT_FILE="$4"

# First, substitute context values into defaults using gomplate
TEMP_DEFAULTS=$(mktemp)

# Use gomplate to substitute variables in defaults file
# gomplate can read JSON data and substitute variables
echo "$CONTEXT_DATA" | gomplate -d 'context=stdin:?type=application/json' -f "$DEFAULTS_FILE" > "$TEMP_DEFAULTS"

# Then merge with user values (user values take precedence)
if [ -f "$USER_VALUES_FILE" ]; then
  # Use yq to deep merge, with user values overriding defaults
  yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' "$TEMP_DEFAULTS" "$USER_VALUES_FILE" > "$OUTPUT_FILE"
else
  cp "$TEMP_DEFAULTS" "$OUTPUT_FILE"
fi

rm "$TEMP_DEFAULTS" 