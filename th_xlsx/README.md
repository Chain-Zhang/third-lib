# 介绍

xlsx 是对新版本的Excel进行简单读写操作的golang第三方库，支持的文件格式是`.xlsx`

# 源码

https://github.com/tealeg/xlsx

# 安装

```
go get github.com/tealeg/xlsx
```

# 使用

## Read

excel 文件内容：

![excel数据](http://opgmvuzyu.bkt.clouddn.com/1519979756070.png)

读取excel文件代码：

```
import(
	"fmt"
	"github.com/tealeg/xlsx"
)

var (
	inFile = "/Users/chain/Downloads/student1.xlsx"
)

func Import(){
	// 打开文件
    xlFile, err := xlsx.OpenFile(inFile)
    if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遍历sheet页读取
    for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		//遍历行读取
        for _, row := range sheet.Rows {
			// 遍历每行的列读取
            for _, cell := range row.Cells {
				text := cell.String()
                fmt.Printf("%20s", text)
			}
			fmt.Print("\n")
        }
	}
	fmt.Println("\n\nimport success")
}
```

测试代码：

```
import(
	"testing"
)

func TestImport(t *testing.T){
	Import()
}
```

结果：

![导入结果](http://opgmvuzyu.bkt.clouddn.com/1519979958069.png)

## Write

```
import(
	"strconv"
	"fmt"
	"github.com/tealeg/xlsx"
)

var (
	inFile = "/Users/chain/Downloads/student1.xlsx"
	outFile = "/Users/chain/Downloads/out_student.xlsx"
)

type Student struct{
	Name string
	age int
	Phone string
	Gender string
	Mail string
}

func Export(){
    file := xlsx.NewFile()
    sheet, err := file.AddSheet("student_list")
    if err != nil {
        fmt.Printf(err.Error())
	}
	stus := getStudents()
	//add data
	for _, stu := range stus{
		row := sheet.AddRow()
		nameCell := row.AddCell()
		nameCell.Value = stu.Name

		ageCell := row.AddCell()
		ageCell.Value = strconv.Itoa(stu.age)

		phoneCell := row.AddCell()
		phoneCell.Value = stu.Phone

		genderCell := row.AddCell()
		genderCell.Value = stu.Gender

		mailCell := row.AddCell()
		mailCell.Value = stu.Mail
	}
    err = file.Save(outFile)
    if err != nil {
        fmt.Printf(err.Error())
	}
	fmt.Println("\n\nexport success")
}

func getStudents()[]Student{
	students := make([]Student, 0)
	for i := 0; i < 10; i++{
		stu := Student{}
		stu.Name = "name" + strconv.Itoa(i + 1)
		stu.Mail = stu.Name + "@chairis.cn"
		stu.Phone = "1380013800" + strconv.Itoa(i)
		stu.age = 20
		stu.Gender = "男"
		students = append(students, stu)
	}
	return students
}
```

测试代码：

```
import(
	"testing"
)

func TestExport(t *testing.T){
	Export()
}
```

导出结果：

![导出结果](http://opgmvuzyu.bkt.clouddn.com/1519980434042.png)