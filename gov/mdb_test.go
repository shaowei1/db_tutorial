package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

// 在测试函数开始前创建数据库文件，在测试结束后删除数据库文件
func setupAndTeardownDBFile() string {
	dbPath := createAndRemoveDBFile()
	// defer os.Remove(dbPath)
	return dbPath
}

func TestBTreeStructurePrinting(t *testing.T) {
	setupAndTeardownDBFile()
	var script []string
	for i := 1; i <= 14; i++ {
		script = append(script, fmt.Sprintf("insert %d user%d person%d@example.com", i, i, i))
	}
	script = append(script, ".btree")
	script = append(script, "insert 15 user15 person15@example.com")
	script = append(script, ".exit")

	result := runScript(script)

	expected := []string{
		"db > Tree:",
		"- internal (size 1)",
		"  - leaf (size 7)",
		"    - 1",
		"    - 2",
		"    - 3",
		"    - 4",
		"    - 5",
		"    - 6",
		"    - 7",
		"  - key 7",
		"  - leaf (size 7)",
		"    - 8",
		"    - 9",
		"    - 10",
		"    - 11",
		"    - 12",
		"    - 13",
		"    - 14",
		"db > Need to implement searching an internal node",
		"",
	}

	startIndex := 14 // Index from where the expected result starts
	if len(result) < startIndex+len(expected) {
		t.Errorf("Expected result length is longer than actual result length")
	}
	if !isEqual(result[startIndex:], expected) {
		t.Errorf("Expected output: \n%v, got: \n%v", expected, result[startIndex:])
	}
}

func TestDuplicateIDErrorMessage(t *testing.T) {
	setupAndTeardownDBFile()
	script := []string{
		"insert 1 user1 person1@example.com",
		"insert 1 user1 person1@example.com",
		"select",
		".exit",
	}
	result := runScript(script)
	expectedOutput := []string{
		"db > Executed.",
		"db > Error: Duplicate key.",
		"db > (1, user1, person1@example.com)",
		"Executed.",
		"db > ",
	}
	if !isEqual(result, expectedOutput) {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, result)
	}
}

func TestPrintBTreeStructure(t *testing.T) {
	setupAndTeardownDBFile()
	script := []string{
		"insert 3 user3 person3@example.com",
		"insert 1 user1 person1@example.com",
		"insert 2 user2 person2@example.com",
		".btree",
		".exit",
	}

	// 运行脚本并获取输出结果
	result := runScript(script)

	// 期望的输出结果
	expected := []string{
		"db > Executed.",
		"db > Executed.",
		"db > Executed.",
		"db > Tree:",
		"- leaf (size 3)",
		"  - 1",
		"  - 2",
		"  - 3",
		"db > ",
	}

	// 检查输出结果是否符合预期
	if !isEqual(result, expected) {
		t.Errorf("Expected output: %v, got: %v", expected, result)
	}
}

func TestPrintConstants(t *testing.T) {
	// 准备测试脚本
	script := []string{
		".constants",
		".exit",
	}

	// 运行脚本并获取输出结果
	result := runScript(script)

	// 期望的输出结果
	expected := []string{
		"db > Constants:",
		"ROW_SIZE: 293",
		"COMMON_NODE_HEADER_SIZE: 6",
		"LEAF_NODE_HEADER_SIZE: 10",
		"LEAF_NODE_CELL_SIZE: 297",
		"LEAF_NODE_SPACE_FOR_CELLS: 4086",
		"LEAF_NODE_MAX_CELLS: 13",
		"db > ",
	}

	// 检查输出结果是否符合预期
	if !isEqual(result, expected) {
		t.Errorf("Expected output: %v, got: %v", expected, result)
	}
}

func TestKeepsDataAfterClosingConnection(t *testing.T) {
	setupAndTeardownDBFile()

	// Run the first script
	result1 := runScript([]string{
		"insert 1 user1 person1@example.com",
		".exit",
	})
	expected1 := []string{
		"db > Executed.",
		"db > ",
	}
	if !isEqual(result1, expected1) {
		t.Errorf("Result1: got %v, want %v", result1, expected1)
	}

	// Run the second script
	result2 := runScript([]string{
		"select",
		".exit",
	})
	expected2 := []string{
		"db > (1, user1, person1@example.com)",
		"Executed.",
		"db > ",
	}
	if !isEqual(result2, expected2) {
		t.Errorf("Result2: got %v, want %v", result2, expected2)
	}
}

func TestNegativeIDErrorMessage(t *testing.T) {
	setupAndTeardownDBFile()
	script := []string{
		"insert -1 cstack foo@bar.com",
		"select",
		".exit",
	}
	result := runScript(script)

	expectedOutput := []string{
		"db > ID must be positive.",
		"db > Executed.",
		"db > ",
	}
	if !isEqual(result, expectedOutput) {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, result)
	}
}

func TestStringTooLongErrorMessage(t *testing.T) {
	longUsername := strings.Repeat("a", 33)
	longEmail := strings.Repeat("a", 256)
	script := []string{
		"insert 1 " + longUsername + " " + longEmail,
		"select",
		".exit",
	}
	result := runScript(script)

	expectedOutput := []string{
		"db > String is too long.",
		"db > Executed.",
		"db > ",
	}
	if !isEqual(result, expectedOutput) {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, result)
	}
}

func TestInsertMaxStringLength(t *testing.T) {
	longUsername := strings.Repeat("a", 32)
	longEmail := strings.Repeat("a", 255)
	script := []string{
		"insert 1 " + longUsername + " " + longEmail,
		"select",
		".exit",
	}

	result := runScript(script)

	expectedUsername := longUsername
	expectedEmail := longEmail
	expectedOutput := []string{
		"db > Executed.",
		"db > (1, " + expectedUsername + ", " + expectedEmail + ")",
		"Executed.",
		"db > ",
	}
	if !isEqual(result, expectedOutput) {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, result)
	}
}

// 测试插入和检索一行
func TestInsertAndRetrieveRow(t *testing.T) {
	setupAndTeardownDBFile()
	commands := []string{
		"insert 1 user1 person1@example.com",
		"select",
		".exit",
	}

	// 运行脚本并获取输出
	output := runScript(commands)

	// 验证输出是否符合预期
	expectedOutput := []string{
		"db > Executed.",
		"db > (1, user1, person1@example.com)",
		"Executed.",
		"db > ",
	}
	if !isEqual(output, expectedOutput) {
		t.Errorf("Unexpected output:\nGot: %v\nExpected: %v", output, expectedOutput)
	}
}

func TestTableFullErrorMessage(t *testing.T) {
	var script []string
	for i := 1; i <= 1401; i++ {
		script = append(script, "insert "+strconv.Itoa(i)+" user"+strconv.Itoa(i)+" person"+strconv.Itoa(i)+"@example.com")
	}
	script = append(script, ".exit")

	result := runScript(script)

	expectedErrorMessage := "db > Error: Table full."
	if result[len(result)-2] != expectedErrorMessage {
		t.Errorf("Expected error message: %s, got: %s", expectedErrorMessage, result[len(result)-2])
	}
}

// 运行脚本并返回输出结果
func runScript(commands []string) []string {
	dbPath := "test.db"
	cmd := exec.Command("../mdb", dbPath)

	// 创建输入输出缓冲区
	var out bytes.Buffer
	var in bytes.Buffer

	// 将命令写入输入缓冲区
	for _, command := range commands {
		in.WriteString(command + "\n")
	}

	// 将输入缓冲区关联到命令的标准输入
	cmd.Stdin = &in

	// 将输出缓冲区关联到命令的标准输出
	cmd.Stdout = &out

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("command execute error:", err)
		panic(err)
	}

	// 将输出按行分割并返回
	outputLines := bytes.Split(out.Bytes(), []byte("\n"))
	outputStrings := make([]string, len(outputLines))
	for i, line := range outputLines {
		outputStrings[i] = string(line)
	}

	return outputStrings
}

// 检查两个字符串切片是否相等
func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		fmt.Println("len(a):", len(a), "len(b):", len(b))
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			fmt.Println("a[i]:", a[i], "b[i]:", b[i])
			return false
		}
	}
	return true
}

func createAndRemoveDBFile() string {
	dbPath := "test.db"
	os.Remove(dbPath) // 删除已存在的数据库文件
	return dbPath
}
