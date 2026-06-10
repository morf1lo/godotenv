package godotenv

import (
	"bufio"
	"os"
	"strings"
)

func Load(filenames ...string) error {
	if len(filenames) == 0 {
		filenames = append(filenames, ".env")
	}

	for _, filename := range filenames {
		if err := loadFromFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func loadFromFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		pair := scanner.Text()

		keyvalue := strings.SplitN(strings.TrimSpace(pair), "=", 2)
		if len(keyvalue) != 2 {
			continue
		}

		key := removeQuotes(keyvalue[0])
		value := removeQuotes(keyvalue[1])

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func removeQuotes(s string) string {
	s = strings.TrimSpace(s)
	
	chars := strings.Split(s, "")

	if len(chars) < 2 {
		return s
	}

	if (chars[0] == "\"" && chars[len(chars)-1] == "\"") || (chars[0] == "'" && chars[len(chars)-1] == "'") {
		chars = chars[1:len(chars)-1]
	}

	return strings.Join(chars, "")
}
