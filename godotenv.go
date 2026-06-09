package godotenv

import (
	"bufio"
	"os"
	"strings"
)

func Load(paths ...string) error {
	if len(paths) == 0 {
		paths = append(paths, ".env")
	}

	for _, path := range paths {
		if err := loadFromFile(path); err != nil {
			return err
		}
	}

	return nil
}

func loadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		pair, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		keyvalue := strings.SplitN(strings.TrimSpace(pair), "=", 2)
		if len(keyvalue) != 2 {
			continue
		}

		key := removeQuotes(strings.TrimSpace(keyvalue[0]))
		value := removeQuotes(strings.TrimSpace(keyvalue[1]))

		os.Setenv(key, value)
	}

	return nil
}

func removeQuotes(s string) string {
	chars := strings.Split(s, "")

	if len(chars) == 0 {
		return s
	}

	if (chars[0] == "\"" && chars[len(chars)-1] == "\"") || (chars[0] == "'" && chars[len(chars)-1] == "'") {
		chars = chars[1:len(chars)-1]
	}

	return strings.Join(chars, "")
}
