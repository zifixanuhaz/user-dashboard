// parser.js

/**
 * Parses a CSV string into an array of objects.
 *
 * @param {string} csvString - The CSV string to parse.
 * @param {object} options - Configuration options for parsing.
 * @param {string} [options.delimiter=','] - The delimiter character.
 * @param {boolean} [options.header=true] - Whether the first row is a header row.
 * @returns {Array<object>} An array of objects representing the parsed CSV data.
 * @throws {Error} If the CSV string is invalid.
 */
function parseCSV(csvString, options = {}) {
  const delimiter = options.delimiter || ',';
  const hasHeader = options.header !== false; // Default to true if not explicitly false

  if (typeof csvString !== 'string') {
    throw new Error('Invalid input: CSV string must be a string.');
  }

  const lines = csvString.trim().split('\n');
  if (lines.length === 0) {
    return []; // Return empty array for empty CSV
  }

  let header = [];
  let data = [];

  if (hasHeader) {
    header = lines.shift().split(delimiter).map(h => h.trim());
  } else {
    // Create default headers if no header row is provided
    header = Array.from({ length: lines[0].split(delimiter).length }, (_, i) => `column${i + 1}`);
  }

  for (const line of lines) {
    const values = line.split(delimiter).map(v => v.trim());
    if (values.length !== header.length) {
      console.warn("Row length does not match header length. Skipping row.");
      continue; // Skip rows with mismatching length. Log warning for debugging.
    }

    const obj = {};
    for (let i = 0; i < header.length; i++) {
      obj[header[i]] = values[i];
    }
    data.push(obj);
  }

  return data;
}

/**
 * Parses a JSON string into a JavaScript object.
 *
 * @param {string} jsonString - The JSON string to parse.
 * @returns {object} The parsed JavaScript object.
 * @throws {Error} If the JSON string is invalid.
 */
function parseJSON(jsonString) {
  if (typeof jsonString !== 'string') {
    throw new Error('Invalid input: JSON string must be a string.');
  }

  try {
    return JSON.parse(jsonString);
  } catch (error) {
    throw new Error('Invalid JSON string: ' + error.message);
  }
}

/**
 * Parses a URL query string into an object.
 *
 * @param {string} queryString - The URL query string to parse (e.g., "?key1=value1&key2=value2").
 * @returns {object} An object representing the parsed query parameters.
 */
function parseQueryString(queryString) {
    const params = {};
    if (!queryString) {
        return params; // Return empty object if no query string
    }

    const trimmedQueryString = queryString.startsWith('?') ? queryString.slice(1) : queryString;
    const pairs = trimmedQueryString.split('&');

    for (const pair of pairs) {
        const [key, value] = pair.split('=').map(decodeURIComponent);
        if (key) {
            params[key] = value || ''; // Store empty string if no value
        }
    }

    return params;
}

export { parseCSV, parseJSON, parseQueryString };