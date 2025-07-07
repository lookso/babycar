package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var _ = decimal.New(1, 1)

var dataMap = map[string]func(columnType gorm.ColumnType) (dataType string){
	"tinyint": func(columnType gorm.ColumnType) (dataType string) {
		return "int32"
	},
	"decimal": func(columnType gorm.ColumnType) (dataType string) {
		// 返回完整的类型定义（包含包名和精度）
		return "decimal.Decimal" // 不再拼接 detailType
	},
}

// GenerateModelFromDB 使用 gorm.io/gen 生成 GORM 模型
func GenerateModelFromDB(dsn string) error {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	// 检查并创建 dal 目录
	dalDir := "../internal/dal"
	if _, err := os.Stat(dalDir); os.IsNotExist(err) {
		err := os.MkdirAll(dalDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create dal directory: %v", err)
			return err
		}
	}

	// 初始化代码生成器
	g := gen.NewGenerator(gen.Config{
		OutPath: filepath.Join(dalDir, "query"),                                     // 模型输出路径
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成默认查询方法
		// FieldNullable:    true,
		// FieldWithTypeTag: true,
	})

	// 设置数据库连接
	g.UseDB(db)
	g.WithDataTypeMap(dataMap)
	g.ApplyBasic(
		// g.GenerateModelAs("order_service", "Order"),
		// g.GenerateModelAs("order_service_logs", "OrderLog"),
		// g.GenerateModelAs("order_service_details", "OrderItem"),
		// g.GenerateModelAs("order_service_except", "ExceptOrder"),
		// g.GenerateModelAs("order_service_except_logs", "ExceptOrderLog"),
		// g.GenerateModelAs("order_service_refunds", "OrderRefund"),
		// g.GenerateModelAs("order_service_refund_details", "OrderRefundItem"),
		// g.GenerateModelAs("order_payment_failed", "PaymentFailInfo"),
		// g.GenerateModelAs("order_platform_change_log", "OrderChangeLog"),

		g.GenerateAllTable()...,
	)
	// 执行生成
	g.Execute()

	fmt.Println("Models generated successfully")
	return nil
}

func main() {
	// 示例调用
	dsn := "root:Abc@1234@tcp(127.0.0.1:3306)/new_hh_product?parseTime=true&timeout=10s"
	if err := GenerateModelFromDB(dsn); err != nil {
		log.Fatalf("Error generating models: %v", err)
	}
}
