package th_xlsx

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

// 导出到excel
func Export(){
	// 创建文件
	file := xlsx.NewFile()
	//添加sheet页
    sheet, err := file.AddSheet("student_list")
    if err != nil {
        fmt.Printf(err.Error())
	}
	stus := getStudents()
	//add data
	for _, stu := range stus{
		// 添加行
		row := sheet.AddRow()

		// 添加数据
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
	// 保存到指定文件
    err = file.Save(outFile)
    if err != nil {
        fmt.Printf(err.Error())
	}
	fmt.Println("\n\nexport success")
}

// 获取数据
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