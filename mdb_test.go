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
	defer os.Remove(dbPath) // not executed
	return dbPath
}

func TestAllowPrintingOutStructureOfSevenLeafNodeBTree(t *testing.T) {
	setupAndTeardownDBFile()
	script := []string{
		"insert 58 user58 person58@example.com",
		"insert 56 user56 person56@example.com",
		"insert 8 user8 person8@example.com",
		"insert 54 user54 person54@example.com",
		"insert 77 user77 person77@example.com",
		"insert 7 user7 person7@example.com",
		"insert 25 user25 person25@example.com",
		"insert 71 user71 person71@example.com",
		"insert 13 user13 person13@example.com",
		"insert 22 user22 person22@example.com",
		"insert 53 user53 person53@example.com",
		"insert 51 user51 person51@example.com",
		"insert 59 user59 person59@example.com",
		"insert 32 user32 person32@example.com",
		"insert 36 user36 person36@example.com",
		"insert 79 user79 person79@example.com",
		"insert 10 user10 person10@example.com",
		"insert 33 user33 person33@example.com",
		"insert 20 user20 person20@example.com",
		"insert 4 user4 person4@example.com",
		"insert 35 user35 person35@example.com",
		"insert 76 user76 person76@example.com",
		"insert 49 user49 person49@example.com",
		"insert 24 user24 person24@example.com",
		"insert 70 user70 person70@example.com",
		"insert 48 user48 person48@example.com",
		"insert 39 user39 person39@example.com",
		"insert 15 user15 person15@example.com",
		"insert 47 user47 person47@example.com",
		"insert 30 user30 person30@example.com",
		"insert 86 user86 person86@example.com",
		"insert 31 user31 person31@example.com",
		"insert 68 user68 person68@example.com",
		"insert 37 user37 person37@example.com",
		"insert 66 user66 person66@example.com",
		"insert 63 user63 person63@example.com",
		"insert 40 user40 person40@example.com",
		"insert 78 user78 person78@example.com",
		"insert 19 user19 person19@example.com",
		"insert 46 user46 person46@example.com",
		"insert 14 user14 person14@example.com",
		"insert 81 user81 person81@example.com",
		"insert 72 user72 person72@example.com",
		"insert 6 user6 person6@example.com",
		"insert 50 user50 person50@example.com",
		"insert 85 user85 person85@example.com",
		"insert 67 user67 person67@example.com",
		"insert 2 user2 person2@example.com",
		"insert 55 user55 person55@example.com",
		"insert 69 user69 person69@example.com",
		"insert 5 user5 person5@example.com",
		"insert 65 user65 person65@example.com",
		"insert 52 user52 person52@example.com",
		"insert 1 user1 person1@example.com",
		"insert 29 user29 person29@example.com",
		"insert 9 user9 person9@example.com",
		"insert 43 user43 person43@example.com",
		"insert 75 user75 person75@example.com",
		"insert 21 user21 person21@example.com",
		"insert 82 user82 person82@example.com",
		"insert 12 user12 person12@example.com",
		"insert 18 user18 person18@example.com",
		"insert 60 user60 person60@example.com",
		"insert 44 user44 person44@example.com",
		".btree",
		".exit",
	}
	result := runScript(script)

	excepted := []string{
		"db > Tree:",
		"- internal (size 1)",
		"  - internal (size 2)",
		"    - leaf (size 7)",
		"      - 1",
		"      - 2",
		"      - 4",
		"      - 5",
		"      - 6",
		"      - 7",
		"      - 8",
		"    - key 8",
		"    - leaf (size 11)",
		"      - 9",
		"      - 10",
		"      - 12",
		"      - 13",
		"      - 14",
		"      - 15",
		"      - 18",
		"      - 19",
		"      - 20",
		"      - 21",
		"      - 22",
		"    - key 22",
		"    - leaf (size 8)",
		"      - 24",
		"      - 25",
		"      - 29",
		"      - 30",
		"      - 31",
		"      - 32",
		"      - 33",
		"      - 35",
		"  - key 35",
		"  - internal (size 3)",
		"    - leaf (size 12)",
		"      - 36",
		"      - 37",
		"      - 39",
		"      - 40",
		"      - 43",
		"      - 44",
		"      - 46",
		"      - 47",
		"      - 48",
		"      - 49",
		"      - 50",
		"      - 51",
		"    - key 51",
		"    - leaf (size 11)",
		"      - 52",
		"      - 53",
		"      - 54",
		"      - 55",
		"      - 56",
		"      - 58",
		"      - 59",
		"      - 60",
		"      - 63",
		"      - 65",
		"      - 66",
		"    - key 66",
		"    - leaf (size 7)",
		"      - 67",
		"      - 68",
		"      - 69",
		"      - 70",
		"      - 71",
		"      - 72",
		"      - 75",
		"    - key 75",
		"    - leaf (size 8)",
		"      - 76",
		"      - 77",
		"      - 78",
		"      - 79",
		"      - 81",
		"      - 82",
		"      - 85",
		"      - 86",
		"db > ",
	}
	if !isEqual(result[64:], excepted) {
		t.Errorf("Expected output: %v, got: %v", excepted, result[64:])
	}
}

func TestAllowPrintingOutStructureOfFourLeafNodeBTree(t *testing.T) {
	setupAndTeardownDBFile()
	script := []string{"insert 18 user18 person18@example.com",
		"insert 7 user7 person7@example.com",
		"insert 10 user10 person10@example.com",
		"insert 29 user29 person29@example.com",
		"insert 23 user23 person23@example.com",
		"insert 4 user4 person4@example.com",
		"insert 14 user14 person14@example.com",
		"insert 30 user30 person30@example.com",
		"insert 15 user15 person15@example.com",
		"insert 26 user26 person26@example.com",
		"insert 22 user22 person22@example.com",
		"insert 19 user19 person19@example.com",
		"insert 2 user2 person2@example.com",
		"insert 1 user1 person1@example.com",
		"insert 21 user21 person21@example.com",
		"insert 11 user11 person11@example.com",
		"insert 6 user6 person6@example.com",
		"insert 20 user20 person20@example.com",
		"insert 5 user5 person5@example.com",
		"insert 8 user8 person8@example.com",
		"insert 9 user9 person9@example.com",
		"insert 3 user3 person3@example.com",
		"insert 12 user12 person12@example.com",
		"insert 27 user27 person27@example.com",
		"insert 17 user17 person17@example.com",
		"insert 16 user16 person16@example.com",
		"insert 13 user13 person13@example.com",
		"insert 24 user24 person24@example.com",
		"insert 25 user25 person25@example.com",
		"insert 28 user28 person28@example.com",
		".btree",
		".exit",
	}
	result := runScript(script)
	excepted := []string{
		"db > Tree:",
		"- internal (size 3)",
		"  - leaf (size 7)",
		"    - 1",
		"    - 2",
		"    - 3",
		"    - 4",
		"    - 5",
		"    - 6",
		"    - 7",
		"  - key 7",
		"  - leaf (size 8)",
		"    - 8",
		"    - 9",
		"    - 10",
		"    - 11",
		"    - 12",
		"    - 13",
		"    - 14",
		"    - 15",
		"  - key 15",
		"  - leaf (size 7)",
		"    - 16",
		"    - 17",
		"    - 18",
		"    - 19",
		"    - 20",
		"    - 21",
		"    - 22",
		"  - key 22",
		"  - leaf (size 8)",
		"    - 23",
		"    - 24",
		"    - 25",
		"    - 26",
		"    - 27",
		"    - 28",
		"    - 29",
		"    - 30",
		"db > ",
	}
	if !isEqual(result[30:], excepted) {
		t.Errorf("Expected output: %v, got: %v", excepted, result[31:])
	}
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
		"db > Executed.",
		"db > ",
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
		"LEAF_NODE_HEADER_SIZE: 14",
		"LEAF_NODE_CELL_SIZE: 297",
		"LEAF_NODE_SPACE_FOR_CELLS: 4082",
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
	setupAndTeardownDBFile()
	var script []string
	for i := 1; i <= 1401; i++ {
		script = append(script, "insert "+strconv.Itoa(i)+" user"+strconv.Itoa(i)+" person"+strconv.Itoa(i)+"@example.com")
	}
	script = append(script, ".exit")

	result := runScript(script)

	expectedErrorMessage := []string{
		"db > Executed.",
		"db > Tried to fetch page number out of bounds. 101 > 100",
	}
	result = result[len(result)-2:]
	if !isEqual(result, expectedErrorMessage) {
		t.Errorf("Expected error message: %s, got: %s", expectedErrorMessage, result)
	}
}

// 运行脚本并返回输出结果
func runScript(commands []string) []string {
	dbPath := "test.db"
	cmd := exec.Command("./mdb", dbPath)

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
		// panic(err)
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
			fmt.Println("given[", i, "]:", a[i], "glength", len(a[i]), "expected[", i, "]:", b[i], "elength", len(b[i]))
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

func printResult(result []string) {
	for _, line := range result {
		fmt.Println(line)
	}
}
