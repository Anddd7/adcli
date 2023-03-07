package blog

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a blog to sidebar",
	RunE: func(cmd *cobra.Command, args []string) error {
		return publish()
	},
}

var filepath string

func init() {
	publishCmd.Flags().StringVarP(&filepath, "filepath", "f", "", "file of the blog post")
	publishCmd.MarkFlagRequired("title")
}

func publish() error {
	wd, _ := os.Getwd()
	path := fmt.Sprintf("%s/%s", wd, filepath)
	paths := strings.Split(filepath, "/")

	title, err := parseTitle(path)
	if err != nil {
		return err
	}
	year, _ := strconv.Atoi(paths[2])
	filename := paths[3][:len(paths[3])-3]

	obj := &blog{
		title:    title,
		year:     year,
		filename: filename,
	}

	err = obj.insert2sidebar()
	if err != nil {
		return err
	}

	return nil
}

func parseTitle(path string) (string, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		return "", err
	}

	br := bufio.NewReader(f)
	title, _, err := br.ReadLine()
	if err != nil {
		return "", err
	}

	return string(title)[2:], nil
}

func (b *blog) insert2sidebar() error {
	wd, _ := os.Getwd()
	path := fmt.Sprintf("%s/%s", wd, "docs/_sidebar.md")
	var lines []string

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	indicator := fmt.Sprintf("* [%d](/)", b.year)
	for i, line := range lines {
		if line == indicator {
			lines = append(lines[:i+1], lines[i:]...)
			lines[i+1] = fmt.Sprintf("** [%s](post/%d/%s \"%s\")", b.title, b.year, b.filename, b.title)

			break
		}
	}

	fo, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0666)
	defer fo.Close()
	writer := bufio.NewWriter(fo)
	for _, line := range lines {
		println(line)
		_, _ = writer.WriteString(line + "\n")
	}
	writer.Flush()

	return nil
}
